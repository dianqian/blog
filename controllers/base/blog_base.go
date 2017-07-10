package base

import (
    "github.com/astaxie/beego/logs"
)

/**
 blog业务Controller的基础公共类，利用继承的特性
 */
type HomeCommonCtr struct {
    CommonCtr
}

/**
 预处理
 */
func (this *HomeCommonCtr) Prepare() {
    this.CommonBase()

    //this.Data["Socials"] = models.FindAllSocial()

    this.Layout = "visitor/home_layout.html"
    logs.Debug("HomeCommonCtr Prepare......")
}
