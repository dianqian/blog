package visitor

import (
	"github.com/astaxie/beego/logs"
	"blog/controllers/base"
	"fmt"
	"blog/models"
)

type HomeController struct {
	base.HomeCommonCtr
}

func (h *HomeController) Get() {
	h.HomeBase()

	h.Data["Website"] = "beego.me"
	h.Data["Email"] = "astaxie@gmail.com"

	// 首页文章列表
	a := new(models.Article)
	_, err := a.Count(100)
	if err != nil {
		// todo: 错误处理
		logs.Error(fmt.Sprintf("home article count failed: %s", err.Error()))
	}
	//cnt = cnt

	// 指定页码
	_, err = h.GetInt("pn")
	if err != nil {

	}
	//pn = pn



	//this.Layout = "blog/home_layout.html"
	h.Data["Prev"] = 1
	h.Data["Next"] = 2

	h.TplName = "visitor/home.html"
	logs.Debug("HomeController Get......")
	return
}

