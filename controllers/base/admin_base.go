package base

import (
    "net/http"
)

/**
 admin管理的Controller的公共基础类
 */
type AdminCommonCtr struct {
    CommonCtr
}

/**
 todo： 这里参考beego中(https://beego.me/docs/mvc/controller/controller.md)的实现，后期可以做进一步的优化处理
 */
func (this *AdminCommonCtr) AdminBase() {
    // 执行基础内容
    this.CommonBase()

    // 进行在线判断
    if !this.IsLogin {
        // 将当前的ur获取，存入session中，然后在login之后，从session中取出，恢复指定页面
        if this.Ctx.Request.RequestURI != "/admin/login" {
            this.SetSession("next_uri", this.Ctx.Request.RequestURI)
        }
        this.Redirect("/admin/login", http.StatusFound)
        return
    }

    this.Data["UserName"] = this.UserInfo.Name
    this.Layout = "admin/admin_layout.html"
}
