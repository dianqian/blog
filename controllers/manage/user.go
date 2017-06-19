package manage

import (
    "blog/controllers"
)


type UserManageController struct {
    controllers.AdminCommonCtr
}

func (this *UserManageController) Get ()  {
    this.PreBase()

    this.Prepare()
    this.TplName = "admin/user.html"
    //this.TplName = "admin/example.html"
    return
}