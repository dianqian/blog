package admin

import (
    "nest/controllers/base"
    "time"
    "github.com/astaxie/beego/logs"
    "fmt"
    "nest/models/html/htmladmin"
    "nest/common"
)

/**
 用户信息管理
 */
type UserManageController struct {
    base.AdminCommonCtr
}

/**
 @Description：db 初始化设置
 @Param:
 @Return：
 */
func (u *UserManageController) Get ()  {
    u.AdminBase()

    account := htmladmin.HTMLUserData{}
    if u.UserInfo != nil {
        account.User.Name = u.UserInfo.Name
        account.User.NickName = u.UserInfo.NickName

        //account.Avatar = "/static/img/jiaoshou.jpg"
        account.User.Avatar = u.UserInfo.Avatar

        account.User.Email = u.UserInfo.Email
        account.User.WebSite = u.UserInfo.WebSite
        account.User.Wechat = u.UserInfo.Wechat

        account.User.Signature = u.UserInfo.Signature

        account.User.CreateTime = time.Unix(u.UserInfo.Create, 0)
        account.User.LastLoginTime = time.Unix(u.UserInfo.LastLoginTime, 0)
        account.User.LastLogoutTime = time.Unix(u.UserInfo.LastLogoutTime, 0)
    }

    u.Data["HTMLUserData"] = account

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
 @Description：修改账户信息
 @Param:
 @Return：
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
        logMsg := fmt.Sprintf("update email/web_site/wechat/nick_name/signature failed: %s", err.Error())
        logs.Error(logMsg)
        u.Data["json"] = common.CreateErrResponse(common.RESP_CODE_SYSTEM_ERR, logMsg, nil)
        u.ServeJSON()
        return
    }

    u.Data["json"] = common.CreateOkResponse(nil)
    u.ServeJSON()
    return
}

/**
 @Description：修改登陆密码
 @Param:
 @Return：
*/
func (u *UserManageController) ChangePassword()  {

    oldpw := u.GetString("old")
    newpw := u.GetString("new")
    confirm := u.GetString("confirm")
    if newpw == "" || confirm == "" || oldpw == "" {
        logMsg := fmt.Sprintf("input null, old/new/confirm.")
        logs.Error(logMsg)
        u.Data["json"] = common.CreateErrResponse(common.RESP_CODE_PARAMS_ERR, logMsg, nil)
        u.ServeJSON()
        return
    }

    if newpw != confirm {
        logMsg := fmt.Sprintf("new password not equal confirm password.")
        logs.Error(logMsg)
        u.Data["json"] = common.CreateErrResponse(common.RESP_CODE_PARAMS_ERR, logMsg, nil)
        u.ServeJSON()
        return
    }

    if oldpw == newpw {
        logMsg := fmt.Sprintf("new password equal old password.")
        logs.Error(logMsg)
        u.Data["json"] = common.CreateErrResponse(common.RESP_CODE_PARAMS_ERR, logMsg, nil)
        u.ServeJSON()
        return
    }

    // 校验old password的合法性
    user, err := Authenticate(u.UserInfo.Name, oldpw)
    if err != nil {
        logMsg := fmt.Sprintf("authenticate failed: %s.", err.Error())
        logs.Error(logMsg)
        u.Data["json"] = common.CreateErrResponse(common.RESP_CODE_PARAMS_ERR, logMsg, nil)
        u.ServeJSON()
        return
    }

    // 设置新password
    user.PassWord = Password(newpw).Md5()
    user.Updated = time.Now().Unix()
    err = user.Update("pass_word", "updated")
    if err != nil {
        logMsg := fmt.Sprintf("save new password failed: %s.", err.Error())
        logs.Error(logMsg)
        u.Data["json"] = common.CreateErrResponse(common.RESP_CODE_PARAMS_ERR, logMsg, nil)
        u.ServeJSON()
        return
    }

    u.Data["json"] = common.CreateOkResponse(nil)
    u.ServeJSON()
    return
}