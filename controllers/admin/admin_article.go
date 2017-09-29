package admin

import (
    "blog/controllers/base"
    "blog/common"
    "github.com/astaxie/beego/logs"
    "fmt"
    "net/http"
    "blog/models/db"
    "github.com/astaxie/beego"
    "time"
    "strconv"
)

/**
 article相关的topic
 */
type TopicOfArticle struct {
    ID          int
    Name        string
    Checked     string                    // 是否被当前文章选中
}

/**
 获取topics
 */
func (a *ArticleEditController) GetTopics() []*TopicOfArticle {
    var tpOfarticle []*TopicOfArticle
    tp := new(db.Topic)
    tps, err := tp.SelectAll()
    if err == nil {
        for _, item := range tps {
            tpOfarticle = append(tpOfarticle, &TopicOfArticle{ID: item.Id, Name: item.Name})
        }
    }

    return tpOfarticle
}

/**
 为article设置checked的topic
 */
func (a *ArticleEditController) SetTopicChecked(aID int, tpOfarticle []*TopicOfArticle) ([]*TopicOfArticle, string) {
    isDefault := "checked"

    ar := new(db.ArticleTopic)
    ats, err := ar.SelectByArticle(aID)
    if err != nil {
        for _, item := range ats {
            for _, tp := range tpOfarticle {
                if item.TopicId == tp.ID {
                    tp.Checked = "checked"
                    isDefault = ""
                    break
                }
            }
        }
    }

    return tpOfarticle, isDefault
}

/**
 article相关的tag
 */
type TagOfArticle struct {
    ID          int
    Name        string
}

/**
 编辑信息
 */
type Edit struct {
    Title           string
    IsDraft         bool
    Slug            string
    Content         string
    PublishTime     time.Time
}

type ArticleEditController struct {
    base.AdminCommonCtr
}

/**
 进入文章编辑、新建文章的界面
 */
func (a *ArticleEditController) Get ()  {
    // 预加载执行
    a.AdminBase()

    // 专题内容准备
    isDefaultTp := ""
    tpOfArticle := a.GetTopics()
    articleID, err := a.GetInt("articleid")
    if err != nil || articleID == 0 {
        // todo: 表示新建文章
    } else {
        // todo: 编辑文章
        tpOfArticle, isDefaultTp = a.SetTopicChecked(articleID, tpOfArticle)
        a.SetSession("articleid", articleID)

        // 准备article的内容
        article := new(db.Article)
        article.Id = articleID
        err = article.Read("id")
        if err != nil {
            logs.Error(fmt.Sprintf("get article'%d' info failed: %s", articleID, err.Error()))
        }

        edit := Edit{
            Title: article.Title,
            Slug: article.Url,
            Content: article.Content,
            PublishTime: time.Unix(article.PublishTime, 0),
            IsDraft: true,
        }
        a.Data["Edit"] = edit
    }
    a.Data["DefaultTp"] = isDefaultTp
    a.Data["Series"] = tpOfArticle

    // 准备域名
    a.Data["Domain"] = beego.AppConfig.String(beego.BConfig.RunMode + "::domain")

    // todo: tag处理
    //tag := models.Tag{}
    //tags, err := tag.SelectAll()

    // 加载template
    a.TplName = "admin/admin_article.html"
    return
}

/**
 文章提交的操作
 todo: 参考(https://beego.me/docs/mvc/controller/flash.md)flash数据处理的方式，表单提交错误的处理方式
 */
