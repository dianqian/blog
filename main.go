package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/astaxie/beego/logs"
	"blog/models"
	"blog/common"
)

func init() {
	// db init
	// fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", this.User, this.Passwd, this.Ip, this.Port, this.Database)
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		beego.AppConfig.String("db.username"),
		beego.AppConfig.String("db.password"),
		beego.AppConfig.String("db.domain"),
		beego.AppConfig.String("db.port"),
		beego.AppConfig.String("db.database"),
	)
	logs.Info("data source:" + dataSource)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dataSource, 30)
	orm.RegisterModel(
		//new(models.User),
		//new(models.TmpUser),
		//new(models.ArticleInfo),
		//new(models.ArticleContent),
		//new(models.ArticleStat),
		//new(models.Comment),
		//new(models.CommentStat),
		//new(models.Category),
		//new(models.TagInfo),
		new(models.Social),
	)
	orm.Debug = false
	orm.RunSyncdb("default", false, true)
}

func main() {
	// 使能session
	beego.BConfig.WebConfig.Session.SessionOn = true

	// 自定义在模板中使用的函数
	// 使用《模板函数》特性：https://beego.me/docs/mvc/view/template.md
	beego.AddFuncMap("TimeIsZero", common.TimeIsZero)

	beego.Run()
}

