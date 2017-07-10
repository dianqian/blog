package models

import "github.com/astaxie/beego/orm"

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

/**
 采用init的方式加载，而非在main.go文件中执行加载的模式
 */
//func init()  {
//    orm.RegisterModel(new(Social))
//}
