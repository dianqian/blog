package controllers

import (
    
)

type ArticleController struct {
    HomeCommonCtr
}

func (this *ArticleController) Get() {
    this.Data["IsFalse"] = true

    this.PreBase()
    this.Prepare()
    this.TplName = "blog/article.html"
    return
}
