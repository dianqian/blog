package manage

import (
    "blog/controllers/base"
)


type AdminController struct {
    base.AdminCommonCtr
}

func (this *AdminController) Get ()  {
    //this.PreBase()

    this.Prepare()
    //this.TplName = "admin/admin_layout.html"
    this.TplName = "admin/article.html"
    return
}