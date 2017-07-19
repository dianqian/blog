package manage

import (
    "blog/controllers/base"
    "blog/models"
    "time"
    "github.com/astaxie/beego/logs"
    "fmt"
)

type DraftInfo struct {
    ID              int
    Title           string
    Author          string
    Topic           string
    Create          time.Time
    Update          time.Time
}

type DraftManageController struct {
    base.AdminCommonCtr
}

func (d *DraftManageController) Get ()  {
    // 基础调用函数
    d.AdminBase()

    var drafts []*DraftInfo

    article := new(models.Article)
    articles, err := article.SelectByStatus(base.ARTICLE_STATUS_DRAFT)
    if err != nil {
        logs.Error(fmt.Sprintf("list draft article failed: %s", err.Error()))
    }
    for _, item := range articles {
        draft := DraftInfo{ID: item.Id, Title: item.Title}

        // 准备author name
        u := new(models.User)
        u.Id = item.Author
        err := u.Read("id")
        if err != nil {
            logs.Error(fmt.Sprintf("get author'%s' failed: %s", item.Title, err.Error()))
        } else {
            if u.NickName != "" {
                draft.Author = u.NickName
            } else {
                draft.Author = u.Name
            }
        }

        // 准备Topic
        at := new(models.ArticleTopic)
        ats, err := at.SelectByArticle(item.Id)
        if err != nil {
            logs.Error(fmt.Sprintf("get topic'%s' failed: %s", item.Title, err.Error()))
        }
        if len(ats) == 1 {
            t := new(models.Topic)
            t.Id = ats[0].TopicId
            err := t.Read("id")
            if err != nil {
                logs.Error(fmt.Sprintf("get topic name'%s' failed: %s", item.Title, err.Error()))
            }
            draft.Topic = t.Name
        }

        draft.Create = time.Unix(item.Create, 0)
        draft.Update = time.Unix(item.Updated, 0)

        drafts = append(drafts, &draft)
    }

    d.Data["List"] = drafts

    d.TplName = "admin/drafts.html"
    return
}