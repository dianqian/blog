package manage

import (
    "blog/controllers/base"
    "github.com/astaxie/beego/logs"
    "fmt"
    "net/http"
    "blog/models"
)


type ArticleEditController struct {
    base.AdminCommonCtr
}

/**
 article相关的topic
 */
type TopicOfArticle struct {
    ID          int
    Name        string
    Checked     bool                    // 是否被当前文章选中
}

/**
 获取topics
 */
func (a *ArticleEditController) GetTopics() []*TopicOfArticle {
    var tpOfarticle []*TopicOfArticle
    tp := new(models.Topic)
    tps, err := tp.SelectAll()
    if err == nil {
        for _, item := range tps {
            tpOfarticle = append(tpOfarticle, &TopicOfArticle{ID: item.Id, Name: item.Name})
        }
    }

    return tpOfarticle
}

/**
 为article设置checked的topic
 */
func (a *ArticleEditController) SetTopicChecked(aID int, tpOfarticle []*TopicOfArticle) []*TopicOfArticle {
    ar := new(models.ArticleTopic)
    ats, err := ar.SelectByArticle(aID)
    if err != nil {
        for _, item := range ats {
            for _, tp := range tpOfarticle {
                if item.TopicId == tp.ID {
                    tp.Checked = true
                    break
                }
            }
        }
    }

    return tpOfarticle
}

/**
 article相关的tag
 */
type TagOfArticle struct {
    ID          int
    Name        string
}

/**
 进入文章编辑、新建文章的界面
 */
func (a *ArticleEditController) Get ()  {
    // 预加载执行
    a.AdminBase()

    tpOfarticle := a.GetTopics()

    articleID, err := a.GetInt("articleid")
    if err != nil || articleID == 0 {
        // todo: 表示新建文章
    } else {
        // todo: 编辑文章
        tpOfarticle = a.SetTopicChecked(articleID, tpOfarticle)
    }
    //tag := models.Tag{}
    //tags, err := tag.SelectAll()

    // 专题内容准备
    a.Data["Series"] = tpOfarticle

    //this.Data["Tags"] =

    // 加载template
    a.TplName = "admin/article.html"
    return
}

/**
 文章提交的操作
 todo: 参考(https://beego.me/docs/mvc/controller/flash.md)flash数据处理的方式，表单提交错误的处理方式
 */
func (a *ArticleEditController) Post()  {
    var err error
    // 预加载
    a.AdminBase()

    // 动作方式的处理
    do := a.GetString("do")
    logs.Debug(fmt.Sprintf("action=%s", do))
    // 处理方式，利用defer在最后执行
    defer func() {
        if do == "save" {
            if err != nil {
                a.ServeJSON()
                a.StopRun()
                return
            }
            a.Redirect("/admin/manage-drafts", http.StatusFound)
        } else if do == "publish" {
            a.Redirect("/admin/manage-articles", http.StatusFound)
        } else if do == "auto" {
            // todo: 自动保存的处理
        }
    }()

    // 数据获取
    title := a.GetString("title")
    slug := a.GetString("slug")
    text := a.GetString("text")
    date := a.GetString("date")
    serie := a.GetString("serie")
    tags := a.GetString("tags")
    isUpdate := a.GetString("update")
    if title == "" || slug == "" || text == "" {
        err = fmt.Errorf("param error")
        logs.Error(fmt.Sprintf("param error: title"))
        return
    }
    logs.Debug(fmt.Sprintf("title: %s", title))
    logs.Debug(fmt.Sprintf("slug: %s", slug))
    logs.Debug(fmt.Sprintf("text: %s", text))
    logs.Debug(fmt.Sprintf("date: %s", date))
    logs.Debug(fmt.Sprintf("serie: %s", serie))
    logs.Debug(fmt.Sprintf("tags: %s", tags))
    logs.Debug(fmt.Sprintf("update: %s", isUpdate))


    return
}