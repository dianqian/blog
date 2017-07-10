package routers

import (
	"blog/controllers"
	"blog/controllers/manage"
	"blog/controllers/visitor"
	"github.com/astaxie/beego"
)

func init() {
	/*********************************
	 博客页
	 */
	// 首页
    beego.Router("/", &visitor.HomeController{})

	// 文章显示
	beego.Router("/article", &visitor.ArticleController{})

	// user相关的页面
	//

	/*********************************
	 todo: administrator相关的内容
	       admin相关的页面
	 */
	// admin login的页面
	beego.Router("/admin/login", &manage.LoginController{})
	beego.Router("/admin/logout", &manage.LogoutController{})

	// 管理首页
	beego.Router("/admin", &manage.AdminController{})

	// 管理员user信息管理
	beego.Router("/admin/user", &manage.UserManageController{})

	// 编辑文档页
	beego.Router("/admin/edit-article", &manage.ArticleEditController{})
	// 管理文档页
	beego.Router("/admin/manage-articles", &manage.ArticleManageController{})
	// 草稿箱管理
	beego.Router("/admin/manage-drafts", &manage.DraftManageController{})
	// 回收箱管理
	beego.Router("/admin/manage-trashes", &manage.TrashManageController{})

	// 编辑专题
	beego.Router("/admin/edit-topic", &manage.TopicEditController{})
	// 管理专题topic
	beego.Router("/admin/manage-topics", &manage.TopicManageController{})

	// 管理标签tags
	beego.Router("/admin/manage-tags", &manage.TagManageController{})

	// todo: 待确认，具体是否需要
	// 常规基本设置页
	beego.Router("/admin/general-setting", &manage.GeneralSettingController{})
	// 评论相关
	beego.Router("/admin/discussion", &manage.DiscussionController{})


	/*********************************
	 todo: 其他错误、异常处理相关
	 */
	// 404错误页
	beego.Router("/404.html", &controllers.Error404Controller{})
	// 500错误页
	beego.Router("/500.html", &controllers.Error500Controller{})

	// test
	beego.Router("/test/", &controllers.TestController{})

}