func (a *ArticleEditController) Post()  {
    // 预加载
    a.AdminBase()

    // 数据获取
    title := a.GetString("title")               // 标题
    slug := a.GetString("slug")                 // 连接url
    text := a.GetString("text")                 // 文本
    date := a.GetString("date")                 // 日期，格式为：2017-07-18 11:26   不涉及到秒
    serie := a.GetString("serie")               // 专题
    tags := a.GetString("tags")                 // 标签
    isUpdate := a.GetString("update")           // 标记是否需要更新时间

    // todo: 检查serie的合法性
    tpID, err := strconv.Atoi(serie)
    if err != nil {
        logs.Error(fmt.Sprintf("input topic id error: %s", err.Error()))
        return
    }

    // 入参检查
    if title == "" || slug == "" {
        logs.Error(fmt.Sprintf("param error: title/slug"))
        return
    }
    // todo: tags
    tags = tags

    // 构造数据
    article := new(db.Article)
    article.Title = title
    article.Url = slug
    article.Author = a.UserInfo.Id
    article.Updated = time.Now().Unix()
    article.Content = text

    // 判读是insert还是update处理
    articleID := a.GetSession("articleid")
    if articleID != nil {
        // update
        article.Id = articleID.(int)
        if isUpdate == "true" {
            article.PublishTime = a.dealPublishTime(date).Unix()
        }
    } else {
        // insert
        article.Id = 0
        article.PublishTime = a.dealPublishTime(date).Unix()
        article.Create = time.Now().Unix()
    }

    // 动作方式的处理
    do := a.GetString("do")
    logs.Debug(fmt.Sprintf("action=%s", do))
    if do == "save" {
        article.Status = common.ARTICLE_STATUS_DRAFT                                // 草稿
        err := a.saveArticle(article, tpID)
        if err != nil {
            logs.Error(err.Error())
            return
        }
        a.Redirect("/admin/manage-drafts", http.StatusFound)
    } else if do == "publish" {
        article.Status = common.ARTICLE_STATUS_PUBLISH                              // 草稿
        err := a.saveArticle(article, tpID)
        if err != nil {
            logs.Error(err.Error())
            return
        }
        a.Redirect("/admin/manage-articles", http.StatusFound)
    } else if do == "auto" {
        // todo: 自动保存的处理
    }

    return
}

/**
 对于文章保存的处理
 */
func (a *ArticleEditController) saveArticle(article *db.Article, tpID int) error {
    if article.Id != 0 {
        err := article.Update("title", "url", "author", "publish_time", "content", "create", "updated", "status")
        if err != nil {
            return fmt.Errorf("article'%d' update failed: %s", article.Title, err.Error())
        }
        // 处理article topic
        err = a.dealTopic(article.Id, tpID)
        if err != nil {
            return fmt.Errorf("when update article, deal topic failed: %s", err.Error())
        }
    } else {
        // 新建文章
        arID, err := article.Insert()
        if err != nil {
            return fmt.Errorf("article'%s' insert failed: %s", err.Error(), article.Title)
        }
        // 处理article topic
        err = a.dealTopic(int(arID), tpID)
        if err != nil {
            return fmt.Errorf("when insert article, deal topic failed: %s", err.Error())
        }
        // todo: 新建ArticleTag数据
    }

    return nil
}

/**
 更新article对应的
 */
func (a *ArticleEditController) dealTopic(articleID int, tpID int) error {
    // 先确认topic的有效性
    tp := &db.Topic{Id: tpID}
    err := tp.Read("id")
    if err != nil {
        logs.Warning(fmt.Sprintf("read topic for '%d' failed: %s", tpID, err.Error()))
        tp.Id = 1                      // 就将1定义为缺省的topic id值，后续需要录入表中
    }

    // 确认article_topic表中存在对应数据
    arTp := &db.ArticleTopic{ArticleId: articleID, TopicId: tp.Id}
    err = arTp.Read("article_id")
    if err != nil {
        // 不存在数据
        arTp.Create = time.Now().Unix()
        arTp.Updated = arTp.Create
        arTp.Status = common.STATUS_VALID
        _, err := arTp.Insert()
        if err != nil {
            logs.Error(fmt.Sprintf("insert article topic failed: %s", err.Error()))
            return fmt.Errorf("insert article topic failed: %s", err.Error())
        }
    } else {
        // 存在数据
        if arTp.TopicId != tp.Id {
            // 更新topic id值
            arTp.TopicId = tp.Id
            arTp.Updated = time.Now().Unix()
            err = arTp.Update("topic_id", "updated")
            if err != nil {
                logs.Error(fmt.Sprintf("update article_topic failed: %s", err.Error()))
                return fmt.Errorf("update article_topic failed: %s", err.Error())
            }
        }
    }

    return nil
}

/**
 处理发布时间
 */
func (a *ArticleEditController) dealPublishTime(date string) time.Time {
    if date == "" {
        return time.Now()
    }

    pt, err := time.Parse(common.TIME_LAYOUT_STR, date + ":00")
    if err != nil {
        logs.Error(fmt.Sprintf("publish time parse error:%s", err.Error()))
        pt = time.Now()
    }

    return pt
}