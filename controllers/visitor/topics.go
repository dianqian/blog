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

type Topics struct {
    base.HomeCommonCtr
}

func (t *Topics) Get() {
    // 基础的执行
    t.HomeBase()

    // 准备数据
    tpForMd := new(TopicsInfoForMd)
    err := t.getPreSay(tpForMd)
    if err != nil {
        logs.Error(fmt.Sprintf("get topic presay error, %s", err.Error()))
    }
    err = t.getTopics(tpForMd)
    if err != nil {
        logs.Error(fmt.Sprintf("get topics error, %s", err.Error()))
    }
    err = t.getArticles(tpForMd)
    if err != nil {

    }

    logs.Error(tpForMd.generateTopicsMarkdown())
    t.Data["Header"], t.Data["Content"] = common.GetHeaderAndContent([]byte(tpForMd.generateTopicsMarkdown()))

    t.TplName = "visitor/series.html"
}

/**
 获取专题的前说
 */
func (t *Topics) getPreSay(tpForMd *TopicsInfoForMd) error {
    blog := new(models.Blogger)
    err := blog.Read()
    if err != nil {
        return fmt.Errorf("read blog info failed: %s", err.Error())
    }

    tpForMd.Name = "专题"
    tpForMd.PreSay = blog.SeriesSay
    return nil
}

/**
 获取所有的topic信息
 */
func (t *Topics) getTopics(tpForMd *TopicsInfoForMd) error {
    tp := new(models.Topic)
    tps, err := tp.SelectAll()
    if err != nil {
        return fmt.Errorf("read all topic failed: %s", err.Error())
    }
    for _, item := range tps {
        one := TopicInfo{ID: item.Id, Name: item.Name, Slug: item.Slug, Description: item.Desc}
        tpForMd.Topics = append(tpForMd.Topics, &one)
    }

    // 增加一个默认的专题，这样在后续的get article中可以获取到相关的article
    one := TopicInfo{ID: 0, Name: "默认专题", Slug: "默认专题", Description: "当没有设置专题时，则划归为默认专题"}
    tpForMd.Topics = append(tpForMd.Topics, &one)

    return nil
}

/**
 获取article的基本信息
 */
func (t *Topics) getArticles(tpForMd *TopicsInfoForMd) error {
    for _, item := range tpForMd.Topics {
        arTp := new(models.ArticleTopic)
        aTps, err := arTp.SelectByTp(item.ID)
        if err != nil {
            logs.Error(fmt.Sprintf("no article topics record for topic(%d): %s", item.ID, err.Error()))
            continue
        }

        for _, i := range aTps {
            ar := new(models.Article)
            ar.Id = i.ArticleId
            err := ar.Read("id")
            if err != nil {
                logs.Error(fmt.Sprintf("read article(id=%d) failed: %s", i.ArticleId, err.Error()))
                continue
            }

            ab := ArticleBriefForTopic{Name: ar.Title, Slug: ar.Url, PublishTime: time.Unix(ar.PublishTime, 0)}
            item.ArticleBrief = append(item.ArticleBrief, &ab)
        }
    }
    return nil
}


type TopicsInfoForMd struct {
    Name            string
    PreSay          string                  // 专题前言
    Topics          []*TopicInfo            //
}

type TopicInfo struct {
    ID              int                 // 专题ID值
    Name            string              // 专题名
    Slug            string              // 略缩名
    Description     string              // 专题描述信息
    ArticleBrief    []*ArticleBriefForTopic
    // 专题下面的article信息
}

type ArticleBriefForTopic struct {
    Name            string              // 文章名
    Slug            string              // 文章访问的url
    PublishTime     time.Time
}

/**
 基于指定内容创建markdown的内容
 生成topics的md
 */
func (tpInfo *TopicsInfoForMd) generateTopicsMarkdown() string {
    var buffer bytes.Buffer

    // 专题前说
    buffer.WriteString(fmt.Sprintf("# %s", tpInfo.Name))
    buffer.WriteString("\n")
    buffer.WriteString(tpInfo.PreSay)
    buffer.WriteString("\n\n")

    // 展示不同的专题
    for _, item := range tpInfo.Topics {
        buffer.WriteString(fmt.Sprintf("## %s{#toc-%d}", item.Name, item.ID))
        buffer.WriteString("\n")
        buffer.WriteString(item.Description)
        for _, ar := range item.ArticleBrief {
            // list方式的markdown语法： * [标题](/post/slug.html) <span class="date">(Man 02, 2006)</span>
            buffer.WriteString(fmt.Sprintf("* [%s](/post/%s.html) <span class=\"date\">(%s)</span>", ar.Name, ar.Slug, ar.PublishTime.Format("Jan 02, 2006")))
        }
        buffer.WriteString("\n")
    }

    return string(buffer.Bytes())
}