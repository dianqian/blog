package admin

import (
    "nest/controllers/base"
    "nest/models/html/htmlvisitor"
    "nest/models/db"
    "nest/models/html/htmladmin"
    "nest/common"
    "fmt"
    "github.com/astaxie/beego/logs"
    "github.com/astaxie/beego"
)

// @Description：预览加载的实现
type PreviewController struct {
    base.AdminCommonCtr
}

/**
 @Description：获取文章预览的效果
 @Param:
 @Return：
*/
func (p *PreviewController) Get()  {
    // 针对preview的admin处理
    p.AdminForPreview()

    htmlData := htmladmin.HTMLArticlePreview{}

    articleID, err := p.GetInt("id")
    if err != nil {
        // id输入失败，则表示有问题
        logMsg := fmt.Sprintf("input article id failed: %s", err.Error())
        logs.Error(logMsg)
        htmlData.ErrorInfo = logMsg
        p.Data["HTMLArticlePreview"] = htmlData
        p.TplName = "admin/admin_preview.html"
        return
    }

    article, err := p.dealVisitNavbar(articleID)
    if err != nil {
        logMsg := fmt.Sprintf("deal preview nav failed: %s", err.Error())
        logs.Error(logMsg)
        htmlData.ErrorInfo = logMsg
        p.Data["HTMLArticlePreview"] = htmlData
        p.TplName = "admin/admin_preview.html"
        return
    }
    htmlData.ArticleInfo.ID = article.Id
    htmlData.ArticleInfo.Title = article.Title
    htmlData.ArticleInfo.Slug = article.Url
    htmlData.ArticleInfo.Header, htmlData.ArticleInfo.Content = common.GetHeaderAndContent([]byte(article.Content))

    p.Data["HTMLArticlePreview"] = htmlData
    p.TplName = "admin/admin_preview.html"
    return
}

/**
 @Description：处理预览的右侧导航栏
 @Param:
 @Return：
*/
func (p *PreviewController) dealVisitNavbar(id int) (*db.Article, error) {
    var visitHTML htmlvisitor.HTMLVisitData

    originAr := &db.Article{Id: id}
    err := originAr.Read("id")
    if err != nil {
        return nil, err
    }

    u := &db.User{Id: originAr.Author}
    err = u.Read("id")
    if err != nil {
        return nil, err
    }
    visitHTML.Icon = u.Avatar
    visitHTML.Name = u.NickName
    visitHTML.Description = u.Signature
    if visitHTML.Icon == "" {
        visitHTML.Icon = beego.AppConfig.String("default_config.icon")
    }
    if visitHTML.Name == "" {
        visitHTML.Name = beego.AppConfig.String("default_config.name")
    }
    if visitHTML.Description == "" {
        visitHTML.Description = beego.AppConfig.String("default_config.description")
    }

    p.Data["HTMLVisitData"] = visitHTML
    p.Layout = "visit/visit_layout.html"

    return originAr, nil
}