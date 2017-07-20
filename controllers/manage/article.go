package manage

import (
    "blog/controllers/base"
    "blog/common"
    "github.com/astaxie/beego/logs"
    "fmt"
    "net/http"
    "blog/models"
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
    Checked     bool                    // 是否被当前文章选中
}

/**
 获取topics
 */
func (a *ArticleEditController) GetTopics() []*TopicOfArticle {
    var tpOfarticle []*TopicOfArticle
    tp := new(models.Topic)
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
func (a *ArticleEditController) SetTopicChecked(aID int, tpOfarticle []*TopicOfArticle) []*TopicOfArticle {
    ar := new(models.ArticleTopic)
    ats, err := ar.SelectByArticle(aID)
    if err != nil {
        for _, item := range ats {
            for _, tp := range tpOfarticle {
                if item.TopicId == tp.ID {
                    tp.Checked = true
                    break
                }
            }
        }
    }

    return tpOfarticle
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
    tpOfarticle := a.GetTopics()
    articleID, err := a.GetInt("articleid")
    if err != nil || articleID == 0 {
        // todo: 表示新建文章
    } else {
        // todo: 编辑文章
        tpOfarticle = a.SetTopicChecked(articleID, tpOfarticle)
        a.SetSession("articleid", articleID)

        // 准备article的内容
        article := new(models.Article)
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
    a.Data["Series"] = tpOfarticle

    // 准备域名
    a.Data["Domain"] = beego.AppConfig.String(beego.BConfig.RunMode + "::domain")

    // todo: tag处理
    //tag := models.Tag{}
    //tags, err := tag.SelectAll()

    // 加载template
    a.TplName = "admin/article.html"
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
    if title == "" || slug == "" || text == "" || date == "" || serie == ""{
        //err := fmt.Errorf("param error")
        logs.Error(fmt.Sprintf("param error: title/slug/text/date/serie"))
        return
    }
    logs.Debug(fmt.Sprintf("title: %s", title))
    logs.Debug(fmt.Sprintf("slug: %s", slug))
    logs.Debug(fmt.Sprintf("text: %s", text))
    logs.Debug(fmt.Sprintf("date: %s", date))
    logs.Debug(fmt.Sprintf("serie: %s", serie))
    logs.Debug(fmt.Sprintf("tags: %s", tags))
    logs.Debug(fmt.Sprintf("update: %s", isUpdate))

    // 构造数据
    article := new(models.Article)
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
            pt, err := time.Parse(common.TIME_LAYOUT_STR, date + ":00")
            if err != nil {
                logs.Error(fmt.Sprintf("publish time parse error:%s", err.Error()))
                pt = time.Now()
            }
            article.PublishTime = pt.Unix()
        }
    } else {
        // insert
        article.Id = 0
        pt, err := time.Parse(common.TIME_LAYOUT_STR, date + ":00")
        if err != nil {
            logs.Error(fmt.Sprintf("publish time parse error:%s", err.Error()))
            pt = time.Now()
        }
        article.PublishTime = pt.Unix()
        article.Create = time.Now().Unix()
    }

    // 动作方式的处理
    do := a.GetString("do")
    logs.Debug(fmt.Sprintf("action=%s", do))
    if do == "save" {
        article.Status = common.ARTICLE_STATUS_DRAFT
        if article.Id != 0 {
            err := article.Update("title", "url", "author", "publish_time", "content", "create", "updated", "status")
            if err != nil {
                logs.Error(fmt.Sprintf("article'%s' update failed: %s", err.Error(), article.Title))
            }
        } else {
            // 新建文章
            arID, err := article.Insert()
            if err != nil {
                logs.Error(fmt.Sprintf("article'%s' insert failed: %s", err.Error(), article.Title))
            }
            // 新建ArticleTopic
            at := &models.ArticleTopic{ArticleId: int(arID), TopicId: tpID, Create: time.Now().Unix(),
                Updated: time.Now().Unix(), Status: common.STATUS_VALID}
            // todo:
            at = at

        }
        a.Redirect("/admin/manage-drafts", http.StatusFound)
    } else if do == "publish" {
        a.Redirect("/admin/manage-articles", http.StatusFound)
    } else if do == "auto" {
        // todo: 自动保存的处理
    }

    return
}