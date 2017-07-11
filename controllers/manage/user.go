package manage

import (
    "blog/controllers/base"
    "blog/models"
    "time"
)

/**
 用户信息管理
 */
type UserManageController struct {
    base.AdminCommonCtr
}

func (this *UserManageController) Get ()  {
    this.AdminBase()

    account := &models.Account{}
    account.UserName = this.UserInfo.Name
    account.CreateTime = time.Unix(this.UserInfo.Create, 0)
    account.BlogName = "weisen's blog"
    account.Avatar = "/static/img/jiaoshou.jpg"

    this.Data["Account"] = account

    this.TplName = "admin/user.html"
    return
}

/**
 修改账户信息
 */
func (this *UserManageController) ChangeAccountInfo()  {
    return
}

/**
 修改blog信息
 */
func (this *UserManageController) ChangeBlogInfo()  {
    return 
}

/**
 修改登陆密码
 */
func (this *UserManageController) ChangePassword()  {
    return
}