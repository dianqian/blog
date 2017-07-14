package models

/**
  文章的统计信息
 */
type ArticleStat struct {
    Id          int

    ArticleId       int                     // 文章id，是ArticleInfo中的Id值
    CommentCount    int                     // 文章被评论条数
    FavourCount     int                     // 文章被点赞的次数

    Create      int64
    Updated     int64
    Status      int
}
