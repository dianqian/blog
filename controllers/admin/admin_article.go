package admin

import (
    "nest/controllers/base"
    "nest/common"
    "github.com/astaxie/beego/logs"
    "fmt"
    "nest/models/db"
    "time"
    "strconv"
    "nest/models/html/htmladmin"
    "strings"
)

type ArticleEditController struct {
    base.AdminCommonCtr
}

/**
 @Description：进入文章编辑、新建文章的界面
 @Param:
 @Return：
*/
func (a *ArticleEditController) Get ()  {
    // 预加载执行
    a.AdminBase()

    htmlArticleEditData := htmladmin.HTMLArticleEditData{}

    // 专题内容准备
    isDefaultTp := ""
    tp := &htmladmin.TopicOfArticle{}
    tps := tp.GetTopics()

    toa := htmladmin.TagOfArticle{}
    articleID, err := a.GetInt("articleid")
    if err != nil || articleID == 0 {
        // todo: 表示新建文章
    } else {
        // todo: 编辑文章
        tps, isDefaultTp = tp.SetTopicChecked(articleID, tps)
        a.SetSession("articleid", articleID)

        // 准备article的内容
        article := &db.Article{Id: articleID}
        err = article.Read("id")
        if err != nil {
            logs.Error(fmt.Sprintf("get article'%d' info failed: %s", articleID, err.Error()))
        }

        edit := htmladmin.Edit{
            IsEdit: true,
            Title: article.Title,
            Slug: article.Url,
            Content: article.Content,
            PublishTime: time.Unix(article.PublishTime, 0),
            IsDraft: true,
        }
        htmlArticleEditData.Edit = edit

        // 获取article相关的tag值
        toas, err := toa.GetTags(articleID)
        if err != nil {
            logs.Error(fmt.Sprintf("get tag for article'%s' failed: %s", article.Title, err.Error()))
            htmlArticleEditData.ArticleTags = "test, test1, test2"
        } else {
            htmlArticleEditData.ArticleTags = strings.Join(toas, ",")
        }
    }

    htmlArticleEditData.AllTags, _ = toa.GetAllTags()
    htmlArticleEditData.DefaultTp = isDefaultTp
    htmlArticleEditData.Series = tps

    // 加载template，设置html和data
    a.Data["HTMLArticleEditData"] = htmlArticleEditData
    a.TplName = "admin/admin_article.html"
    return
}

/**
 @Description：草稿保存处理
 @Param:
 @Return：
*/
func (a *ArticleEditController) dealDrafts() error {
    // 数据获取
    title := a.GetString("title")               // 标题
    slug := a.GetString("slug")                 // 连接url
    text := a.GetString("text")                 // 文本
    date := a.GetString("date")                 // 日期，格式为：2017-07-18 11:26   不涉及到秒
    topic := a.GetString("topic")               // 专题
    tags := a.GetString("tags")                 // 标签
    isUpdate := a.GetString("update")           // 标记是否需要更新时间

    if title == "" {
        return fmt.Errorf("`title` input empty")
    }
    if slug == "" {
        return fmt.Errorf("`slug` input empty")
    }
    tpID, err := strconv.Atoi(topic)
    if err != nil {
        return fmt.Errorf("`topic` input error")
    }

    // 直接执行保存处理
    article := &db.Article{
        Title: title,
        Url: slug,
        Author: a.UserInfo.Id,
        Updated: time.Now().Unix(),
        Content: text,
    }
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
        article.Updated = article.Create
    }

    article.Status = common.ARTICLE_STATUS_DRAFT
    // todo: tags的处理
    err = a.saveArticle(article, tpID, tags)
    if err != nil {
        return fmt.Errorf("deal drafts error: %s", err.Error())
    }

    return nil
}

/**
 @Description：文章发布处理
 @Param:
 @Return：
*/
func (a *ArticleEditController) dealPublish () error {
    // 数据获取
    title := a.GetString("title")               // 标题
    slug := a.GetString("slug")                 // 连接url
    text := a.GetString("text")                 // 文本
    date := a.GetString("date")                 // 日期，格式为：2017-07-18 11:26   不涉及到秒
    topic := a.GetString("topic")               // 专题
    tags := a.GetString("tags")                 // 标签
    isUpdate := a.GetString("update")           // 标记是否需要更新时间

    if title == "" {
        return fmt.Errorf("`title` input empty")
    }
    if slug == "" {
        return fmt.Errorf("`slug` input empty")
    }
    if text == "" {
        return fmt.Errorf("`text` input empty")
    }
    if date == "" {
        return fmt.Errorf("`date` input empty")
    }
    tpID, err := strconv.Atoi(topic)
    if err != nil {
        return fmt.Errorf("`topic` input error")
    }
    if tags == "" {
        return fmt.Errorf("`tags` input empty")
    }

    article := &db.Article{
        Title: title,
        Url: slug,
        Author: a.UserInfo.Id,
        Updated: time.Now().Unix(),
        Content: text,
    }
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
        article.Updated = article.Create
    }

    article.Status = common.ARTICLE_STATUS_PUBLISH
    err = a.saveArticle(article, tpID, tags)
    if err != nil {
        return fmt.Errorf("deal publish error: %s", err.Error())
    }

    return nil
}

/**
 @Description：编辑的文章提交
 @Param:
 @Return：
*/
func (a *ArticleEditController) Post()  {
    // 预加载
    a.AdminBase()

    var err error
    do := a.GetString("do")                     // 处理的动作
    if do == "save" {
        err = a.dealDrafts()
    } else if do == "publish" {
        err = a.dealPublish()
    } else {
        logMsg := fmt.Sprintf("input `do` type error: `%s`", do)
        a.Data["json"] = common.CreateErrResponse(common.RESP_CODE_PARAM_ERR, logMsg, nil)
        a.ServeJSON()
        return
    }

    if err != nil {
        logMsg := fmt.Sprintf("deal article post error: %s", err.Error())
        a.Data["json"] = common.CreateErrResponse(common.RESP_CODE_SYSTEM_ERR, logMsg, nil)
        a.ServeJSON()
        return
    }

    a.Data["json"] = common.CreateOkResponse(nil)
    a.ServeJSON()
    return
}

/**
 对于文章保存的处理
 */
func (a *ArticleEditController) saveArticle(article *db.Article, tpID int, tags string) error {
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