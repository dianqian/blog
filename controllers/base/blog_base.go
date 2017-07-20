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
func (h *HomeCommonCtr) HomeBase() {
    h.CommonBase()

    //this.Data["Socials"] = models.FindAllSocial()

    h.Layout = "visitor/home_layout.html"
    logs.Debug("HomeCommonCtr Prepare......")
}
