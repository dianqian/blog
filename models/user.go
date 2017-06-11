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
 社交信息
 */
type Social struct {
    Id              int
    SortId          int                 // 排序值
    Name            string              // 名字
    ExtraUrl        string              // 外部链接
    Icon            string              // 图标

    Create          int64
    Updated         int64
    Status          int
}

func FindAllSocial() []*Social {
    o := orm.NewOrm()

    socialInfos := new(Social)
    qs := o.QueryTable(socialInfos)

    var list []*Social
    qs.RelatedSel().All(&list)

    return list
}