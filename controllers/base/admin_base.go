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

func (this *AdminCommonCtr) AdminBase() {
    // 执行基础内容
    this.CommonBase()

    // 进行在线判断
    if !this.IsLogin {
        // 将当前的ur获取，存入session中，然后在login之后，从session中取出，恢复指定页面
        this.SetSession("next_uri", this.Ctx.Request.RequestURI)
        this.Redirect("/admin/login", http.StatusFound)
        return
    }

    this.Data["UserName"] = this.UserInfo.Name
    this.Layout = "admin/admin_layout.html"
}
