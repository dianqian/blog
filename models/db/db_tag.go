package db

import (
    "github.com/astaxie/beego/orm"
    "fmt"
)

/**
 标签相关
 */

/**
 标签
 */
type Tag struct {
    Id      int
    Name    string

    Create      int64
    Updated     int64
    Status      int
}

func (tg *Tag) TableName() string {
    return "tag"
    //return reflect.TypeOf(*tg).Name()
}

/**
 tag read
 */
func (tg *Tag) Read(fields ...string) error {
    err := orm.NewOrm().Read(tg, fields...)
    if err != nil {
        return err
    }
    return nil
}

/**
 tag update
 */
func (tg *Tag) Update(fields ...string) error {
    _, err := orm.NewOrm().Update(tg, fields...)
    if err != nil {
        return err
    }
    return nil
}

/**
 tag insert
 */
func (tg *Tag) Insert() error {
    _, err := orm.NewOrm().Insert(tg)
    if err != nil {
        return err
    }
    return nil
}

/**
 all tags
 todo: 缺少对于status的过滤
 */
func (tg *Tag) SelectAll() ([]*Tag, error) {
    var tags []*Tag

    qs := orm.NewOrm().QueryTable(tg)
    _, err := qs.All(&tags)
    if err != nil {
        return nil, err
    }

    return tags, nil
}

// 用于编码方便加入的
const (
    SELECT_COND_FALSE = "1=2"
    SELECT_COND_TRUE = "1=1"
)

/**
 @Description：根据ids获取tag的对象列表
 @Param:
 @Return：
 */
func (tg *Tag) GetByIDs(ids []int) ([]*Tag, error) {
    var tags []*Tag

    if len(ids) <= 0 {
        return nil, fmt.Errorf("no input id")
    }

    qs := orm.NewOrm().QueryTable(tg).Filter("id__in", ids)

    _, err := qs.All(&tags)
    if err != nil {
        return nil, err
    }

    return tags, nil
}

/**
 初始化
 */
func init() {
    orm.RegisterModel(new(Tag))
}
