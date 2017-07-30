package visitor

import (
    "blog/controllers/base"
    "time"
    "bytes"
    "fmt"
    "blog/models"
    "github.com/astaxie/beego/logs"
    "blog/common"
)

type ArchiveController struct {
    base.HomeCommonCtr
}

func (a *ArchiveController) Get()  {
    // 基础的执行
    a.HomeBase()

    // 准备数据
    ap := new(ArchivePage)
    err := a.getPreSay(ap)
    if err != nil {
        logs.Error(fmt.Sprintf("get archive presay failed: %s", err.Error()))
    }
    err = a.getArchivesAndArticles(ap)
    if err != nil {
        logs.Error(fmt.Sprintf("get archive and article failed: %s", err.Error()))
    }

    // 转换为markdown的text
    a.Data["Header"], a.Data["Content"] = common.GetHeaderAndContent([]byte(ap.generateArchiveMarkdown()))

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
    ars, err := ar.Select(0, 1000, 100)
    if err != nil {
        return fmt.Errorf("find article error: %s", err.Error())
    }

    // 遍历文章
    for _, item := range ars {
        //logs.Error(item.PublishTime)
        ay, am, _ := time.Unix(item.PublishTime, 0).Date()            // 文章的日期
        isMatch := false
        // 遍历archive
        for _, one := range ap.Archives {
            if ay == one.Year && am == one.Month {
                // 匹配到archive
                one.ArticleBrief = append(one.ArticleBrief, &ArticleBriefForArchive{
                    Name: item.Title, Slug: item.Url, PublishTime: time.Unix(item.PublishTime, 0),
                })
                isMatch = true

                break                           // 跳出该循环
            }
        }

        // 没有匹配，则新增一个Archive
        if isMatch == false {
            oneAr := &Archive{
                Date: time.Unix(item.PublishTime, 0),
                Year: ay, Month: am,
            }
            oneAr.ArticleBrief = append(oneAr.ArticleBrief, &ArticleBriefForArchive{
                Name: item.Title, Slug: item.Url, PublishTime: time.Unix(item.PublishTime, 0),
            })

            ap.Archives = append(ap.Archives, oneAr)
        }
    }

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
    Year                int                 // 年
    Month               time.Month          // 月
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
            buffer.WriteString("\n")
        }
        buffer.WriteString("\n\n")
    }

    return string(buffer.Bytes())
}
