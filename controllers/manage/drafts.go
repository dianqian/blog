package manage

import (
    "blog/controllers"
)


type DraftManageController struct {
    controllers.AdminCommonCtr
}

func (this *DraftManageController) Get ()  {
    this.PreBase()

    this.Prepare()
    this.TplName = "admin/drafts.html"
    //this.TplName = "admin/example.html"
    return
}