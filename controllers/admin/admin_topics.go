package admin

import (
    "blog/controllers/base"
    "blog/models/db"
    "time"
    "github.com/astaxie/beego/logs"
    "fmt"
)

/**
 专题的简要信息
 */
type TopicBriefInfo struct {
    Id                  int
    Name                string
    CreateTime          time.Time
    ArticleCount        int
}

type TopicManageController struct {
    base.AdminCommonCtr
}

func (t *TopicManageController) Get ()  {
    // 基础组件的执行
    t.AdminBase()

    // 准备数据
    var topicsData []*TopicBriefInfo
    topic := new(db.Topic)
    topics, err := topic.SelectAll()
    if err != nil {
        // todo: select出错处理
        logs.Error(fmt.Sprintf("topic select all failed: %s", err.Error()))
    } else {
        for _, item := range topics {
            topicsData = append(topicsData, &TopicBriefInfo{Id: item.Id, Name: item.Name,
                CreateTime: time.Unix(item.Create, 0), ArticleCount: item.ArticleCount})
        }
    }

    // 将数据带出
    t.Data["TopicList"] = topicsData
    t.TplName = "admin/admin_topics.html"
    return
}