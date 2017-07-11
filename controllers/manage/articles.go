package manage

import (
    "blog/controllers/base"
)


type ArticleManageController struct {
    base.AdminCommonCtr
}

func (this *ArticleManageController) Get ()  {

    this.AdminBase()

    this.TplName = "admin/articles.html"
    //this.TplName = "admin/example.html"
    return
}