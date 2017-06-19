package manage

import (
    "blog/controllers"
)


type ArticleEditController struct {
    controllers.AdminCommonCtr
}

func (this *ArticleEditController) Get ()  {
    this.PreBase()

    this.Prepare()
    this.TplName = "admin/article.html"
    //this.TplName = "admin/example.html"
    return
}