package admin

import (
    "nest/controllers/base"
)


type DiscussionController struct {
    base.AdminCommonCtr
}

func (d *DiscussionController) Get ()  {

    d.Prepare()
    d.TplName = "admin/admin_discussion.html"
    //this.TplName = "admin/example.html"
    return
}