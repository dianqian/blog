package models

import "github.com/astaxie/beego/orm"

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

/**
 初始化
 */
func init() {
    orm.RegisterModel(new(Tag))
}
