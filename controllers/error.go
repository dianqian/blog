package controllers

import (
    "github.com/astaxie/beego"
)

/**
 404错误的控制器
 */
type Error404Controller struct {
    beego.Controller
}

func (this *Error404Controller) Get() {
    this.TplName = "errors/404.html"
}

/**
 500错误的控制器
 */
type Error500Controller struct {
    beego.Controller
}

func (this *Error500Controller) Get()  {
    this.TplName = "errors/500.html"
}

