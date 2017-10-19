package htmladmin

import (
    "nest/common"
    "time"
    "nest/models/db"
)

type Topic struct {
    Id              int
    Name            string          // 名字
    Slug            string          // 略缩名
    Desc            string          // 描述信息
    ArticleCount    int     // 该topic的文章统计数

    Create      int64
    Updated     int64
    Status      int
}

/**
 @Description：topic的简要信息
 */
type TopicBriefInfo struct {
    Id              int
    Name            string
    Slug            string
    ArticleCount    int
    Created         time.Time
    Updated         time.Time
    Status          int
}

/**
 @Description：topic list的html的返回数据
 */
type HTMLTopicsData struct {
    ErrorInfo       string
    PageInfo        common.Page
    TopicInfo       []*TopicBriefInfo
}


/**
 @Description：一个topic的内容
 */
type HTMLOneTopic struct {
    ErrorInfo       string
    Topic           db.Topic
}