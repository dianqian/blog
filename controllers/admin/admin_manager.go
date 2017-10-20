package admin

import (
    "nest/controllers/base"
    "nest/models/db"
    "github.com/astaxie/beego/logs"
    "fmt"
    "time"
    "nest/models/html/htmladmin"
    "nest/common"
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


type BlogMgController struct {
    base.AdminCommonCtr
}

/**
 @Description：blog相关的配置信息
 @Param:
 @Return：
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

    blogData := htmladmin.HTMLBlogData{
        Blogger: htmladmin.Blogger{
            Title: blog.Title,
            SubTitle: blog.SubTitle,
            //Icon: blog.Icon,
            Icon: "/static/img/jiaoshou.jpg",
            BeiAn: blog.BeiAn,
            Copyright: blog.Copyright,
            SeriesSay: blog.SeriesSay,
            ArchivesSay: blog.ArchivesSay,
            LastUpdateTime: time.Unix(blog.Updated, 0),
        },
    }
    b.Data["HTMLBlogData"] = blogData
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
 @Description：修改配置信息
 @Param:
 @Return：
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
        logMsg := fmt.Sprintf("blog info update failed: %s", err.Error())
        logs.Error(logMsg)
        b.Data["json"] = common.CreateErrResponse(common.RESP_CODE_PARAMS_ERR, logMsg, nil)
        b.ServeJSON()
        return
    }

    b.Data["json"] = common.CreateOkResponse(nil)
    b.ServeJSON()
    return
}