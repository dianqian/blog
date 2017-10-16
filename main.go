package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "nest/models/db"
	_ "nest/routers"
	"fmt"
	"github.com/astaxie/beego/logs"
	"nest/common"
)

func main() {
	// 使能session
	beego.BConfig.WebConfig.Session.SessionOn = true

	// 自定义在模板中使用的函数
	// 使用《模板函数》特性：https://beego.me/docs/mvc/view/template.md
	beego.AddFuncMap("TimeIsZero", common.TimeIsZero)

	beego.Run()
}

/**
 @Description：init
 @Param:
 @Return：
 */
func init()  {
	dbInit()
}

/**
 @Description：db 初始化设置
 @Param:
 @Return：
 */
func dbInit() {
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
	orm.Debug = false
	orm.RunSyncdb("default", false, true)
}
