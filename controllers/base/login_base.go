package base

import (
    "github.com/astaxie/beego"
    "blog/models"
    "github.com/astaxie/beego/logs"
)

/**
 login的基础类
 */
type LoginBaseCtr struct {
    beego.Controller

    UserInfo    *models.User                // 用户信息
    IsLogin     bool                        // 标记是否发生登陆
}

/**
 登陆的预加载操作，Prepare函数会在所有函数之前完成执行
 */
func (this * LoginBaseCtr) Prepare() {

    this.IsLogin = this.GetSession("user_info") != nil              // 没有该session，则认为没有登陆过
    if this.IsLogin {
        logs.Debug("the user is login")
        // 获取该用户的相关信息
        user := this.GetLogin()
        if user != nil {
            this.UserInfo = user
        } else {
            // todo: 获取登陆用户信息失败
        }
    } else {
        logs.Debug("the user is no login")
    }

    return
}

/**
 获取登陆的用户
 */
func (this *LoginBaseCtr) GetLogin() *models.User {
    user := &models.User{Id: this.GetSession("user_info").(int)}
    err := user.Read("id")
    if err != nil {
        return nil
    }

    return user
}

/**
 记录登陆状态
 */
func (this *LoginBaseCtr) SetLogin(user *models.User) {
    this.IsLogin = true
    // 设置sesstion
    this.SetSession("user_info", user.Id)
    return
}

/**
 注销login信息
 */
func (this *LoginBaseCtr) DelLogin() {
    this.IsLogin = false
    // session中删除事件
    this.DelSession("user_info")
    return
}


