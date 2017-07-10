package models

import "github.com/astaxie/beego/orm"

/**
 注册用户信息
 */
type User struct {
    Id                  int

    Name                string                  // 用户名
    PassWord            string                  // 密码
    Avatar              string                  // 头像
    Signature           string                  // 签名
    Email               string                  // 邮箱
    WebSite             string                  // 站点

    Create      int64
    Updated     int64
    Status      int
}

/**
 基于指定字段读取数据
 */
func (this *User) Read(fields ...string) error {
    err := orm.NewOrm().Read(this, fields...)
    if err != nil {
        return err
    }
    return nil
}

/**
 更新指定字段的处理
 */
func (this *User) Update(fields ...string) error {
    _, err := orm.NewOrm().Update(this, fields...)
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