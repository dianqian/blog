package manage

import (
    "blog/controllers"
)


type TagManageController struct {
    controllers.AdminCommonCtr
}

func (this *TagManageController) Get ()  {
    this.PreBase()

    this.Prepare()
    this.TplName = "admin/tags.html"
    //this.TplName = "admin/example.html"
    return
}