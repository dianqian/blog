package htmlvisitor

import "time"

type ArticleIntro struct {
    ID                      int
    Title                   string
    Excerpt                 string

    PublishTime             time.Time
    Slug                    string

    CommentCount            int
}

type HTMLHomeData struct {
    List            []ArticleIntro
    Prev            int
    Next            int
}
