package controllers

import (
    "github.com/astaxie/beego"
    "io/ioutil"
    "encoding/json"
)

type TestController struct {
    beego.Controller
}

func (this *TestController) Get() {
    resultData := make(map[string]interface{})
    resultData["get"] = "this is get do"
    this.Data["json"] = resultData

    this.ServeJSON()
    return
}

/**
 todo: post提交数据的几种方式：
 https://beego.me/docs/mvc/controller/params.md
 */
func (this *TestController) Post() {
    resultData := make(map[string]interface{})

    // todo: 方式1：基于key的方式获取post参数
    //input := this.GetString("name")

    // todo: 方式2：基于body直接解析post json数据
    body, err := ioutil.ReadAll(this.Ctx.Request.Body)
    if err != nil {
        this.Data["json"] = resultData
        this.ServeJSON()
        return
    }

    type TurtleInputJson struct {

    }
    input := new(TurtleInputJson)
    err = json.Unmarshal(body, input)
    if err != nil {
        this.Data["json"] = resultData
        this.ServeJSON()
        return
    }

    resultData["get"] = input
    this.Data["json"] = resultData

    this.ServeJSON()
    return
}

