package base

import (
    "github.com/astaxie/beego"
    "nest/models/html/htmlvisitor"
    "strings"
    "nest/models/db"
    "net/url"
    "fmt"
    "nest/common"
)

/**
 blog业务Controller的基础公共类，利用继承的特性
 */
type HomeCommonCtr struct {
    CommonCtr
}

/**
 预处理
 */
func (h *HomeCommonCtr) HomeBase() {
    h.CommonBase()

    var visitHTML htmlvisitor.HTMLVisitData
    if strings.Contains(h.Ctx.Input.URL(), "post") == true {
        // 具体文章页面
        slug, err := h.parseArticleSlug()
        if err == nil {
            originAr := &db.Article{Url: slug}
            err := originAr.Read("url")
            if err == nil {
                u := &db.User{Id: originAr.Author}
                err := u.Read("id")
                if err == nil {
                    visitHTML.Icon = u.Avatar
                    visitHTML.Name = u.NickName
                    visitHTML.Description = u.Signature
                }
            }
        }
    } else {
        // 总博客页面
        blog := new(db.Blogger)
        err := blog.Read()
        if err == nil {
            visitHTML.Icon = blog.Icon
            visitHTML.Name = blog.Title
            visitHTML.Description = blog.SubTitle
        }
    }
    if visitHTML.Icon == "" {
        visitHTML.Icon = "/static/img/jiaoshou.jpg"
    }
    if visitHTML.Name == "" {
        visitHTML.Name = "上善若水"
    }
    if visitHTML.Description == "" {
        visitHTML.Description = "上善若水煮肉片，心乱如麻婆豆腐"
    }

    visitHTML.Version = beego.AppConfig.String("common.version")
    visitHTML.Domain = beego.AppConfig.String(beego.BConfig.RunMode + "::domain")           // 域名
    h.Data["HTMLVisitData"] = visitHTML

    h.Layout = "visit/visit_layout.html"
}


/**
 解析URI得到article的slug的信息
 */
func (h *HomeCommonCtr) parseArticleSlug() (string, error) {

    u, err := url.Parse(h.Ctx.Request.URL.EscapedPath())
    if err != nil {
        return "", err
    }
    inputURI := u.Path

    if strings.HasPrefix(inputURI, common.POST_TAG) != true ||
        strings.HasSuffix(inputURI, common.HTML_TAG) != true {
        return "", fmt.Errorf("input uri invalid: %s", inputURI)
    }

    begin := strings.Index(inputURI, common.POST_TAG) + len(common.POST_TAG)
    end := strings.LastIndex(inputURI, common.HTML_TAG)
    if begin >= end {
        return "", fmt.Errorf("input uri invalid: %s", inputURI)
    }

    return inputURI[begin: end], nil
}