package manage

import (
    "blog/controllers/base"
)


type TopicManageController struct {
    base.AdminCommonCtr
}

func (this *TopicManageController) Get ()  {
    //this.PreBase()

    this.Prepare()
    this.TplName = "admin/topics.html"
    //this.TplName = "admin/example.html"
    return
}