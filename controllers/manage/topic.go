package manage

import (
    "blog/controllers"
)


type TopicEditController struct {
    controllers.AdminCommonCtr
}

func (this *TopicEditController) Get ()  {
    this.PreBase()

    this.Prepare()
    this.TplName = "admin/topic.html"
    //this.TplName = "admin/example.html"
    return
}