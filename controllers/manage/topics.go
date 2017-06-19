package manage

import (
    "blog/controllers"
)


type TopicManageController struct {
    controllers.AdminCommonCtr
}

func (this *TopicManageController) Get ()  {
    this.PreBase()

    this.Prepare()
    this.TplName = "admin/topics.html"
    //this.TplName = "admin/example.html"
    return
}