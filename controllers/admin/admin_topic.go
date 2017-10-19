package admin

import (
    "nest/controllers/base"
    "github.com/astaxie/beego/logs"
    "fmt"
    "nest/models/db"
    "time"
    "net/http"
    "nest/models/html/htmladmin"
)


type TopicEditController struct {
    base.AdminCommonCtr
}

/**
 展示topic
 */
func (t *TopicEditController) Get ()  {
    // 预加载
    t.AdminBase()

    topicData := htmladmin.HTMLOneTopic{}

    // 尝试获取参数
    topicId, err := t.GetInt("id")
    if err == nil {
        // 是修改，需要保存在topic id值
        topic := new(db.Topic)
        topic.Id = topicId
        err := topic.Read("id")
        if err != nil {
            topicData.ErrorInfo = err.Error()
            t.Data["HTMLOneTopic"] = topicData
            t.TplName = "admin/admin_topic.html"
            return
        }

        // 编辑
        t.SetSession("modify_topic_id", topicId)
        topicData.Topic = *topic
        t.Data["HTMLOneTopic"] = topicData
        t.TplName = "admin/admin_topic.html"
        return
    } else {
        // 否则，新增
    }

    t.Data["HTMLOneTopic"] = topicData
    t.TplName = "admin/admin_topic.html"
    return
}

/**
 提交topic
 */
func (t *TopicEditController) Post()  {

    // 预处理的
    t.AdminBase()

    // 检查入参
    name := t.GetString("name")
    slugName := t.GetString("slug")
    desc := t.GetString("description")
    // todo: 需要做一个参数检查
    logs.Debug(fmt.Sprintf("%s, %s, %s", name, slugName, desc))

    // 准备数据，进行处理
    topic := new(db.Topic)
    topic.Name = name
    topic.Slug = slugName
    topic.Desc = desc
    topic.Updated = time.Now().Unix()

    var err error
    err = nil
    topicInter := t.GetSession("modify_topic_id")
    if topicInter == nil {
        // 没有modify topic id值，执行insert
        topic.Create = time.Now().Unix()
        topic.Status = 1
        err = topic.Insert()
        if err != nil {
            logs.Error(fmt.Sprintf("topic insert db failed: %s", err.Error()))
        }
    } else {
        // 是修改，执行update
        topic.Id = topicInter.(int)
        err = topic.Update("name", "slug", "desc", "updated")
        if err != nil {
            logs.Error(fmt.Sprintf("topic update db failed: %s", err.Error()))
        }
    }

    // 跳转操作
    if err != nil {
        // todo：编辑失败，需要进行提示和重新操作
    } else {
        // 编辑成功，跳转到list中
        t.Redirect("/admin/topics.html", http.StatusFound)
    }

    return
}