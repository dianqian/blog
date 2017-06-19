package manage

import (
    "blog/controllers"
)


type GeneralSettingController struct {
    controllers.AdminCommonCtr
}

func (this *GeneralSettingController) Get ()  {
    this.PreBase()

    this.Prepare()
    this.TplName = "admin/general.html"
    //this.TplName = "admin/example.html"
    return
}