package manage

import (
    "blog/controllers"
)


type ArticleManageController struct {
    controllers.AdminCommonCtr
}

func (this *ArticleManageController) Get ()  {
    this.PreBase()

    this.Prepare()
    this.TplName = "admin/articles.html"
    //this.TplName = "admin/example.html"
    return
}