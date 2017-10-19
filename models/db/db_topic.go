package db

import (
    "github.com/astaxie/beego/orm"
    "nest/common"
)

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
 @Description：select by status，并且携带offset和limit
 @Param:
 @Return：
*/
func (t *Topic) Select(offset int, limit int, status int) ([]*Topic, error) {
    var topics []*Topic

    qs := orm.NewOrm().QueryTable(t).Offset(offset).Limit(limit)
    if status != common.STATUS_ALL {
        qs = qs.Filter("status", status)
    }

    _, err := qs.All(&topics)
    if err != nil {
        return nil, err
    }
    return topics, nil
}

/**
 @Description：计算status的count值
 @Param:
 @Return：
 */
func (t *Topic) Count(status int) (int64, error) {

    qs := orm.NewOrm().QueryTable(t)
    if status != common.STATUS_ALL {
        qs = qs.Filter("status", status)
    }

    count, err := qs.Count()
    if err != nil {
        return 0, err
    }

    return count, nil
}

/**
 采用init的方式加载，而非main.go文件中执行加载的模式

 */
func init()  {
    orm.RegisterModel(new(Topic))
}