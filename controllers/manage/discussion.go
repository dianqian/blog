package manage

import (
    "blog/controllers"
)


type DiscussionController struct {
    controllers.AdminCommonCtr
}

func (this *DiscussionController) Get ()  {
    this.PreBase()

    this.Prepare()
    this.TplName = "admin/discussion.html"
    //this.TplName = "admin/example.html"
    return
}