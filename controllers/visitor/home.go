package visitor

import (
	"github.com/astaxie/beego/logs"
	"blog/controllers/base"
)

type HomeController struct {
	base.HomeCommonCtr
}

func (this *HomeController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"

	// 首页文章列表
	//this.Layout = "blog/home_layout.html"
	this.TplName = "visitor/article_list.html"
	logs.Debug("HomeController Get......")
	return
}

