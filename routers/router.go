package routers

import (
	"blog/controllers"
	"blog/controllers/manage"
	"github.com/astaxie/beego"
)

func init() {
	// 首页
    beego.Router("/", &controllers.HomeController{})

	// 文章显示
	beego.Router("/article", &controllers.ArticleController{})

	// user相关的页面
	//

	// admin相关的页面
	// 管理首页
	beego.Router("/admin", &manage.AdminController{})
	// 编辑文档页
	beego.Router("/admin/editor", &manage.EditorController{})

	// 404错误页
	beego.Router("/404.html", &controllers.Error404Controller{})
	// 500错误页
	beego.Router("/500.html", &controllers.Error500Controller{})

}
