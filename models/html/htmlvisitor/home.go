package htmlvisitor

import (
    "time"
    "nest/common"
)

/**
 @Description: 文章列表
 */
type HTMLHomeData struct {
    ErrorInfo       string
    PageInfo        common.Page
    Articles        []ArticleIntro
}

type ArticleIntro struct {
    ID                      int
    Title                   string
    Excerpt                 string

    PublishTime             time.Time
    Slug                    string

    CommentCount            int
}


/**
 @Description: 文章的具体信息
*/
type HTMLArticleData struct {
    ErrorInfo           string
    ArticleInfo         Article
    Extend              ExtendInfo
    CopyRight           string
}

type Article struct {
    ID              int                     // 文章ID
    Title           string                  // 标题
    Slug            string                  // 略缩信息，即用于url
    Author          string                  // 作者
    Header          string                  // 文章目录
    Content         string                  // 文章内容
    PublishTime     time.Time               // 发布时间
    UpdateTime      time.Time               // 更新时间
    Days            int
}

type ExtendInfo struct {
    Tags            []string                // 标签
    Topic           TopicArticle            // 专题
    CommentCount    int                     // 评论数
    Prev            *ArticleBrother          // 上一篇
    Next            *ArticleBrother          // 下一篇
}

type ArticleBrother struct {
    ID              int
    Title           string
    Slug            string
    PublishTime     time.Time
}

type TopicArticle struct {
    ID              int                     // topic id值
    Name            string                  // topic name值
    Articles        []ArticleBrother        // 同一个topic下文章信息
}

/**
 @Description: 文章的具体信息
*/
type HTMLAbout struct {
    ErrorInfo       string
    About           AboutArticle
}
type AboutArticle struct {
    ID          int
    Title       string
    Header      string              // 头目录
    Content     string
    Slug        string              // 略缩字
}

/**
 @Description: 友链的
 */
type HTMLRollInfo struct {
    ErrorInfo       string
    Roll            RollInfo
}

type RollInfo struct {
    ID          int
    Title       string
    Header      string
    Content     string
}

/**
 @Description: 归档
 */
type HTMLArchive struct {
    ErrorInfo       string
    Header          string
    Content         string
}

/**
 @Description: 专题
 */
type HTMLTopicData struct {
    ErrorInfo       string
    Header          string
    Content         string
}
