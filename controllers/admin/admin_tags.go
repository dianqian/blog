package admin

import (
    "blog/controllers/base"
    "blog/models/db"
    "github.com/astaxie/beego/logs"
    "fmt"
)


type TagManageController struct {
    base.AdminCommonCtr
}

func (tgMCtrl *TagManageController) Get ()  {
    // 基础执行
    tgMCtrl.AdminBase()

    // 准备数据
    tg := new(db.Tag)
    tgs, err := tg.SelectAll()
    if err != nil {
        // todo: select出错
        logs.Error(fmt.Sprintf("tag select all failed: %s.", err.Error()))
    } else {
        tgMCtrl.Data["TagList"] = tgs
    }

    tgMCtrl.TplName = "admin/admin_tags.html"
    return
}