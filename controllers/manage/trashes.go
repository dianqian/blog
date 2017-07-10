package manage

import (
    "blog/controllers/base"
)


type TrashManageController struct {
    base.AdminCommonCtr
}

func (this *TrashManageController) Get ()  {
    //this.PreBase()

    this.Prepare()
    this.TplName = "admin/trashes.html"
    //this.TplName = "admin/example.html"
    return
}