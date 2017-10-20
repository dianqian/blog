package db

import (
    "github.com/astaxie/beego/orm"
    "fmt"
    "nest/common"
)

/**
 文章相关
 */

/**
 文章的基本信息
 */
type Article struct {
    Id          int

    Title       string                  // 标题
    Url         string                  // url
    Author      int                     // 作者，对应的user id
    PublishTime int64                   // 发布时间

    Content     string                  // 文章内容，采用markdown方式，longText的类型

    Create      int64                   // 创建时间
    Updated     int64                   // 更新时间
    Status      int                     // 状态：0：回收(删除)；2：草稿；1：正式文章
}

/**
 根据指定字段读取数据
 */
func (a *Article) Read(fields...string) error {
    err := orm.NewOrm().Read(a, fields...)
    if err != nil {
        // todo: orm.ErrNoRows
        return err
    }
    return nil
}

/**
 执行update操作
 */
func (a *Article) Update(fields...string) error {
    _, err := orm.NewOrm().Update(a, fields...)
    if err != nil {
        return err
    }
    return nil
}

/**
 插入一条数据
 */
func (a *Article) Insert() (int64, error) {
    id, err := orm.NewOrm().Insert(a)
    if err != nil {
        return 0, err
    }
    return id, nil
}

/**
 计算count
 */
func (a *Article) Count(status int) (int64, error) {
    qs := orm.NewOrm().QueryTable(a)
    if status != common.ARTICLE_STATUS_ALL {
        qs = qs.Filter("status", status)
    }

    // 得到count
    cnt, err := qs.Count()
    if err != nil {
        return 0, fmt.Errorf("table Article: %s", err.Error())
    }

    return cnt, nil
}

/**
 携带offset、limit的配置信息
 */
func (a *Article) Select(offset int, limit int, status int) ([]*Article, error) {
    var articles []*Article

    qs := orm.NewOrm().QueryTable(a).Offset(offset).Limit(limit)
    if status != common.ARTICLE_STATUS_ALL {
        qs = qs.Filter("status", status)
    }

    _, err := qs.All(&articles)
    if err != nil {
        return articles, err
    }

    return articles, nil
}

/**
 获取article的pre和next
 上一篇、下一篇，采用publish time时间，发布的时间
 direction: 1----next;-1----pre
 */
func (a *Article) SelectBrother(direction int, publishTime int64) (error) {

    qs := orm.NewOrm().QueryTable(a)

    if direction == common.BROTHER_PREV {
        qs = qs.Filter("publish_time__lt", publishTime)
    } else if direction == common.BROTHER_NEXT {
        qs = qs.Filter("publish_time__gt", publishTime)
    }

    var articles []*Article
    qs = qs.Limit(1)
    _, err := qs.All(&articles)
    if err != nil {
        return err
    }
    if len(articles) != 1 {
        return fmt.Errorf("no brother article")
    }

    a.Id = articles[0].Id
    a.Title = articles[0].Title
    a.Url = articles[0].Url
    a.PublishTime = articles[0].PublishTime
    return nil
}


/**
 根据status字选择
 */
func (a *Article) SelectByStatus(status int) ([]*Article, error) {
    var articles []*Article

    qs := orm.NewOrm().QueryTable(a).Filter("status", status)
    _, err := qs.All(&articles)
    if err != nil {
        return nil, err
    }

    return articles, nil
}


/**
 article和topic对应的表
 */
type ArticleTopic struct {
    Id          int

    ArticleId   int                     // 文章article id
    TopicId     int                     // topic id值

    Create      int64                   // 创建时间
    Updated     int64                   // 更新时间
    Status      int                     // 状态：1--有效；0--无效
}

/**
 根据指定字段读取数据
 */
func (a *ArticleTopic) Read(fields...string) error {
    err := orm.NewOrm().Read(a, fields...)
    if err != nil {
        return err
    }
    return nil
}

/**
 更新行数据
 */
func (a *ArticleTopic) Update(fields...string) error {
    _, err := orm.NewOrm().Update(a, fields...)
    if err != nil {
        return err
    }
    return nil
}

/**
 插入一条数据
 */
func (a *ArticleTopic) Insert() (int64, error) {
    id, err := orm.NewOrm().Insert(a)
    if err != nil {
        return 0, err
    }
    return id, nil
}

/**
 select by article id.
 */
func (a *ArticleTopic) SelectByArticle(aID int) ([]*ArticleTopic, error) {
    var at []*ArticleTopic

    qs := orm.NewOrm().QueryTable(a)
    _, err := qs.Filter("article_id", aID).All(&at)
    if err != nil {
        return nil, err
    }
    return at, nil
}

/**
 基于article id查相关的文章数
 */
func (a *ArticleTopic) CountForArticle(tpID int) (int, error)  {
    qs := orm.NewOrm().QueryTable(a)
    qs= qs.Filter("topic_id", tpID)

    // 得到count
    cnt, err := qs.Count()
    if err != nil {
        return 0, fmt.Errorf("table ArticleTopic: %s", err.Error())
    }

    return int(cnt), nil
}

/**
 select by topic id.
 */
func (a *ArticleTopic) SelectByTp(tID int) ([]*ArticleTopic, error) {
    var at []*ArticleTopic

    qs := orm.NewOrm().QueryTable(a)
    _, err := qs.Filter("topic_id", tID).All(&at)
    if err != nil {
        return nil, err
    }
    return at, nil
}

/**
 article和tag对应的表
 */
type ArticleTag struct {
    Id          int

    ArticleId   int                     // article id
    TagId       int                     // topic id值
    Create      int64                   // 创建时间
    Updated     int64                   // 更新时间
    Status      int                     // 状态：1--有效；0--无效
}

/**
 select by article id.
 */
func (a *ArticleTag) SelectByArticle(aID int) ([]*ArticleTag, error) {
    var at []*ArticleTag

    qs := orm.NewOrm().QueryTable(a)
    _, err := qs.Filter("article_id", aID).All(&at)
    if err != nil {
        return nil, err
    }
    return at, nil
}

/**
 select by topic id.
 */
func (a *ArticleTag) SelectByTp(tID int) ([]*ArticleTag, error) {
    var at []*ArticleTag

    qs := orm.NewOrm().QueryTable(a)
    _, err := qs.Filter("tag_id", tID).All(&at)
    if err != nil {
        return nil, err
    }
    return at, nil
}

/**
 附件文章的，todo：待处理
 */
type AttachFile struct {
    Id          int

    Name        string                  // 文件名
    Url         string                  // 文件存储的url
}


/**
 采用init的方式加载，而非在main.go文件中执行加载的模式
 */
func init()  {
    orm.RegisterModel(new(Article))
    orm.RegisterModel(new(ArticleTopic))
    orm.RegisterModel(new(ArticleTag))
}