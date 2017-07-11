package manage

import (
    "blog/controllers/base"
    "github.com/astaxie/beego/logs"
)


type AdminController struct {
    base.AdminCommonCtr
}

func (this *AdminController) Get ()  {
    // 显示执行基类
    this.AdminBase()

    //this.TplName = "admin/admin_layout.html"
    this.TplName = "admin/article.html"
    logs.Debug("AdminController Get ....")
    return
}