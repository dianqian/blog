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
	// 首页文章列表
	//this.Layout = "blog/home_layout.html"
	this.TplName = "blog/article_list.html"
	return
}

