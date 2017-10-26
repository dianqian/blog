// @Description: article文章相关的数据模型的处理
package htmladmin

import (
    "nest/models/db"
    "nest/common"
    "time"
)

/**
 @Description：article相关的专题信息
*/
type TopicOfArticle struct {
    ID          int                         // 专题id
    Name        string                      // 专题name
    Checked     bool                       // 是否被当前文章选中
}

/**
 @Description：获取全部的
 @Param:
 @Return：
 */
func (t *TopicOfArticle) GetTopics() []*TopicOfArticle {
    var tpOfarticle []*TopicOfArticle
    tp := new(db.Topic)
    tps, err := tp.Select(common.OFFSET_ZEOR,common.LIMIT_ALL, common.STATUS_ALL)
    if err == nil {
        for _, item := range tps {
            tpOfarticle = append(tpOfarticle, &TopicOfArticle{ID: item.Id, Name: item.Name, Checked: false})
        }
    }

    return tpOfarticle
}

/**
 @Description：为article设置checked的topic
 */
func (t *TopicOfArticle) SetTopicChecked(aID int, tpOfarticle []*TopicOfArticle) ([]*TopicOfArticle) {

    ar := new(db.ArticleTopic)
    ats, err := ar.SelectByArticle(aID)
    if err == nil {
        for _, item := range ats {
            for _, tp := range tpOfarticle {
                if item.TopicId == tp.ID {
                    tp.Checked = true
                    break
                }
            }
        }
    }

    return tpOfarticle
}

/**
 @Description：article相关的tag
 */
type TagOfArticle struct {
    ID         string       `json:"id"`
    Name       string       `json:"name"`
}

/**
 @Description：获取所有的tag
 @Param:
 @Return：
 */
func (t *TagOfArticle) GetAllTags() ([]TagOfArticle, error) {
    tg := db.Tag{}
    tgs, err := tg.SelectAll()
    if err != nil {
        return nil, err
    }

    var tags []TagOfArticle
    for _, one := range tgs {
        tags = append(tags, TagOfArticle{ID: one.Name, Name: one.Name})
    }

    return tags, nil
}

/**
 @Description：获取article相关的tag
 @Param:
 @Return：
 */
func (t *TagOfArticle) GetTags(arID int) ([]string, error) {
    at := db.ArticleTag{ArticleId: arID}
    ats, err := at.SelectByArticle(arID)
    if err != nil {
        return nil, err
    }

    var tagIDs []int
    for _, one := range ats {
        tagIDs = append(tagIDs, one.TagId)
    }

    tg := db.Tag{}
    tgs, err := tg.GetByIDs(tagIDs)
    if err != nil {
        return nil, err
    }

    var toa []string
    for _, one := range tgs {
        toa = append(toa, one.Name)
    }

    return toa, nil
}

/**
 @Description：编辑信息
*/
type ArticleData struct {
    IsEdit          bool                // 是否是编辑，编辑：true；创建：false
    Title           string
    Status          int
    Slug            string
    Content         string
    PublishTime     time.Time
}

/**
 @Description：html article edit data information
               编辑的文章
*/
type HTMLArticleEditData struct {
    ErrorInfo   string                      // 错误的提示信息
    Article     ArticleData
    Topics      []*TopicOfArticle
    ArticleTags string
}

/**
 @Description：html article edit data information
               发布的文章数据
*/
type HTMLArticlesPublishData struct {
    ErrorInfo       string
    PageInfo        common.Page
    Articles        []*ArticlePublishData
}

/**
 @Description: article publish data
 */
type ArticlePublishData struct {
    ID              int
    Title           string
    Comment         int                     // 评论数
    Author          string
    Topic           string
    Create          time.Time
    Update          time.Time
    Publish         time.Time               // 发布时间
}

/**
 @Description：article draft html info
               回收文章数据
*/
type HTMLArticleDraftData struct {
    ErrorInfo           string
    PageInfo            common.Page
    ArticleInfo         []*ArticleDraftInfo
}

/**
 @Description: article draft info
 */
type ArticleDraftInfo struct {
    ID              int
    Title           string
    Author          string
    Topic           string
    Create          time.Time
    Update          time.Time
}


/**
 @Description：article trash html info
               回收站文章数据
*/
type HTMLArticleTrashData struct {
    ErrorInfo           string
    PageInfo            common.Page
    ArticleInfo         []*ArticleTrashInfo
}

/**
 @Description: article trash info
 */
type ArticleTrashInfo struct {
    ID              int
    Title           string
    Author          string
    Topic           string
    Create          time.Time
    Update          time.Time
}

/**
 @Description：预览文章的article
*/
type HTMLArticlePreview struct {
    ErrorInfo           string
    ArticleInfo         Article
}

// @Description：文章信息
type Article struct {
    ID              int                     // 文章ID
    Title           string                  // 标题
    Slug            string                  // 略缩信息，即用于url
    Header          string                  // 文章目录
    Content         string                  // 文章内容
    PublishTime     time.Time               // 发布时间
}