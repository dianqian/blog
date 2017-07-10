package manage

import (
    "blog/controllers/base"
)


type GeneralSettingController struct {
    base.AdminCommonCtr
}

func (this *GeneralSettingController) Get ()  {
    //this.PreBase()

    this.Prepare()
    this.TplName = "admin/general.html"
    //this.TplName = "admin/example.html"
    return
}