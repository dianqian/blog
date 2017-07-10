package manage

import (
    "blog/controllers/base"
)


type DiscussionController struct {
    base.AdminCommonCtr
}

func (this *DiscussionController) Get ()  {

    this.Prepare()
    this.TplName = "admin/discussion.html"
    //this.TplName = "admin/example.html"
    return
}