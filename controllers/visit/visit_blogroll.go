package visit

import (
    "nest/controllers/base"
    "nest/models/db"
    "github.com/astaxie/beego/logs"
    "fmt"
    "nest/common"
    "nest/models/html/htmlvisitor"
)



/**
 友链展示
 */
type BlogRoll struct {
    base.HomeCommonCtr
}

func (b *BlogRoll) Get()  {
    // 基础的执行
    b.HomeBase()

    htmlData := htmlvisitor.HTMLRollInfo{}

    // 准备数据
    ar := &db.Article{Url: "blog_roll"}
    err := ar.Read("url")
    if err != nil {
        logMsg := fmt.Sprintf("no found blog roll html, error: %s", err.Error())
        logs.Error(logMsg)
        htmlData.ErrorInfo = logMsg
        b.Data["HTMLRollInfo"] = htmlData
        b.TplName = "visit/visit_roll.html"
        return
    }

    roll := htmlvisitor.RollInfo{ID: ar.Id, Title: ar.Title}
    roll.Header, roll.Content = common.GetHeaderAndContent([]byte(ar.Content))
    htmlData.Roll = roll

    b.Data["HTMLRollInfo"] = htmlData
    b.TplName = "visit/visit_roll.html"
    return
}
