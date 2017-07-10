package base

import (
    "github.com/astaxie/beego"
    "blog/models"
)

/**
 login的基础类
 */
type LoginBaseCtr struct {
    beego.Controller

    IsLogin     bool                        // 标记是否发生登陆
}

/**
 登陆的预加载操作，Prepare函数会在所有函数之前完成执行
 */
func (this * LoginBaseCtr) Prepare() {

    this.IsLogin = this.GetSession("user_info") != nil              // 没有该session，则认为没有登陆过
    if this.IsLogin {
        // todo: 获取该用户的相关信息
        //logs.Debug("the user is login")
    } else {
        //logs.Debug("the user is no login")
    }

    //this.IsLogin = true
    return
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


