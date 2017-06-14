package manage

import (
    "blog/controllers"
)


type AdminController struct {
    controllers.AdminCommonCtr
}

func (this *AdminController) Get ()  {
    this.PreBase()

    //this.Prepare()
    this.TplName = "admin/admin_layout.html"
    this.TplName = "admin/editor.html"
    return
}