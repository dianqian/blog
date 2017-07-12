package models

/**
 文章的基本信息
 */
type ArticleInfo struct {
    Id          int

    Title       string                  // 标题
    SubTitle    string                  // 子标题
    Author      string                  // 作者
    Tags        string                  // 标签
    Categories  string                  // 分类

    Create      int64
    Updated     int64
    Status      int
}

/**
 文章内容信息
 */
type ArticleContent struct {
    Id          int
    ArticleId   int                     // 文章id，是ArticleInfo中的Id值
    Content     string                  // 文章内容，采用markdown方式，longText的类型

    Create      int64
    Updated     int64
    Status      int
}

/**
 分类信息
 */
type Category struct {
    Id          int

    SortId      int                     // 排序值
    Name        string                  // category的名字
    Count       int                     // 该category文章的数量
    Description string                  // 描述信息

    Create      int64
    Update      int64
    Status      int
}


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