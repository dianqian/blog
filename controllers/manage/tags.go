package manage

import (
    "blog/controllers/base"
)


type TagManageController struct {
    base.AdminCommonCtr
}

func (this *TagManageController) Get ()  {
    //this.PreBase()

    this.Prepare()
    this.TplName = "admin/tags.html"
    //this.TplName = "admin/example.html"
    return
}