package visit

import (
    "nest/controllers/base"
    "nest/models/db"
    "github.com/astaxie/beego/logs"
    "fmt"
    "nest/common"
)

type RollInfo struct {
    ID          int
    Title       string
    Header      string
    Content     string
}

/**
 友链展示
 */
type BlogRoll struct {
    base.HomeCommonCtr
}

func (b *BlogRoll) Get()  {
    // 基础的执行
    b.HomeBase()

    // 准备数据
    ar := &db.Article{Url: "blog_roll"}
    err := ar.Read("url")
    if err != nil {
        logs.Error(fmt.Sprintf("no found blog roll html, error: %s", err.Error()))
    } else {
        roll := RollInfo{ID: ar.Id, Title: ar.Title}
        roll.Header, roll.Content = common.GetHeaderAndContent([]byte(ar.Content))

        b.Data["Article"] = roll
    }

    b.TplName = "visit/visit_roll.html"
    return
}
