// @Description: admin相关的注册
package routers

import (
    "github.com/astaxie/beego"
    "nest/controllers/admin"
)

func init()  {
    /*********************************
     todo: administrator相关的内容
           admin相关的页面
     */
    // admin login的页面
    beego.Router("/admin/login.html", &admin.LoginController{})
    beego.Router("/admin/logout.html", &admin.LogoutController{})

    // 管理首页
    beego.Router("/admin", &admin.PublishManageController{})
    beego.Router("/admin/manager.html", &admin.BlogMgController{})

    // 管理员user信息管理
    beego.Router("/admin/user.html", &admin.UserManageController{})
    // todo: 配置相关的几个待实现
    beego.Router("/admin/user/account", &admin.UserManageController{}, "post:ChangeAccountInfo")
    beego.Router("/admin/user/avatar", &admin.UserManageController{}, "post:ChangeAvatarInfo")
    beego.Router("/admin/user/password", &admin.UserManageController{}, "post:ChangePassword")

    /*** 文档管理 ***/
    // 编辑文档页
    beego.Router("/admin/article.html", &admin.ArticleEditController{})
    beego.Router("/admin/preview.html", &admin.PreviewController{})
    // 管理文档页
    beego.Router("/admin/articles.html", &admin.PublishManageController{})
    // 草稿箱管理
    beego.Router("/admin/drafts.html", &admin.DraftManageController{})
    // 回收箱管理
    beego.Router("/admin/trashes.html", &admin.TrashManageController{})

    /*** 专题相关 ***/
    // 编辑专题
    beego.Router("/admin/topic.html", &admin.TopicEditController{})
    // 管理专题topic
    beego.Router("/admin/topics.html", &admin.TopicManageController{})

    /*** 标签相关 ***/
    // 管理标签tags
    beego.Router("/admin/tags.html", &admin.TagManageController{})
    beego.Router("/admin/tags/list", &admin.TagManageController{}, "get:TagsList")

    // todo: 待确认，具体是否需要
    beego.Router("/admin/manager.html", &admin.BlogMgController{})
    beego.Router("/admin/manager/icon", &admin.BlogMgController{}, "post:ChangeIcon")
    beego.Router("/admin/manager/info", &admin.BlogMgController{}, "post:ChangeInfo")
    // 常规基本设置页
    beego.Router("/admin/general-setting", &admin.GeneralSettingController{})
    // 评论相关
    beego.Router("/admin/discussion.html", &admin.DiscussionController{})
}
