package admin

import (
    "nest/controllers/base"
    "time"
    "github.com/astaxie/beego/logs"
    "fmt"
    "net/http"
)

/**
 账号信息
 */
type Account struct {
    Name            string
    NickName        string

    Avatar          string
    Signature       string

    Email           string
    WebSite         string
    Wechat          string

    CreateTime          time.Time
    LastLoginTime       time.Time
    LastLogoutTime      time.Time

    LoginIp         string
}

/**
 用户信息管理
 */
type UserManageController struct {
    base.AdminCommonCtr
}

func (u *UserManageController) Get ()  {
    u.AdminBase()

    account := Account{}
    if u.UserInfo != nil {
        account.Name = u.UserInfo.Name
        account.NickName = u.UserInfo.NickName

        //account.Avatar = "/static/img/jiaoshou.jpg"
        account.Avatar = u.UserInfo.Avatar

        account.Email = u.UserInfo.Email
        account.WebSite = u.UserInfo.WebSite
        account.Wechat = u.UserInfo.Wechat

        account.Signature = u.UserInfo.Signature

        account.CreateTime = time.Unix(u.UserInfo.Create, 0)
        account.LastLoginTime = time.Unix(u.UserInfo.LastLoginTime, 0)
        account.LastLogoutTime = time.Unix(u.UserInfo.LastLogoutTime, 0)
    }

    u.Data["Account"] = account

    u.TplName = "admin/admin_user.html"
    return
}

/**
 修改头像avatar:
    1. 文件上传；
    2. 修改avatar的url链接；
 */
func (u *UserManageController) ChangeAvatarInfo()  {
    // todo: 待实现
    return
}

/**
 修改账户信息
 */
func (u *UserManageController) ChangeAccountInfo()  {

    u.UserInfo.Email = u.GetString("email")
    u.UserInfo.WebSite = u.GetString("website")
    u.UserInfo.Wechat = u.GetString("wechat")
    u.UserInfo.NickName = u.GetString("nickname")
    u.UserInfo.Signature = u.GetString("signature")

    u.UserInfo.Updated = time.Now().Unix()

    err := u.UserInfo.Update("email", "web_site", "wechat", "nick_name", "signature", "updated")
    if err != nil {
        logs.Error(fmt.Sprintf("update email/web_site/wechat/nick_name/signature failed: %s", err.Error()))
    }

    u.Redirect("/admin/user.html", http.StatusFound)

    return
}


/**
 修改登陆密码
 */
func (u *UserManageController) ChangePassword()  {

    oldpw := u.GetString("old")
    newpw := u.GetString("new")
    confirm := u.GetString("confirm")
    if newpw == "" || confirm == "" || oldpw == "" {
        logs.Error(fmt.Sprintf("input null, old/new/confirm."))
        u.Redirect("/admin/user.html", http.StatusFound)
        return
    }

    if newpw != confirm {
        logs.Error(fmt.Sprintf("new password not equal confirm password."))
        u.Redirect("/admin/user.html", http.StatusFound)
        return
    }

    if oldpw == newpw {
        logs.Error(fmt.Sprintf("new password equal old password."))
        u.Redirect("/admin/user.html", http.StatusFound)
        return
    }

    // 校验old password的合法性
    user, err := Authenticate(u.UserInfo.Name, oldpw)
    if err != nil {
        logs.Error(fmt.Sprintf("authenticate failed: %s.", err.Error()))
        u.Redirect("/admin/user.html", http.StatusFound)
        return
    }

    // 设置新password
    user.PassWord = Password(newpw).Md5()
    user.Updated = time.Now().Unix()
    err = user.Update("pass_word", "updated")
    if err != nil {
        logs.Error(fmt.Sprintf("save new password failed: %s.", err.Error()))
        u.Redirect("/admin/user.html", http.StatusFound)
        return
    }

    logs.Info(fmt.Sprintf("update password ok."))
    u.DelLogin()
    u.Redirect("/admin/user.html", http.StatusFound)
    return
}