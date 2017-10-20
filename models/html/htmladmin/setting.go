package htmladmin

import "time"

/**
 账号信息
 */
type Account struct {
    Name            string
    NickName        string

    Avatar          string
    Signature       string

    Email           string
    WebSite         string
    Wechat          string

    CreateTime          time.Time
    LastLoginTime       time.Time
    LastLogoutTime      time.Time

    LoginIp         string
}

type HTMLUserData struct {
    ErrorInfo       string
    User            Account
}

/**
 blog基本信息配置相关
 */
type Blogger struct {
    Title           string          // 标题
    SubTitle        string          // 子标题
    Icon            string          // 图标

    BeiAn           string          // 备案号
    Copyright       string

    SeriesSay       string          // 专题前说
    ArchivesSay     string          // 归档前说

    LastUpdateTime  time.Time       // 上次更新的时间
}


type HTMLBlogData struct {
    ErrorInfo       string
    Blogger         Blogger
}