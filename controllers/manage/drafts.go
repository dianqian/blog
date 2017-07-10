package manage

import (
    "blog/controllers/base"
)


type DraftManageController struct {
    base.AdminCommonCtr
}

func (this *DraftManageController) Get ()  {

    this.Prepare()
    this.TplName = "admin/drafts.html"
    //this.TplName = "admin/example.html"
    return
}