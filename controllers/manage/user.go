package manage

import (
    "blog/controllers/base"
)


type UserManageController struct {
    base.AdminCommonCtr
}

func (this *UserManageController) Get ()  {
    //this.PreBase()

    this.Prepare()
    this.TplName = "admin/user.html"
    //this.TplName = "admin/example.html"
    return
}