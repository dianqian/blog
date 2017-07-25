package wolagent2server

import (
    "encoding/json"
)

// wolagent和wolserver之间的通信协议
// 走tls的长连接通信，所以不用token的校验

/*******************
 协议消息类型
 */
const (
    HEART_BEAT          = 2  // 心跳包协议
    WOLVERINE_ID        = 3  // 金刚狼上报协议
    CONN_DUPLICATE      = 4  // 连接存在协议
    VERSION             = 5  // 版本信息

    CMD_MSG             = 1  // 执行命令协议
    WEAPON_REQ_MSG      = 6  // plugin的weapon的request信息
    WEAPON_RESP_MSG     = 7  // plugin的weapon的response信息
)

// 一级
type ProtoMsg struct {
    MsgType         int         `json:"msg_type"`
    MsgContent      []byte      `json:"msg_content"`
}

func (this *ProtoMsg) Marshal() []byte {
    //json格式化
    buf, err := json.Marshal(this)
    if err != nil {
        buf = nil
    }

    return buf
}

func (this *ProtoMsg) Unmarshal(buf []byte) bool {
    //json解析
    err := json.Unmarshal(buf, this)
    if err != nil {
        return false
    }

    return true
}

// 二级
// cmd消息协议
type CmdMsg struct {
    WolId           string          `json:"wol_id"`
    TaskId          string          `json:"task_id"`

    Cmd             string          `json:"cmd"`                // 命令
    Params          string          `json:"params"`             // 命令参数
    TimeOut         int             `json:"time_out"`

    Result string `json:"result"`   // 结果
    Status int    `json:"status"`   // 命令执行的结果
}

func (this *CmdMsg) Marshal() []byte {

    buf, err := json.Marshal(this)
    if err != nil {
        buf = nil
    }

    return buf
}

func (this *CmdMsg) Unmarshal(buf []byte) bool {

    err := json.Unmarshal(buf, this)
    if err != nil {
        return false
    }

    return true
}

// 二级
// wolverine id的上报协议
type WolverineId struct {
    HaveId      bool        `json:"have_id"`        // 标示是否有金刚狼id
    Id          string      `json:"id"`             // 金刚狼id
    IsCreate    bool        `json:"is_create"`      // 标记该wol id是创建or校验通过

    Pid         int         `json:"pid"`
}

func (this *WolverineId) Marshal() []byte {

    buf, err := json.Marshal(this)
    if err != nil {
        buf = nil
    }

    return buf
}

func (this *WolverineId) Unmarshal(buf []byte) bool {

    err := json.Unmarshal(buf, this)
    if err != nil {
        return false
    }

    return true
}

// 版本信息
type VersionInfo struct {
    Version     int     `json:"version"`
    Platform    string  `json:"platform"`           // agent 运行的平台
}

func (this *VersionInfo) Marshal() []byte {

    buf, err := json.Marshal(this)
    if err != nil {
        buf = nil
    }

    return buf
}

func (this *VersionInfo) Unmarshal(buf []byte) bool {

    err := json.Unmarshal(buf, this)
    if err != nil {
        return false
    }

    return true
}


/**
 weapon request msg struct
 weaponId/name/version/md5/url/type      wolId/taskId/
 */
type WeaponReqMsg struct {
    WolId           string          `json:"wol_id"`
    TaskId          string          `json:"task_id"`

    WeaponId        int             `json:"weapon_id"`
    WeaponName      string          `json:"weapon_name"`
    WeaponMd5       string          `json:"weapon_md_5"`
    WeaponUrl       string          `json:"weapon_url"`
    WeaponType      int             `json:"weapon_type"`

    TimeOut         int             `json:"time_out"`
    InputParams     string          `json:"input_params"`             // weapon执行时，需传入的参数
}

func (this *WeaponReqMsg) Marshal() []byte {

    buf, err := json.Marshal(this)
    if err != nil {
        buf = nil
    }

    return buf
}

func (this *WeaponReqMsg) Unmarshal(buf []byte) bool {

    err := json.Unmarshal(buf, this)
    if err != nil {
        return false
    }

    return true
}

/**
 weapon response msg  struct
 */
type WeaponRespMsg struct {
    WolId       string      `json:"wol_id"`
    TaskId      string      `json:"task_id"`
    WeaponId    int         `json:"weapon_id"`
    Status      int         `json:"status"`
    Input       string      `json:"input"`
    Output      string      `json:"output"`
}

func (this *WeaponRespMsg) Marshal() []byte {

    buf, err := json.Marshal(this)
    if err != nil {
        buf = nil
    }

    return buf
}

func (this *WeaponRespMsg) Unmarshal(buf []byte) bool {

    err := json.Unmarshal(buf, this)
    if err != nil {
        return false
    }

    return true
}