package controllers

import (

)

type HomeController struct {
	HomeCommonCtr
}

func (this *HomeController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"

	this.PreBase()
	this.Prepare()
	this.TplName = "blog/article_list.html"
	return
}

