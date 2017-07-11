package manage

import (
    "blog/controllers/base"
)


type TagManageController struct {
    base.AdminCommonCtr
}

func (this *TagManageController) Get ()  {

    this.AdminBase()

    this.TplName = "admin/tags.html"
    return
}