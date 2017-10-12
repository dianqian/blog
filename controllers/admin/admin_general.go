package admin

import (
    "blog/controllers/base"
    "blog/models/db"
    "github.com/astaxie/beego/logs"
    "fmt"
    "time"
    "net/http"
)


type GeneralSettingController struct {
    base.AdminCommonCtr
}

func (g *GeneralSettingController) Get ()  {
    //this.PreBase()

    g.Prepare()
    g.TplName = "admin/admin_general.html"
    //this.TplName = "admin/example.html"
    return
}

/**
 blog基本信息配置相关
 */
type Blogger struct {
    Title           string          // 标题
    SubTitle        string          // 子标题
    Icon            string          // 图标

    BeiAn           string          // 备案号
    Copyright       string

    SeriesSay       string          // 专题前说
    ArchivesSay     string          // 归档前说

    LastUpdateTime  time.Time       // 上次更新的时间
}

type BlogMgController struct {
    base.AdminCommonCtr
}

/**
 展示blog信息
 */
func (b *BlogMgController)Get()  {
    b.AdminBase()

    blog := new(db.Blogger)
    err := blog.Read()
    if err != nil {
        logs.Info(fmt.Sprintf("no blog setting data, will create: %s", err.Error()))

        blog.Create = time.Now().Unix()
        blog.Updated = time.Now().Unix()
        err = blog.Insert()
        if err != nil {
            logs.Error(fmt.Sprintf("create blog setting failed: %s", err.Error()))
        }

        // 重新获取一次
        blog.Read()
    }

    bd := Blogger{
        Title: blog.Title,
        SubTitle: blog.SubTitle,
        //Icon: blog.Icon,
        Icon: "/static/img/jiaoshou.jpg",
        BeiAn: blog.BeiAn,
        Copyright: blog.Copyright,
        SeriesSay: blog.SeriesSay,
        ArchivesSay: blog.ArchivesSay,
        LastUpdateTime: time.Unix(blog.Updated, 0),
    }
    b.Data["Blogger"] = bd
    b.TplName = "admin/admin_manager.html"
    return
}

/**
 修改icon
 */
func (b *BlogMgController) ChangeIcon()  {
    // todo: 待实现
}

/**
 修改配置信息
 */
func (b *BlogMgController) ChangeInfo()  {
    blog := new(db.Blogger)
    blog.Title = b.GetString("title")
    blog.SubTitle = b.GetString("subtitle")
    blog.BeiAn = b.GetString("beian")
    blog.SeriesSay = b.GetString("seriessay")
    blog.ArchivesSay = b.GetString("archivessay")

    blog.Updated = time.Now().Unix()
    err := blog.Update("title", "sub_title", "bei_an", "series_say", "archives_say", "updated")
    if err != nil {
        logs.Error(fmt.Sprintf("blog info update failed: %s", err.Error()))
    }

    b.Redirect("/admin/manager.html", http.StatusFound)
    return
}