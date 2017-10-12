// @Description: article文章相关的数据模型的处理
package htmladmin

import (
    "blog/models/db"
    "time"
)

/**
 @Description：article相关的专题信息
*/
type TopicOfArticle struct {
    ID          int                         // 专题id
    Name        string                      // 专题name
    Checked     string                      // 是否被当前文章选中
}

/**
 @Description：获取全部的
 @Param:
 @Return：
 */
func (t *TopicOfArticle) GetTopics() []*TopicOfArticle {
    var tpOfarticle []*TopicOfArticle
    tp := new(db.Topic)
    tps, err := tp.SelectAll()
    if err == nil {
        for _, item := range tps {
            tpOfarticle = append(tpOfarticle, &TopicOfArticle{ID: item.Id, Name: item.Name})
        }
    }

    return tpOfarticle
}

/**
 @Description：为article设置checked的topic
 */
func (t *TopicOfArticle) SetTopicChecked(aID int, tpOfarticle []*TopicOfArticle) ([]*TopicOfArticle, string) {
    isDefault := "checked"

    ar := new(db.ArticleTopic)
    ats, err := ar.SelectByArticle(aID)
    if err != nil {
        for _, item := range ats {
            for _, tp := range tpOfarticle {
                if item.TopicId == tp.ID {
                    tp.Checked = "checked"
                    isDefault = ""
                    break
                }
            }
        }
    }

    return tpOfarticle, isDefault
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
type Edit struct {
    IsEdit          bool                // 是否是编辑，编辑：true；创建：false
    Title           string
    IsDraft         bool
    Slug            string
    Content         string
    PublishTime     time.Time
}

/**
 @Description：html article edit data information
*/
type HTMLArticleEditData struct {
    Edit        Edit
    DefaultTp   string
    Series      []*TopicOfArticle
    ArticleTags string
    AllTags     []TagOfArticle
}