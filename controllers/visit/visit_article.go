package visit

import (
    "nest/controllers/base"
    "time"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
    "fmt"
    "strings"
    "net/url"
    "nest/models/db"
    "nest/common"
    "nest/models/html/htmlvisitor"
)

type ArticleController struct {
    base.HomeCommonCtr
}

/**
 展示文章页
 */
func (a *ArticleController) Get() {
    // 基础的执行
    a.HomeBase()

    // 数据准备处理
    htmlData := &htmlvisitor.HTMLArticleData{}

    // 解析slug
    err := a.parseURI(htmlData)
    if err != nil {
        logMsg := fmt.Sprintf("no input article url invalid '%s'.", a.Ctx.Request.RequestURI)
        logs.Error(logMsg)
        htmlData.ErrorInfo = logMsg
        a.Data["HTMLArticleData"] = htmlData
        a.TplName = "visit/visit_article.html"
        return
    }
    originAr, err := a.setArticle(htmlData)
    if err != nil {
        logMsg := fmt.Sprintf("deal article error: %s", err.Error())
        logs.Error(logMsg)
        htmlData.ErrorInfo = logMsg
        a.Data["HTMLArticleData"] = htmlData
        a.TplName = "visit/visit_article.html"
        return
    }
    err = a.setTags(htmlData, originAr)
    if err != nil {
        logMsg := fmt.Sprintf("set tags error: %s", err.Error())
        logs.Error(logMsg)
        htmlData.ErrorInfo = logMsg
        a.Data["HTMLArticleData"] = htmlData
        a.TplName = "visit/visit_article.html"
        return
    }
    err = a.setTopic(htmlData, originAr)
    if err != nil {
        logMsg := fmt.Sprintf("set topic error: %s", err.Error())
        logs.Error(logMsg)
        htmlData.ErrorInfo = logMsg
        a.Data["HTMLArticleData"] = htmlData
        a.TplName = "visit/visit_article.html"
        return
    }
    err = a.setBrother(htmlData, originAr)
    if err != nil {
        logMsg := fmt.Sprintf("set pre/next error: %s", err.Error())
        logs.Error(logMsg)
        htmlData.ErrorInfo = logMsg
        a.Data["HTMLArticleData"] = htmlData
        a.TplName = "visit/visit_article.html"
        return
    }

    htmlData.CopyRight = beego.AppConfig.String("common.copyright")
    a.Data["HTMLArticleData"] = htmlData
    a.TplName = "visit/visit_article.html"
    return
}

/**
 读取article的信息
 */
func (a *ArticleController) setArticle(ar *htmlvisitor.HTMLArticleData) (*db.Article, error) {
    originAr := &db.Article{Url: ar.ArticleInfo.Slug}
    err := originAr.Read("url")
    if err != nil {
        return nil, fmt.Errorf("no article for '%s'", ar.ArticleInfo.Slug)
    }

    ar.ArticleInfo.ID = originAr.Id
    ar.ArticleInfo.Title = originAr.Title
    ar.ArticleInfo.PublishTime = time.Unix(originAr.PublishTime, 0)
    ar.ArticleInfo.UpdateTime = time.Unix(originAr.Updated, 0)

    // markdown转换为html的文档，涉及header/content
    ar.ArticleInfo.Header, ar.ArticleInfo.Content = common.GetHeaderAndContent([]byte(originAr.Content))
    ar.ArticleInfo.Days = int(time.Now().Sub(ar.ArticleInfo.UpdateTime).Hours()) / 24

    return originAr, nil
}

/**
 设置prev/next内容
 */
func (a *ArticleController) setBrother(htmlData *htmlvisitor.HTMLArticleData, originAr *db.Article) error {

    // prev
    tmpAr := &db.Article{}
    err := tmpAr.SelectBrother(common.BROTHER_PREV, originAr.PublishTime)
    if err != nil {
        logs.Info(fmt.Sprintf("select prev brother failed: %s", err.Error()))
    } else {
        prev := htmlvisitor.ArticleBrother{
            ID: tmpAr.Id,
            Title: tmpAr.Title,
            Slug: tmpAr.Url,
            PublishTime: time.Unix(tmpAr.PublishTime, 0),
        }
        htmlData.Extend.Prev = &prev
    }

    // next
    tmpAr = &db.Article{}
    err = tmpAr.SelectBrother(common.BROTHER_NEXT, originAr.PublishTime)
    if err != nil {
        logs.Info(fmt.Sprintf("select next brother failed: %s", err.Error()))
    } else {
        next := htmlvisitor.ArticleBrother{
            ID: tmpAr.Id,
            Title: tmpAr.Title,
            Slug: tmpAr.Url,
            PublishTime: time.Unix(tmpAr.PublishTime, 0),
        }
        htmlData.Extend.Next = &next
    }

    return nil
}

/**
 todo: 设置tag内容
 */
func (a *ArticleController) setTags(ar *htmlvisitor.HTMLArticleData, originAr *db.Article) error {

    return nil
}


/**
 设置topic内容
 */
func (a *ArticleController) setTopic(htmlData *htmlvisitor.HTMLArticleData, originAr *db.Article) error {

    tp := &db.Topic{Id: originAr.Id}
    err := tp.Read("id")
    if err != nil {
        return fmt.Errorf("the id='%d no find topic object: %s", originAr.Id, err.Error())
    }

    arTp := &db.ArticleTopic{TopicId: tp.Id}
    arTps, err := arTp.SelectByTp(tp.Id)
    if err != nil {
        return fmt.Errorf("find by topic id failed: %s", err.Error())
    }

    topicArticle := htmlvisitor.TopicArticle{ID: tp.Id, Name: tp.Name}
    for _, item := range arTps {
        tmpAr := &db.Article{Id: item.ArticleId}
        err = tmpAr.Read("id")
        if err != nil {
            logs.Error(fmt.Sprintf("find article by id'%d' failed: %s", item.ArticleId, err.Error()))
            continue
        }
        if tmpAr.Status != common.ARTICLE_STATUS_PUBLISH {
            continue
        }

        topicArticle.Articles = append(topicArticle.Articles, htmlvisitor.ArticleBrother{
            ID: tmpAr.Id, Title: tmpAr.Title, Slug: tmpAr.Url, PublishTime: time.Unix(tmpAr.PublishTime, 0),
        })
    }

    htmlData.Extend.Topic = topicArticle

    return nil
}

/**
 解析URI得到article的slug的信息
 */
func (a *ArticleController) parseURI(ar *htmlvisitor.HTMLArticleData) error {

    u, err := url.Parse(a.Ctx.Request.URL.EscapedPath())
    if err != nil {
        return err
    }
    inputURI := u.Path

    if strings.HasPrefix(inputURI, common.POST_TAG) != true ||
        strings.HasSuffix(inputURI, common.HTML_TAG) != true {
        return fmt.Errorf("input uri invalid: %s", inputURI)
    }

    begin := strings.Index(inputURI, common.POST_TAG) + len(common.POST_TAG)
    end := strings.LastIndex(inputURI, common.HTML_TAG)
    if begin >= end {
        return fmt.Errorf("input uri invalid: %s", inputURI)
    }

    ar.ArticleInfo.Slug = inputURI[begin: end]

    return nil
}