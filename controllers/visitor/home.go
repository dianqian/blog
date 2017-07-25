package visitor

import (
	"github.com/astaxie/beego/logs"
	"blog/controllers/base"
	"fmt"
	"blog/models"
	"github.com/astaxie/beego"
	"blog/common"
	"time"
)

type ArticleDesc struct {
	ID                      int
	Title                   string
	Excerpt                 string

	PublishTime             time.Time
	Slug                    string

	CommentCount            int
}

type HomeController struct {
	base.HomeCommonCtr
}

/**
 首页的get获取展示操作
 */
func (h *HomeController) Get() {
	// 基础的执行
	h.HomeBase()

	h.Data["Website"] = "beego.me"
	h.Data["Email"] = "astaxie@gmail.com"

	// 首页文章列表
	a := new(models.Article)
	cnt, err := a.Count(100)
	if err != nil {
		// todo: 错误处理
		logs.Error(fmt.Sprintf("home article count failed: %s", err.Error()))
	}

	// 指定页码
	pn, err := h.GetInt("pn")
	if err != nil {
		logs.Debug(fmt.Sprintf("no input page number"))
		pn = 1
	}
	// 每页的值
	pageSize, err := beego.AppConfig.Int(beego.AppConfig.String("runmode") + "::page_size")
	if err != nil {
		logs.Debug(fmt.Sprintf("no page size config."))
		pageSize = 10
	}

	// 根据page得到数据
	page := common.PageUtil(int(cnt), pn, int(pageSize), nil)
	articles, err := a.Select(page.Offset, page.PageSize, 100)
	if err != nil {
		logs.Error(fmt.Sprintf("article select(offset=%d, limit=%d) failed: %s",
			page.Offset, page.PageSize, err.Error()))
	}

	var arDesc []ArticleDesc
	for _, item := range articles {
		one := ArticleDesc{
			ID: item.Id,
			PublishTime: time.Unix(item.PublishTime, 0),
			Slug: item.Url,
			Excerpt: "",                    // todo: 描述待实现
			Title: item.Title,
		}
		arDesc = append(arDesc, one)
	}
	h.Data["List"] = arDesc

	h.Data["Prev"] = page.PageNo - 1
	h.Data["Next"] = page.PageNo + 1

	h.TplName = "visitor/home.html"
	return
}

