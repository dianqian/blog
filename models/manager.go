package models

import "github.com/astaxie/beego/orm"

//-------------------------------NavInfo-----------------------------------
/**
 标题栏的内容
 */
type NavInfo struct {
    Id          int

    Name        string          // 名字
    ExtraUrl    string          //
    Description string

    Create      int64
    Updated     int64
    Status      int
}

func FindAllNavInfo() []*NavInfo {
    o := orm.NewOrm()

    navInfo := new(NavInfo)
    qs := o.QueryTable(navInfo)

    var list []*NavInfo
    qs.RelatedSel().All(&list)

    return list
}


