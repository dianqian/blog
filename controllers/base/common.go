package base

import (
    "time"
    "encoding/json"
    "github.com/astaxie/beego/logs"
    "io"
)

/**
 整个站点的功能类
 */
type CommonCtr struct {
    LoginBaseCtr
}

func (c *CommonCtr) CommonBase() {
    c.Data["Title"] = "我的蘑菇鸡"

    /* 获取 */
    c.Data["CopyTime"] = time.Now().Year()
}

/**
 json格式返回的数据
 */
func (c *CommonCtr) ResponseJSON(code int, msg string, data interface{})  {
    jsonData := make(map[string]interface{}, 3)

    jsonData["code"] = code
    jsonData["msg"] = msg
    jsonData["data"] = data

    returnJSON, err := json.Marshal(jsonData)
    if err != nil {
        logs.Error(err.Error())
    }

    c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
    io.WriteString(c.Ctx.ResponseWriter, string(returnJSON))
    c.StopRun()
}
