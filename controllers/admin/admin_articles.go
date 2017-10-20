package admin

import (
    "nest/controllers/base"
    "nest/common"
    "github.com/astaxie/beego/logs"
    "fmt"
    "time"
    "nest/models/html/htmladmin"
    "nest/models/db"
)

type PublishManageController struct {
    base.AdminCommonCtr
}

/**
 @Description：获取article的信息
 @Param:
 @Return：
 */
func (p *PublishManageController) dealPublishInfo(ars []*db.Article) ([]*htmladmin.ArticlePublishData, error) {
    var titles []*htmladmin.ArticlePublishData

    for _, item := range ars {
        // todo: comment count
        title := htmladmin.ArticlePublishData{ID: item.Id, Title: item.Title}

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
        title.Publish = time.Unix(item.PublishTime, 0)

        titles = append(titles, &title)
    }

    return titles, nil
}

/**
 @Description：db 初始化设置
 @Param:
 @Return：
 */
func (p *PublishManageController) Get ()  {
    // 基础调用函数
    p.AdminBase()

    // 解析参数
    pageNo, err := p.GetInt("page_no")
    if err != nil {
        pageNo = common.PAGE_NO_DEFAULT
    }
    pageSize, err := p.GetInt("page_size")
    if err != nil {
        pageSize = common.PAGE_SIZE_DEFAULT
    }
    offset, limit := common.GetOffsetLimit(pageNo, pageSize)

    arData := htmladmin.HTMLArticlesPublishData{}
    // 获取数据
    article := new(db.Article)
    count, err := article.Count(common.ARTICLE_STATUS_PUBLISH)
    if err != nil {
        logMsg := fmt.Sprintf("select count article for pubblish error: %s", err.Error())
        logs.Error(logMsg)

        arData.ErrorInfo = logMsg
        p.Data["HTMLArticlesPublishData"] = arData
        p.TplName = "admin/admin_articles.html"
        return
    }
    ars, err := article.Select(offset, limit, common.ARTICLE_STATUS_PUBLISH)
    if err != nil {
        logMsg := fmt.Sprintf("select for`offset: %d, limit: %d` article error: %s", offset, limit, err.Error())
        logs.Error(logMsg)

        arData.ErrorInfo = logMsg
        p.Data["HTMLArticlesPublishData"] = arData
        p.TplName = "admin/admin_articles.html"
        return
    }

    drafts, err := p.dealPublishInfo(ars)
    if err != nil {
        logs.Error(err.Error())
        arData.ErrorInfo = err.Error()
        p.Data["HTMLArticlesPublishData"] = arData
        p.TplName = "admin/admin_articles.html"
        return
    }

    arData.Articles = drafts
    arData.PageInfo = common.PageUtil(int(count), pageNo, pageSize)

    p.Data["HTMLArticlesPublishData"] = arData
    p.TplName = "admin/admin_articles.html"
    return
}