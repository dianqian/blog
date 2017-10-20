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

type DraftManageController struct {
    base.AdminCommonCtr
}

/**
 @Description：获取article的信息
 @Param:
 @Return：
 */
func (d *DraftManageController) dealDraftInfo(ars []*db.Article) ([]*htmladmin.ArticleDraftInfo, error) {
    var titles []*htmladmin.ArticleDraftInfo

    for _, item := range ars {
        title := htmladmin.ArticleDraftInfo{ID: item.Id, Title: item.Title}

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

/**
 @Description：db 初始化设置
 @Param:
 @Return：
 */
func (d *DraftManageController) Get ()  {
    // 基础调用函数
    d.AdminBase()

    // 解析参数
    pageNo, err := d.GetInt("page_no")
    if err != nil {
        pageNo = common.PAGE_NO_DEFAULT
    }
    pageSize, err := d.GetInt("page_size")
    if err != nil {
        pageSize = common.PAGE_SIZE_DEFAULT
    }
    offset, limit := common.GetOffsetLimit(pageNo, pageSize)

    arData := htmladmin.HTMLArticleDraftData{}
    // 获取数据
    article := new(db.Article)
    count, err := article.Count(common.ARTICLE_STATUS_DRAFT)
    if err != nil {
        logMsg := fmt.Sprintf("select count article for draft error: %s", err.Error())
        logs.Error(logMsg)

        arData.ErrorInfo = logMsg
        d.Data["HTMLArticleDraftData"] = arData
        d.TplName = "admin/admin_drafts.html"
        return
    }
    ars, err := article.Select(offset, limit, common.ARTICLE_STATUS_DRAFT)
    if err != nil {
        logMsg := fmt.Sprintf("select for`offset: %d, limit: %d` article error: %s", offset, limit, err.Error())
        logs.Error(logMsg)

        arData.ErrorInfo = logMsg
        d.Data["HTMLArticleDraftData"] = arData
        d.TplName = "admin/admin_drafts.html"
        return
    }

    drafts, err := d.dealDraftInfo(ars)
    if err != nil {
        logs.Error(err.Error())
        arData.ErrorInfo = err.Error()
        d.Data["HTMLArticleDraftData"] = arData
        d.TplName = "admin/admin_drafts.html"
        return
    }

    arData.ArticleInfo = drafts
    arData.PageInfo = common.PageUtil(int(count), pageNo, pageSize)

    d.Data["HTMLArticleDraftData"] = arData
    d.TplName = "admin/admin_drafts.html"
    return
}