package admin

import (
    "net/http"
    "blog/controllers/base"
    "blog/models/db"
    "fmt"
    "crypto/md5"
    "encoding/hex"
    "github.com/astaxie/beego/logs"
)

/**
 登录相关的控制器
 */
type LoginController struct {
    base.LoginBaseCtr
}

/**
 用于进入login登录界面的操作
 */
func (l *LoginController) Get ()  {
    l.Data["Title"] = "我的蘑菇鸡"

    // todo: 1. 判断是否已经登录；2. 如果已经登录，redirect到admin的首页；3. logout也会跳转到login的页面，所以也需要判断；

    // 判断是否已经完成登陆
    if l.IsLogin {
        nextUri := l.GetSession("next_uri")
        if nextUri == nil {
            l.Redirect("/admin", http.StatusFound)
        } else {
            l.Redirect(nextUri.(string), http.StatusFound)
        }

        return
    }

    l.TplName = "admin/admin_login.html"
    return
}

/**
 login登录界面输入：user password后，点击submit后的post操作
 */
func (l *LoginController) Post()  {
    // 如果已经登陆，则跳转到管理页面
    if l.IsLogin {
        nextUri := l.GetSession("next_uri")
        if nextUri == nil {
            l.Redirect("/admin", http.StatusFound)
        } else {
            l.Redirect(nextUri.(string), http.StatusFound)
        }
        return
    }

    // 执行登陆操作
    name := l.GetString("user")
    password := l.GetString("password")
    if name == "" || password == "" {
        logs.Debug(fmt.Sprintf("input username or password is null"))
        l.Redirect("/admin/login.html", http.StatusFound)
        return
    }

    // 校验user和密码;2. 如何显示提示信息
    user, err := Authenticate(name, password)
    if err != nil {
        logs.Info(fmt.Sprintf("authenticate failed: %s", err.Error()))
        // todo: 考虑如何将err msg带出去
        l.Redirect("/admin/login.html", http.StatusFound)
        return
    }

    logs.Debug(fmt.Sprintf("authenticate ok, login......"))
    // 设置登陆set login
    l.SetLogin(user)

    // 登陆完成，执行跳转
    nextUri := l.GetSession("next_uri")
    if nextUri == nil {
        l.Redirect("/admin", http.StatusFound)
    } else {
        l.Redirect(nextUri.(string), http.StatusFound)
    }
    return
}

/*******************************************
 密码相关的封装处理
 */
type Password string

func (p Password) Md5() string {
    h := md5.New()
    h.Write([]byte(p))
    return hex.EncodeToString(h.Sum(nil))
}

/**
 执行认证操作
 */
func Authenticate(name string, password string) (user *db.User, err error) {
    user = &db.User{Name: name}

    if err := user.Read("name"); err != nil {
        if err.Error() == "<QuerySeter> no row found" {
            return nil, fmt.Errorf("invalide name")
        }
        return nil, err
    } else if user.Id < 1 {
        return nil, fmt.Errorf("invalid name")
    } else if user.PassWord != Password(password).Md5() {
        return nil, fmt.Errorf("invalid password")
    } else {
        // update作为修改用户信息的时间time
        return user, nil
    }
}
/*******************************************/


/**
 登出操作
 */
type LogoutController struct {
    base.LoginBaseCtr
}

func (l *LogoutController) Get() {
    // todo: 具体业务，后续实现
    l.DelLogin()

    // 重定向到网站首页，匿名访问模式
    l.Redirect("/", http.StatusFound)
    return
}

