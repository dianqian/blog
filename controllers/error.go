package controllers

import (
    "github.com/astaxie/beego"
)

type ErrorController struct {
    beego.Controller
}

/**
 404错误
 */
func (c *ErrorController) Error404() {
    c.Data["content"] = "page not found"
    c.TplName = "errors/404.html"
}

/**
 500错误
 */
func (c *ErrorController) Error501() {
    c.Data["content"] = "server error"
    c.TplName = "errors/500.html"
}


