package manage

import (
    "blog/controllers/base"
)


type TopicEditController struct {
    base.AdminCommonCtr
}

func (this *TopicEditController) Get ()  {
    //this.PreBase()

    this.Prepare()
    this.TplName = "admin/topic.html"
    //this.TplName = "admin/example.html"
    return
}