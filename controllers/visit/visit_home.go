package visit

import (
	"github.com/astaxie/beego/logs"
	"nest/controllers/base"
	"fmt"
	"nest/models/db"
	"github.com/astaxie/beego"
	"nest/common"
	"time"
	"nest/models/html/htmlvisitor"
)

type HomeController struct {
	base.HomeCommonCtr
}

/**
 首页的get获取展示操作
 */
func (h *HomeController) Get() {
	// 基础的执行
	h.HomeBase()

	var htmlData htmlvisitor.HTMLHomeData

	h.Data["Website"] = "beego.me"
	h.Data["Email"] = "astaxie@gmail.com"

	// 首页文章列表
	a := new(db.Article)
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
	page := common.PageUtil(int(cnt), pn, int(pageSize))
	articles, err := a.Select(0, page.PageSize, 100)
	if err != nil {
		logs.Error(fmt.Sprintf("article select(offset=%d, limit=%d) failed: %s",
			0, page.PageSize, err.Error()))
	}

	var arDesc []htmlvisitor.ArticleIntro
	for _, item := range articles {
		one := htmlvisitor.ArticleIntro{
			ID: item.Id,
			PublishTime: time.Unix(item.PublishTime, 0),
			Slug: item.Url,
			Excerpt: h.generalExcerpt(item),
			Title: item.Title,
		}
		arDesc = append(arDesc, one)
	}
	htmlData.List = arDesc
	htmlData.Prev = page.PageNo - 1
	htmlData.Next = page.PageNo + 1

	h.Data["HTMLData"] = htmlData
	h.TplName = "visit/visit_home.html"
	return
}


func (h *HomeController) generalExcerpt(a *db.Article) string {
	_, content := common.GetHeaderAndContent([]byte(a.Content))
	return common.GeneralExcerpt(content)
}