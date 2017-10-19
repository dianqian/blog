package admin

import (
    "nest/controllers/base"
    "nest/models/db"
    "github.com/astaxie/beego/logs"
    "fmt"
    "nest/common"
    "nest/models/html/htmladmin"
    "time"
)

type TopicManageController struct {
    base.AdminCommonCtr
}

func (t *TopicManageController) Get ()  {
    // 基础组件的执行
    t.AdminBase()

    // 解析参数
    pageNo, err := t.GetInt("page_no")
    if err != nil {
        pageNo = common.PAGE_NO_DEFAULT
    }
    pageSize, err := t.GetInt("page_size")
    if err != nil {
        pageSize = common.PAGE_SIZE_DEFAULT
    }
    offset, limit := common.GetOffsetLimit(pageNo, pageSize)
    status, err := t.GetInt("status")
    if err != nil {
        status = common.STATUS_ALL
    }

    tpData := htmladmin.HTMLTopicsData{}
    // 获取数据
    topic := new(db.Topic)
    count, err := topic.Count(status)
    if err != nil {
        logMsg := fmt.Sprintf("select count topics error: %s", err.Error())
        logs.Error(logMsg)

        tpData.ErrorInfo = logMsg
        t.Data["HTMLTopicsData"] = tpData
        t.TplName = "admin/admin_topics.html"
        return
    }

    tps, err := topic.Select(offset, limit, status)
    if err != nil {
        logMsg := fmt.Sprintf("select for`offset: %d, limit: %d` topics error: %s", offset, limit, err.Error())
        logs.Error(logMsg)

        tpData.ErrorInfo = logMsg
        t.Data["HTMLTopicsData"] = tpData
        t.TplName = "admin/admin_topics.html"
        return
    }

    for _, item := range tps {
        one := &htmladmin.TopicBriefInfo{
            Id: item.Id,
            Name: item.Name,
            Slug: item.Slug,
            ArticleCount: item.ArticleCount,
            Created: time.Unix(item.Create, 0),
            Updated: time.Unix(item.Updated, 0),
            Status: item.Status,
        }
        tpData.TopicInfo = append(tpData.TopicInfo, one)
    }
    tpData.PageInfo = common.PageUtil(int(count), pageNo, pageSize)

    // 将数据带出
    t.Data["HTMLTopicsData"] = tpData
    t.TplName = "admin/admin_topics.html"
    return
}