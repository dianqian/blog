package manage

import (
    "blog/controllers/base"
    "blog/common"
    "blog/models"
    "github.com/astaxie/beego/logs"
    "fmt"
)


/**
 article相关的topic
 */
type TopicBrief struct {
    ID          int
    Name        string
}


type ArticleManageController struct {
    base.AdminCommonCtr
}

/**
 todo: 参考(https://beego.me/docs/mvc/view/page.md)使用分页模式操作
 todo：也可以参考：https://beego.me/docs/utils/page.md
 */
func (a *ArticleManageController) Get ()  {

    a.AdminBase()

    article := new(models.Article)
    articles, err := article.SelectByStatus(common.ARTICLE_STATUS_DRAFT)
    if err != nil {
        logs.Error(fmt.Sprintf("list draft article failed: %s", err.Error()))
    }
    // todo: 待实现的内容
    articles = articles

    tp := new(models.Topic)
    tps, err := tp.SelectAll()
    if err != nil {
        logs.Error(fmt.Sprintf("list all topics failed: %s", err.Error()))
    }
    var tpbs []*TopicBrief
    for _, item := range tps {
        tpbs = append(tpbs, &TopicBrief{ID: item.Id, Name: item.Name})
    }
    a.Data["Series"] = tpbs

    title := ArticleTitleInfo{}
    publishTitles, err := title.GetTitleInfos(common.ARTICLE_STATUS_PUBLISH)
    if err != nil {
        logs.Error(fmt.Sprintf("draft GetTitleInfos failed: %s", err.Error()))
        a.TplName = "admin/articles.html"
        return
    }

    a.Data["List"] = publishTitles
    a.TplName = "admin/articles.html"
    return
}