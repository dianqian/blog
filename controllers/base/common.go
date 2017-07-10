package base

import (
    "github.com/astaxie/beego"
    "time"
    "github.com/astaxie/beego/logs"
)

/**
 整个站点的功能类
 */
type CommonCtr struct {
    beego.Controller
}

func (this *CommonCtr) CommonBase() {
    this.Data["Title"] = "我的蘑菇鸡"

    /* 获取 */
    this.Data["CopyTime"] = time.Now().Year()
    logs.Debug("CommonCtr CommonBase......")
}
