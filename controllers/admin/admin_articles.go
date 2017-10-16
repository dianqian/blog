package admin

import (
    "nest/controllers/base"
    "nest/common"
    "nest/models/db"
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
    // 基础项
    a.AdminBase()

    // 根据status选择文章
    article := new(db.Article)
    articles, err := article.SelectByStatus(common.ARTICLE_STATUS_DRAFT)
    if err != nil {
        logs.Error(fmt.Sprintf("list draft article failed: %s", err.Error()))
    }
    // todo: 待实现的内容
    articles = articles

    tp := new(db.Topic)
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
        a.TplName = "admin/admin_articles.html"
        return
    }

    a.Data["List"] = publishTitles
    a.TplName = "admin/admin_articles.html"
    return
}