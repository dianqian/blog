package base

import (
    "github.com/astaxie/beego"
    "nest/models/db"
    "github.com/astaxie/beego/logs"
    "time"
    "fmt"
)

/**
 login的基础类
 */
type LoginBaseCtr struct {
    beego.Controller

    UserInfo    *db.User                // 用户信息
    IsLogin     bool                        // 标记是否发生登陆
}

/**
 登陆的预加载操作，Prepare函数会在所有函数之前完成执行
 */
func (l * LoginBaseCtr) Prepare() {

    l.IsLogin = l.GetSession("user_info") != nil              // 没有该session，则认为没有登陆过
    if l.IsLogin {
        //logs.Debug("the user is login")
        // 获取该用户的相关信息
        user := l.GetLogin()
        if user != nil {
            l.UserInfo = user
        } else {
            // todo: 获取登陆用户信息失败
        }
    } else {
        //logs.Debug("the user is no login")
    }

    return
}

/**
 获取登陆的用户
 */
func (l *LoginBaseCtr) GetLogin() *db.User {
    user := &db.User{Id: l.GetSession("user_info").(int)}
    err := user.Read("id")
    if err != nil {
        return nil
    }

    return user
}

/**
 记录登陆状态
 */
func (l *LoginBaseCtr) SetLogin(user *db.User) {
    l.IsLogin = true
    // 设置sesstion
    l.SetSession("user_info", user.Id)

    // 记录本次登录的时间
    l.UserInfo = user
    l.UserInfo.LastLoginTime = user.LoginTime
    l.UserInfo.LoginTime = time.Now().Unix()
    err := l.UserInfo.Update("login_time", "last_login_time")
    if err != nil {
        logs.Error(fmt.Sprintf("update login/last login time failed: %s", err.Error()))
    }

    return
}

/**
 注销login信息
 */
func (l *LoginBaseCtr) DelLogin() {
    l.IsLogin = false
    // session中删除事件
    l.DelSession("user_info")

    // 记录本次退出的时间
    if l.UserInfo != nil {
        l.UserInfo.LastLogoutTime = time.Now().Unix()
        err := l.UserInfo.Update("last_logout_time")
        if err != nil {
            logs.Error(fmt.Sprintf("update logout time failed: %s", err.Error()))
        }
    }

    return
}


