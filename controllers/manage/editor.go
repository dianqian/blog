package manage

import (
    "blog/controllers"
)


type EditorController struct {
    controllers.AdminCommonCtr
}

func (this *EditorController) Get ()  {
    this.PreBase()

    //this.Prepare()
    this.TplName = "admin/editor.html"
    return
}