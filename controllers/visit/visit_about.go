package visit

import (
    "nest/controllers/base"
    "nest/models/db"
    "nest/common"
    "github.com/astaxie/beego/logs"
    "fmt"
    "nest/models/html/htmlvisitor"
)

type About struct {
    base.HomeCommonCtr
}

/**
 about页的展示
 */
func (a *About) Get() {
    // 基础的执行
    a.HomeBase()

    htmlData := htmlvisitor.HTMLAbout{}

    // 准备数据
    ar := &db.Article{Url: "about"}
    err := ar.Read("url")
    if err != nil {
        logMsg := fmt.Sprintf("no found about html, error: %s", err.Error())
        logs.Error(logMsg)
        a.Data["HTMLAbout"] = htmlData
        a.TplName = "visit/visit_about.html"
        return
    }

    about := htmlvisitor.AboutArticle{
        ID: ar.Id, Title: ar.Title, Slug: ar.Url,
    }
    about.Header, about.Content = common.GetHeaderAndContent([]byte(ar.Content))
    htmlData.About = about

    a.Data["HTMLAbout"] = htmlData
    a.TplName = "visit/visit_about.html"
    return
}
