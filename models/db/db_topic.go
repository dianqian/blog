package db

import "github.com/astaxie/beego/orm"

/**
 专题相关
 */

/**
 专题处理
 */
type Topic struct {
    Id              int
    Name            string          // 名字
    Slug            string          // 略缩名
    Desc            string          // 描述信息
    ArticleCount    int     // 该topic的文章统计数

    Create      int64
    Updated     int64
    Status      int
}

/**
 读取指定的topic
 */
func (t *Topic) Read(fields ...string) error {
    err := orm.NewOrm().Read(t, fields...)
    if err != nil {
        return err
    }

    return nil
}

/**
 更新指定topic
 */
func (t *Topic) Update(fields ...string) error {
    _, err := orm.NewOrm().Update(t, fields...)
    if err != nil {
        return err
    }

    return nil
}

/**
 insert插入一条topic
 */
func (t *Topic) Insert() error {
    _, err := orm.NewOrm().Insert(t)
    if err != nil {
        return err
    }
    return nil
}

/**
 select出所有的topic
 */
func (t *Topic) SelectAll() ([]*Topic, error) {
    var topics []*Topic

    qs := orm.NewOrm().QueryTable(t)
    _, err := qs.All(&topics)
    if err != nil {
        return nil, err
    }
    return topics, nil
}

/**
 采用init的方式加载，而非main.go文件中执行加载的模式

 */
func init()  {
    orm.RegisterModel(new(Topic))
}