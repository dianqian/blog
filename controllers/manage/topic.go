package manage

import (
    "blog/controllers/base"
    "github.com/astaxie/beego/logs"
    "fmt"
    "blog/models"
    "time"
    "net/http"
)


type TopicEditController struct {
    base.AdminCommonCtr
}

/**
 展示topic
 */
func (this *TopicEditController) Get ()  {
    // 预加载
    this.AdminBase()

    // 尝试获取参数
    topicId, err := this.GetInt("topic_id")
    if err == nil {
        // 是修改，需要保存在topic id值
        topic := new(models.Topic)
        topic.Id = topicId
        err := topic.Read("id")
        if err == nil {
            this.SetSession("modify_topic_id", topicId)
            this.Data["EditTopic"] = topic
        }
    } else {
        // 否则，新增
    }

    this.TplName = "admin/topic.html"
    return
}

/**
 提交topic
 */
func (this *TopicEditController) Post()  {

    // 预处理的
    this.AdminBase()

    // 检查入参
    name := this.GetString("name")
    slugName := this.GetString("slug")
    desc := this.GetString("description")
    // todo: 需要做一个参数检查
    logs.Debug(fmt.Sprintf("%s, %s, %s", name, slugName, desc))

    // 准备数据，进行处理
    topic := new(models.Topic)
    topic.Name = name
    topic.Slug = slugName
    topic.Desc = desc
    topic.Updated = time.Now().Unix()

    var err error
    err = nil
    topicInter := this.GetSession("modify_topic_id")
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
        this.Redirect("/admin/manage-topics", http.StatusFound)
    }

    return
}