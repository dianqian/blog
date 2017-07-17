package models

import "github.com/astaxie/beego/orm"

/**
 系统管理相关相关
 */

type Blogger struct {
    Id              int

    Title           string          // 标题
    SubTitle        string          // 子标题
    Icon            string          // 图标

    BeiAn           string          // 备案号
    Copyright       string

    SeriesSay       string          // 专题前说
    ArchivesSay     string          // 归档前说

    Create      int64
    Updated     int64
    Status      int
}

/**
 读取到Blogger的配置信息
 */
func (b *Blogger) Read() error {

    b.Id = 1                    // 该配置只会有一条，仅支持修改

    err := orm.NewOrm().Read(b, "id")
    if err != nil {
        // todo: 具体错误处理，参考(https://beego.me/docs/mvc/model/object.md)
        // orm.ErrNoRows
        return err
    }
    return nil
}

/**
 插入该设置
 */
func (b *Blogger) Insert() error {
    b.Id = 1

    _, err := orm.NewOrm().Insert(b)
    if err != nil {
        return err
    }
    return nil
}

/**
 更新
 */
func (b *Blogger) Update(fields ...string) error {
    b.Id = 1

    _, err := orm.NewOrm().Update(b, fields...)
    if err != nil {
        return err
    }
    return nil
}


/**
 采用init的方式加载，而非在main.go文件中执行加载的模式
 */
func init()  {
    orm.RegisterModel(new(Blogger))
}


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


