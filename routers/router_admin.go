// @Description: admin相关的注册
package routers

import (
    "github.com/astaxie/beego"
    "blog/controllers/admin"
)

func init()  {
    /*********************************
     todo: administrator相关的内容
           admin相关的页面
     */
    // admin login的页面
    beego.Router("/admin/admin_login.html", &admin.LoginController{})
    beego.Router("/admin/logout.html", &admin.LogoutController{})

    // 管理首页
    beego.Router("/admin/admin_manager.html", &admin.BlogMgController{})

    // 管理员user信息管理
    beego.Router("/admin/user", &admin.UserManageController{})
    // todo: 配置相关的几个待实现
    beego.Router("/admin/user/account", &admin.UserManageController{}, "post:ChangeAccountInfo")
    beego.Router("/admin/user/avatar", &admin.UserManageController{}, "post:ChangeAvatarInfo")
    beego.Router("/admin/user/password", &admin.UserManageController{}, "post:ChangePassword")

    // 编辑文档页
    beego.Router("/admin/edit-article", &admin.ArticleEditController{})
    // 管理文档页
    beego.Router("/admin/manage-articles", &admin.ArticleManageController{})
    // 草稿箱管理
    beego.Router("/admin/manage-drafts", &admin.DraftManageController{})
    // 回收箱管理
    beego.Router("/admin/manage-trashes", &admin.TrashManageController{})

    // 编辑专题
    beego.Router("/admin/edit-topic", &admin.TopicEditController{})
    // 管理专题topic
    beego.Router("/admin/manage-topics", &admin.TopicManageController{})

    // 管理标签tags
    beego.Router("/admin/manage-tags", &admin.TagManageController{})

    // todo: 待确认，具体是否需要
    beego.Router("/admin/blog", &admin.BlogMgController{})
    beego.Router("/admin/blog/icon", &admin.BlogMgController{}, "post:ChangeIcon")
    beego.Router("/admin/blog/info", &admin.BlogMgController{}, "post:ChangeInfo")
    // 常规基本设置页
    beego.Router("/admin/general-setting", &admin.GeneralSettingController{})
    // 评论相关
    beego.Router("/admin/discussion", &admin.DiscussionController{})
}
