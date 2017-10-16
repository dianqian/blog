package base

import (
    "net/http"
    "nest/models/html/htmladmin"
    "github.com/astaxie/beego"
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
func (a *AdminCommonCtr) AdminBase() {
    // 执行基础内容
    a.CommonBase()

    // 进行在线判断
    if !a.IsLogin {
        // 将当前的ur获取，存入session中，然后在login之后，从session中取出，恢复指定页面
        if a.Ctx.Request.RequestURI != "/admin/login.html" {
            a.SetSession("next_uri", a.Ctx.Request.RequestURI)
        }
        a.Redirect("/admin/login.html", http.StatusFound)
        a.StopRun()
        return
    }

    htmlAdData := htmladmin.HTMLAdData{
        UserName: a.UserInfo.Name,
        Domain: beego.AppConfig.String(beego.BConfig.RunMode + "::domain"),
    }

    // 设置data和html信息
    a.Data["HTMLAdData"] = htmlAdData
    a.Layout = "admin/admin_layout.html"
}
