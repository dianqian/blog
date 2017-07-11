package manage

import (
    "blog/controllers/base"
)


type TrashManageController struct {
    base.AdminCommonCtr
}

func (this *TrashManageController) Get ()  {

    this.AdminBase()

    this.TplName = "admin/trashes.html"
    return
}