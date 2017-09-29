package base

import (
    "github.com/astaxie/beego/logs"
    "github.com/astaxie/beego"
    "blog/models/html/htmlvisitor"
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

    var visitHTML htmlvisitor.HTMLVisitData
    visitHTML.Version = beego.AppConfig.String("common.version")
    visitHTML.Domain = beego.AppConfig.String(beego.BConfig.RunMode + "::domain")           // 域名
    h.Data["HTMLVisitData"] = visitHTML

    h.Data["Version"] = beego.AppConfig.String("common.version")
    h.Data["Domain"] = beego.AppConfig.String(beego.BConfig.RunMode + "::domain")           // 域名

    h.Layout = "visit/visit_layout.html"
    logs.Debug("HomeCommonCtr Prepare......")
}
