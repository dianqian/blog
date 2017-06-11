package models

/**
 评论数据结构
 */
type Comment struct {
    Id          int                     // 评论id

    UserType    int                     // 评论的用户类型，注册用户 or  匿名用户，注册用户：1，匿名用户：2
    UserId      int                     // 评论的用户Id，是User或者TmpUser的Id值

    ArticleId   int                     // 文章id，是ArticleInfo中的Id值
    Content     string                  // 评论内容

    Create      int64
    Updated     int64
    Status      int
}

/**
 评价的统计信息
 */
type CommentStat struct {
    Id          int

    CommentId       int         // 评价的id值，是Comment中Id值
    FavourCount     int         // 被👍的次数
    LowCount        int         // 被🌶黑的次数

    Create          int64
    Updated         int64
    Status          int
}