package manage

import (
    "blog/controllers/base"
)


type ArticleManageController struct {
    base.AdminCommonCtr
}

/**
 todo: 参考(https://beego.me/docs/mvc/view/page.md)使用分页模式操作
 todo：也可以参考：https://beego.me/docs/utils/page.md
 */
func (this *ArticleManageController) Get ()  {

    this.AdminBase()

    this.TplName = "admin/articles.html"
    //this.TplName = "admin/example.html"
    return
}