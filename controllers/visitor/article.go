package visitor

import (
    "blog/controllers/base"
)

type ArticleController struct {
    base.HomeCommonCtr
}

func (this *ArticleController) Get() {
    this.Data["IsFalse"] = true

    //this.PreBase()
    this.Prepare()
    this.TplName = "blog/article.html"
    return
}
