package routers

import (
    "github.com/astaxie/beego"
    "blog/controllers/visit"
    "blog/controllers"
)

func init() {
    /*********************************
     todo: 其他错误、异常处理相关
     */
    beego.ErrorController(&controllers.ErrorController{})

    /*********************************
     博客页
     */
    // 首页
    beego.Router("/", &visit.HomeController{})

    // 文章显示
    // http://localhost:8080/post/[article.url].html
    beego.Router("/post/:slug", &visit.ArticleController{})
    // todo: 专题
    beego.Router("topics", &visit.Topics{})
    // todo：归档
    beego.Router("archives", &visit.ArchiveController{})
    // todo: 友链
    beego.Router("blogroll", &visit.BlogRoll{})
    // todo: about
    beego.Router("about", &visit.About{})
    // todo: search搜索

    // user相关的页面
    //



    // test
    beego.Router("/test/", &controllers.TestController{})

}