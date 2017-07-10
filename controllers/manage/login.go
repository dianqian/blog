package manage

import (
    "github.com/astaxie/beego"
    "net/http"
    "blog/controllers/base"
    "blog/models"
    "fmt"
    "crypto/md5"
    "encoding/hex"
    "time"
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
func (this *LoginController) Get ()  {
    this.Data["Title"] = "我的蘑菇鸡"

    // todo: 1. 判断是否已经登录；2. 如果已经登录，redirect到admin的首页；3. logout也会跳转到login的页面，所以也需要判断；

    // 判断是否已经完成登陆
    if this.IsLogin {
        this.Redirect("/admin", http.StatusFound)
        return
    }

    this.TplName = "admin/login.html"
    //this.TplName = "admin/example.html"
    return
}

/**
 login登录界面输入：user password后，点击submit后的post操作
 */
func (this *LoginController) Post()  {
    // 如果已经登陆，则跳转到管理页面
    if this.IsLogin {
        this.Redirect("/admin", http.StatusFound)
        return
    }

    // 执行登陆操作
    name := this.GetString("user")
    password := this.GetString("password")
    if name == "" || password == "" {
        logs.Debug(fmt.Sprintf("input username or password is null"))
        this.Redirect("/admin/login", http.StatusFound)
        return
    }

    // 校验user和密码;2. 如何显示提示信息
    user, err := Authenticate(name, password)
    if err != nil {
        logs.Info(fmt.Sprintf("authenticate failed: %s", err.Error()))
        // todo: 考虑如何将err msg带出去
        this.Redirect("/admin/login", http.StatusFound)
        return
    }

    // 设置登陆set login
    this.SetLogin(user)
    this.Redirect("/admin", http.StatusFound)
    return
}

/**
 密码相关的封装处理
 */
type Password string

func (this Password) Md5() string {
    h := md5.New()
    h.Write([]byte(this))
    return hex.EncodeToString(h.Sum(nil))
}

/**
 执行认证操作
 */
func Authenticate(name string, password string) (user *models.User, err error) {
    user = &models.User{Name: name}

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
        user.Updated = time.Now().Unix()
        user.Update("updated")
        return user, nil
    }
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

