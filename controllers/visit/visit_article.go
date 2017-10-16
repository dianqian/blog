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
)

/**
 文章展示的相关信息
 */
type ArticleInfo struct {
    ID              int                     // 文章ID
    Title           string                  // 标题
    Author          string                  // 作者
    Header          string                  // 开头，实际用于文件目录list

    Tags            []string                // 标签
    Topic           TopicArticle            // 专题

    Content         string                  // 文章内容

    PublishTime     time.Time               // 发布时间
    UpdateTime      time.Time               // 更新时间

    CommentCount    int                     // 评论数

    Slug            string                  // 略缩信息
    Days            int

    Prev            *ArticleBrother          // 上一篇
    Next            *ArticleBrother          // 下一篇
}

type ArticleBrother struct {
    ID              int
    Title           string
    Slug            string
    PublishTime     time.Time
}

type TopicArticle struct {
    ID              int                     // topic id值
    Name            string                  // topic name值
    Articles        []ArticleBrother        // 同一个topic下文章信息
}

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
    article := &ArticleInfo{}

    // 解析slug
    err := a.parseURI(article)
    if err != nil {
        logs.Error(fmt.Sprintf("no input article url invalid '%s'.", a.Ctx.Request.RequestURI))
    }
    originAr, err := a.setArticle(article)
    if err != nil {
        logs.Error(fmt.Sprintf("deal article error: %s", err.Error()))
    }
    err = a.setTags(article, originAr)
    if err != nil {
        logs.Error(fmt.Sprintf("set tags error: %s", err.Error()))
    }
    err = a.setTopic(article, originAr)
    if err != nil {
        logs.Error(fmt.Sprintf("set topic error: %s", err.Error()))
    }
    err = a.setBrother(article, originAr)
    if err != nil {
        logs.Error(fmt.Sprintf("set pre/next error: %s", err.Error()))
    }

    a.Data["Article"] = article

    a.Data["Copyright"] = beego.AppConfig.String("common.copyright")
    a.TplName = "visit/visit_article.html"
    return
}

/**
 读取article的信息
 */
func (a *ArticleController) setArticle(ar *ArticleInfo) (*db.Article, error) {
    originAr := &db.Article{Url: ar.Slug}
    err := originAr.Read("url")
    if err != nil {
        return nil, fmt.Errorf("no article for '%s'", ar.Slug)
    }

    ar.ID = originAr.Id
    ar.Title = originAr.Title
    ar.PublishTime = time.Unix(originAr.PublishTime, 0)
    ar.UpdateTime = time.Unix(originAr.Updated, 0)

    // markdown转换为html的文档，涉及header/content
    ar.Header, ar.Content = common.GetHeaderAndContent([]byte(originAr.Content))
    ar.Days = int(time.Now().Sub(ar.UpdateTime).Hours()) / 24

    ar.CommentCount = 11

    return originAr, nil
}

/**
 设置prev/next内容
 */
func (a *ArticleController) setBrother(ar *ArticleInfo, originAr *db.Article) error {

    // prev
    tmpAr := &db.Article{}
    err := tmpAr.SelectBrother(common.BROTHER_PREV, originAr.PublishTime)
    if err != nil {
        logs.Info(fmt.Sprintf("select prev brother failed: %s", err.Error()))
    } else {
        prev := ArticleBrother{
            ID: tmpAr.Id,
            Title: tmpAr.Title,
            Slug: tmpAr.Url,
            PublishTime: time.Unix(tmpAr.PublishTime, 0),
        }
        ar.Prev = &prev
    }

    // next
    tmpAr = &db.Article{}
    err = tmpAr.SelectBrother(common.BROTHER_NEXT, originAr.PublishTime)
    if err != nil {
        logs.Info(fmt.Sprintf("select next brother failed: %s", err.Error()))
    } else {
        next := ArticleBrother{
            ID: tmpAr.Id,
            Title: tmpAr.Title,
            Slug: tmpAr.Url,
            PublishTime: time.Unix(tmpAr.PublishTime, 0),
        }
        ar.Next = &next
    }

    return nil
}

/**
 todo: 设置tag内容
 */
func (a *ArticleController) setTags(ar *ArticleInfo, originAr *db.Article) error {

    return nil
}


/**
 设置topic内容
 */
func (a *ArticleController) setTopic(ar *ArticleInfo, originAr *db.Article) error {

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

    topicArticle := TopicArticle{ID: tp.Id, Name: tp.Name}
    for _, item := range arTps {
        tmpAr := &db.Article{Id: item.ArticleId}
        err = tmpAr.Read("id")
        if err != nil {
            logs.Error(fmt.Sprintf("find article by id'%d' failed: %s", item.ArticleId, err.Error()))
        }
        topicArticle.Articles = append(topicArticle.Articles, ArticleBrother{
            ID: tmpAr.Id, Title: tmpAr.Title, Slug: tmpAr.Url, PublishTime: time.Unix(tmpAr.PublishTime, 0),
        })
    }

    ar.Topic = topicArticle

    return nil
}

/**
 解析URI得到article的slug的信息
 */
func (a *ArticleController) parseURI(ar *ArticleInfo) error {

    u, err := url.Parse(a.Ctx.Request.URL.EscapedPath())
    if err != nil {
        return err
    }
    inputURI := u.Path

    if strings.HasPrefix(inputURI, "/post/") != true ||
        strings.HasSuffix(inputURI, ".html") != true {
        return fmt.Errorf("input uri invalid: %s", inputURI)
    }

    result := strings.TrimLeft(inputURI, "/post/")
    result = strings.TrimRight(result, ".html")

    ar.Slug = result

    return nil
}