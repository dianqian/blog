package admin

import (
    "nest/controllers/base"
    "github.com/astaxie/beego/logs"
    "fmt"
    "nest/models/db"
    "time"
    "nest/models/html/htmladmin"
    "nest/common"
)


type TopicEditController struct {
    base.AdminCommonCtr
}

/**
 @Description：获取一个topic的详情
 @Param:
 @Return：
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
 @Description：提交一个Topic的信息
 @Param:
 @Return：
*/
func (t *TopicEditController) Post()  {

    // 预处理的
    t.AdminBase()

    // 检查入参
    name := t.GetString("name")
    if name == "" {
        logMsg := "input name empty"
        logs.Error(logMsg)
        t.Data["json"] = common.CreateErrResponse(common.RESP_CODE_PARAMS_ERR, logMsg, nil)
        t.ServeJSON()
        return
    }
    slugName := t.GetString("slug")
    desc := t.GetString("description")

    // id检查
    idData := t.GetSession("modify_topic_id")
    if idData == nil {
        // 新建
        tp := &db.Topic{
            Name: name,
            Slug: slugName,
            Desc: desc,
            Create: time.Now().Unix(),
            Status: common.STATUS_VALID,
        }
        tp.Updated = time.Now().Unix()
        err := tp.Insert()
        if err != nil {
            logMsg := fmt.Sprintf("insert for `%s` error: %s", name, err.Error())
            logs.Error(logMsg)
            t.Data["json"] = common.CreateErrResponse(common.RESP_CODE_SYSTEM_ERR, logMsg, nil)
            t.ServeJSON()
            return
        }
    } else {
        t.DelSession("modify_topic_id")
        // 修改，如果是修改，则应该有id提交
        id, err := t.GetInt("id")
        if err != nil {
            logMsg := fmt.Sprintf("when update topic input id error: %s", err.Error())
            logs.Error(logMsg)
            t.Data["json"] = common.CreateErrResponse(common.RESP_CODE_PARAMS_ERR, logMsg, nil)
            t.ServeJSON()
            return
        }
        idInner := idData.(int)
        if idInner != id {
            logMsg := fmt.Sprintf("input id`%d` no equal with session's id``%d", id, idInner)
            logs.Error(logMsg)
            t.Data["json"] = common.CreateErrResponse(common.RESP_CODE_PARAMS_ERR, logMsg, nil)
            t.ServeJSON()
            return
        }

        tp := new(db.Topic)
        tp.Id = id
        err = tp.Read("id")
        if err != nil {
            logMsg := fmt.Sprintf("select for `%d` error: %s", id, err.Error())
            logs.Error(logMsg)
            t.Data["json"] = common.CreateErrResponse(common.RESP_CODE_PARAMS_ERR, logMsg, nil)
            t.ServeJSON()
            return
        }
        tp.Name = name
        tp.Slug = slugName
        tp.Desc = desc
        tp.Updated = time.Now().Unix()
        tp.Status = common.STATUS_VALID
        err = tp.Update("name", "slug", "desc", "updated", "status")
        if err != nil {
            logMsg := fmt.Sprintf("update for `%d` error: %s", id, err.Error())
            logs.Error(logMsg)
            t.Data["json"] = common.CreateErrResponse(common.RESP_CODE_SYSTEM_ERR, logMsg, nil)
            t.ServeJSON()
            return
        }
    }

    t.Data["json"] = common.CreateOkResponse(nil)
    t.ServeJSON()
    return
}