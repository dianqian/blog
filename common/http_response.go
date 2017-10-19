package common

import (
    "net/http"
    "encoding/json"
)

/**
 一些通用的错误码
 */
const (
    RESP_CODE_OK = 1000

    RESP_CODE_PARAMS_ERR = 1002                 // params error
    RESP_CODE_REQ_BODY_ERR = 1003               // body error
    RESP_CODE_JSON_ERR = 1004                   // json error
    RESP_CODE_PARAM_ERR = 1005                  // param error
    RESP_CODE_TOKEN_ERR = 1006                  // token verify

    RESP_CODE_AUTH_ERR = 1009                  // auth fail

    RESP_CODE_SYSTEM_ERR = 1010                 // system error
)

/**
 基于map实现的基础的response data
 */
func CreateOkResponse(data interface{}) map[string]interface{} {
    status := map[string]interface{}{"code": RESP_CODE_OK, "msg": "success"}
    respData := map[string]interface{}{"status": status, "result": data}
    return respData
}

func CreateErrResponse(code int, msg string, data interface{}) map[string]interface{} {
    status := map[string]interface{}{"code": code, "msg": msg}
    respData := map[string]interface{}{"status": status, "result": data}
    return respData
}

/**
 基于类ResponData执行的response操作
 */
type ResponseStatus struct {
    Code            int         `json:"code"`
    Msg             string      `json:"msg"`
}

type ResponseData struct {
    Status          ResponseStatus      `json:"status"`
    Data            interface{}         `json:"data"`
}

/**
 对于结果的解析中，采用mapstructure模块进行解析该模块
 todo：mapstructure模块需要做一个定制处理，待完成
 */
func (data ResponseData)SendResponse(w http.ResponseWriter) error {
    w.Header().Set("Content-Type", "application/json")

    buf, err := json.Marshal(data)
    if err != nil {
        tmpData := ResponseData{
            Status: ResponseStatus{
                Code: RESP_CODE_SYSTEM_ERR,
                Msg: "system marshal error",
            },
        }
        buf, _ := json.Marshal(tmpData)
        w.Write(buf)

        return err
    }

    w.Write(buf)
    return nil
}