package visitor

import (
    "blog/controllers/base"
    "time"
    "bytes"
    "fmt"
    "blog/models"
)

type ArchiveController struct {
    base.HomeCommonCtr
}

func (a *ArchiveController) Get()  {
    // 基础的执行
    a.HomeBase()

    // 准备数据
    //ap := new(ArchivePage)
    //err :=


    a.Data["Header"] = ""
    a.Data["Content"] = ""

    a.TplName = "visitor/archives.html"
    return
}

/**
 archive pre say setting
 */
func (a *ArchiveController) getPreSay(ap *ArchivePage) error {
    blog := new(models.Blogger)
    err := blog.Read()
    if err != nil {
        return fmt.Errorf("read blog info failed: %s", err.Error())
    }

    ap.Name = "归档"
    ap.PreSay = blog.ArchivesSay
    return nil
}

/**
 archive deal setting, and get articles
 */
func (a *ArchiveController) getArchivesAndArticles(ap *ArchivePage) error {
    // 取出所有的文章
    ar := new(models.Article)
    ars, err := ar.SelectByStatus(100)
    if err != nil {
        return fmt.Errorf("find article error: %s", err.Error())
    }

    ars = ars
    //for _, item := range ars {
    //
    //}


    return nil
}


/**
 用于生成markdown的数据结构
 */
type ArchivePage struct {
    Name        string
    PreSay      string
    Archives    []*Archive
}

type Archive struct {
    Date                time.Time           // 日期
    ArticleBrief        []*ArticleBriefForArchive
}

type ArticleBriefForArchive struct {
    Name            string              // 文章名
    Slug            string              // 文章访问的url
    PublishTime     time.Time
}

/**
 基于指定内容创建markdown的内容
 生成archive的markdown
 */
func (ap *ArchivePage) generateArchiveMarkdown() string {
    var buffer bytes.Buffer

    // 前说
    buffer.WriteString(fmt.Sprintf("# %s", ap.Name))
    buffer.WriteString("\n")
    buffer.WriteString(ap.PreSay)
    buffer.WriteString("\n\n")

    for _, item := range ap.Archives {
        buffer.WriteString(fmt.Sprintf("## %s", item.Date.Format("2006年01月")))
        buffer.WriteString("\n\n")
        for _, ar := range item.ArticleBrief {
            // list方式的markdown语法： * [标题](/post/slug.html) <span class="date">(Man 02, 2006)</span>
            buffer.WriteString(fmt.Sprintf("* [%s](/post/%s.html) <span class=\"date\">(%s)</span>", ar.Name, ar.Slug, ar.PublishTime.Format("Jan 02, 2006")))
        }
        buffer.WriteString("\n\n")
    }

    return string(buffer.Bytes())
}
