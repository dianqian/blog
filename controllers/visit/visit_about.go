package visit

import (
    "blog/controllers/base"
    "blog/models/db"
    "blog/common"
    "github.com/astaxie/beego/logs"
    "fmt"
)

type AboutArticle struct {
    ID          int
    Title       string
    Header      string              // 头目录
    Content     string
    Slug        string              // 略缩字
}

type About struct {
    base.HomeCommonCtr
}

/**
 about页的展示
 */
func (a *About) Get() {
    // 基础的执行
    a.HomeBase()

    // 准备数据
    ar := &db.Article{Url: "about"}
    err := ar.Read("url")
    if err != nil {
        logs.Error(fmt.Sprintf("no found about html, error: %s", err.Error()))
    } else {
        about := AboutArticle{
            ID: ar.Id, Title: ar.Title, Slug: ar.Url,
        }
        about.Header, about.Content = common.GetHeaderAndContent([]byte(ar.Content))
        a.Data["About"] = about
    }

    a.TplName = "visit/visit_about.html"
    return
}
