package admin

import (
    "nest/controllers/base"
    "nest/models/db"
    "github.com/astaxie/beego/logs"
    "fmt"
    "nest/models/html/htmladmin"
    "strings"
    "nest/common"
    "time"
)


type TagManageController struct {
    base.AdminCommonCtr
}

/**
 @Description：db 初始化设置
 @Param:
 @Return：
 */
func (t *TagManageController) Get ()  {
    // 基础执行
    t.AdminBase()

    tgsData := htmladmin.HTMLTagsData{}

    // 准备数据
    tg := new(db.Tag)
    tgs, err := tg.SelectAll()
    if err != nil {
        logMsg := fmt.Sprintf("tag select all failed: %s.", err.Error())
        logs.Error(logMsg)
        tgsData.ErrorInfo = logMsg
    } else {
        var tmp []string
        for _, item := range tgs {
            tmp = append(tmp, item.Name)
        }
        tgsData.Tags = strings.Join(tmp, ",")
    }

    t.Data["HTMLTagsData"] = tgsData
    t.TplName = "admin/admin_tags.html"
    return
}

/**
 @Description：tag管理的提交
 @Param:
 @Return：
 */
func (t *TagManageController) Post() {
    tg := new(db.Tag)
    tgs, err := tg.SelectAll()
    if err != nil {
        logMsg := fmt.Sprintf("select tags error: %s", err.Error())
        logs.Error(logMsg)
        t.Data["json"] = common.CreateErrResponse(common.RESP_CODE_SYSTEM_ERR, logMsg, nil)
        t.ServeJSON()
        return
    }

    tmp := t.GetString("tags")
    inputTags := strings.Split(tmp, ",")

    // 执行删除
    for _, in := range tgs {
        isMartch := false
        for _, out := range inputTags {
            if strings.TrimSpace(out) == "" {
                continue
            }
            if in.Name == out {
                isMartch = true
                break
            }
        }
        if isMartch == false {
            tg.Id = in.Id
            err = tg.Delete()
            if err != nil {
                logs.Error(fmt.Sprintf("delete tag`%s` error: %s", in.Name, err.Error()))
            }
        }
    }

    // insert操作
    for _, out := range inputTags {
        if strings.TrimSpace(out) == "" {
            continue
        }

        isMartch := false
        for _, in := range tgs {
            if out == in.Name {
                isMartch = true
                break
            }
        }
        if isMartch == false {
            newTag := &db.Tag{
                Name: out,
                Create: time.Now().Unix(),
                Status: common.STATUS_VALID,
            }
            newTag.Updated = newTag.Create
            err = newTag.Insert()
            if err != nil {
                logs.Error(fmt.Sprintf("add tag`%s` error: %s", out, err.Error()))
            }
        }
    }

    t.Data["json"] = common.CreateOkResponse(nil)
    t.ServeJSON()
    return
}

/**
 @Description：db 初始化设置
 @Param:
 @Return：
 */
func (t *TagManageController) TagsList()  {
    t.AdminBase()

    tg := new(db.Tag)
    tgs, err := tg.SelectAll()
    if err != nil {
        t.Data["json"] = []string{}
    } else {
        var tmp []string
        for _, item := range tgs {
            tmp = append(tmp, item.Name)
        }
        t.Data["json"] = tmp
    }

    t.ServeJSON()
    return
}