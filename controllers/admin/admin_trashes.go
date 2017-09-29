package admin

import (
    "blog/controllers/base"
)


type TrashManageController struct {
    base.AdminCommonCtr
}

func (this *TrashManageController) Get ()  {

    this.AdminBase()

    this.TplName = "admin/admin_trashes.html"
    return
}