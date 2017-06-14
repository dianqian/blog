package controllers

import (
    "github.com/astaxie/beego"
    "time"
    "blog/models"
)

/**
 整个站点的功能类
 */
type CommonCtr struct {
    beego.Controller
}

func (this *CommonCtr) PreBase() {
    this.Data["Title"] = "我的蘑菇鸡"

    /* 获取 */
    this.Data["CopyTime"] = time.Now().Year()
}

/**
 业务Controller的基础公共类，利用继承的特性
 */
type HomeCommonCtr struct {
    CommonCtr
}

/**
 预处理
 */
func (this *HomeCommonCtr) Prepare() {
    this.Data["Socials"] = models.FindAllSocial()

    this.Layout = "blog/home_layout.html"
}

/**
 admin管理的Controller的公共基础类
 */
type AdminCommonCtr struct {
    CommonCtr
}

func (this *AdminCommonCtr) Prepare() {
    this.Layout = "admin/admin_layout.html"
}