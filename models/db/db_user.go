package db

import "github.com/astaxie/beego/orm"

/**
 用户相关
 */

/**
 注册用户信息
 */
type User struct {
    Id                  int

    Name                string                  // 用户名
    PassWord            string                  // 密码

    NickName            string                  // 昵称
    Avatar              string                  // 头像
    Signature           string                  // 签名

    Email               string                  // 邮箱
    WebSite             string                  // 站点
    Wechat              string                  // 微信

    Create      int64
    Updated     int64
    Status      int

    LoginTime           int64                   // 本次登录的时间
    LastLoginTime       int64                   // 最后一次登录的时间，在退出本次登录时更新
    LastLogoutTime      int64                   // 最后一次登出的时间，在本次退出时更新
}

/**
 基于指定字段读取数据
 */
func (u *User) Read(fields ...string) error {
    err := orm.NewOrm().Read(u, fields...)
    if err != nil {
        // todo: 具体错误处理，参考(https://beego.me/docs/mvc/model/object.md)
        // orm.ErrNoRows
        return err
    }
    return nil
}

/**
 更新指定字段的处理
 */
func (u *User) Update(fields ...string) error {
    _, err := orm.NewOrm().Update(u, fields...)
    if err != nil {
        return err
    }
    return nil
}

/**
 临时用户
 */
type TmpUser struct {
    Id                  int

    Name                string                  // 名字
    Email               string                  // 邮箱
    WebSite             string                  // 站点

    Create              int64
    Updated             int64
    Status              int
}

/**
 采用init的方式加载，而非在main.go文件中执行加载的模式
 */
func init()  {
    orm.RegisterModel(new(User))
    orm.RegisterModel(new(TmpUser))
}