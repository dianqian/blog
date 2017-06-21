package manage

import (
    "github.com/astaxie/beego"
    "net/http"
)

/**
 登录相关的控制器
 */
type LoginController struct {
    beego.Controller
}

/**
 用于进入login登录界面的操作
 */
func (this *LoginController) Get ()  {
    this.Data["Title"] = "我的蘑菇鸡"

    // todo: 1. 判断是否已经登录；2. 如果已经登录，redirect到admin的首页；3. logout也会跳转到login的页面，所以也需要判断；

    this.TplName = "admin/login.html"
    //this.TplName = "admin/example.html"
    return
}

/**
 login登录界面输入：user password后，点击submit后的post操作
 */
func (this *LoginController) Post()  {
    name := this.GetString("name")
    password := this.GetString("password")
    if name == "" || password == "" {
        this.Redirect("/admin/login", http.StatusFound)
        return
    }

    // todo: 1. 校验user和密码;2. 如何显示提示信息
    if false {
        // 校验失败
        this.Redirect("/admin/login", http.StatusFound)
        return
    }

    // todo: 校验通过的相关信息
    this.Redirect("/admin", http.StatusFound)
    return
}


/**
 登出操作
 */
type LogoutController struct {
    beego.Controller
}

func (this *LogoutController) Get() {
    // todo: 具体业务，后续实现

    // todo：重定向到网站首页，匿名访问模式
    this.Redirect("/", http.StatusFound)
    return
}

