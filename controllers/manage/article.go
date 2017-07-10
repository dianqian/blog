package manage

import (
    "blog/controllers/base"
)


type ArticleEditController struct {
    base.AdminCommonCtr
}

func (this *ArticleEditController) Get ()  {

    this.Prepare()
    this.TplName = "admin/article.html"
    //this.TplName = "admin/example.html"
    return
}