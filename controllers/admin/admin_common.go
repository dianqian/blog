package admin

import (
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego/logs"
    "fmt"
    "time"
    "blog/models/db"
)

type ArticleTitleInfo struct {
    ID              int
    Title           string
    Author          string
    Topic           string
    Create          time.Time
    Update          time.Time
}

func (a *ArticleTitleInfo) GetTitleInfos(status int) ([]*ArticleTitleInfo, error) {
    var titles []*ArticleTitleInfo

    article := new(db.Article)
    articles, err := article.SelectByStatus(status)
    if err != nil {
        if err != orm.ErrNoRows {
            return nil, err
        }
    }
    for _, item := range articles {
        title := ArticleTitleInfo{ID: item.Id, Title: item.Title}

        // 准备author name
        u := new(db.User)
        u.Id = item.Author
        err := u.Read("id")
        if err != nil {
            logs.Error(fmt.Sprintf("get author'%s' failed: %s", item.Title, err.Error()))
        } else {
            if u.NickName != "" {
                title.Author = u.NickName
            } else {
                title.Author = u.Name
            }
        }

        // 准备Topic
        at := new(db.ArticleTopic)
        ats, err := at.SelectByArticle(item.Id)
        if err != nil {
            logs.Error(fmt.Sprintf("get topic'%s' failed: %s", item.Title, err.Error()))
        }
        if len(ats) == 1 {
            t := new(db.Topic)
            t.Id = ats[0].TopicId
            err := t.Read("id")
            if err != nil {
                logs.Error(fmt.Sprintf("get topic name'%s' failed: %s", item.Title, err.Error()))
            }
            title.Topic = t.Name
        }

        title.Create = time.Unix(item.Create, 0)
        title.Update = time.Unix(item.Updated, 0)

        titles = append(titles, &title)
    }

    return titles, nil
}