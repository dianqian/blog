package manage

import (
    "blog/controllers"
)


type TrashManageController struct {
    controllers.AdminCommonCtr
}

func (this *TrashManageController) Get ()  {
    this.PreBase()

    this.Prepare()
    this.TplName = "admin/trashes.html"
    //this.TplName = "admin/example.html"
    return
}