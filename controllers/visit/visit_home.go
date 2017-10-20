package visit

import (
	"github.com/astaxie/beego/logs"
	"nest/controllers/base"
	"fmt"
	"nest/models/db"
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

	// 解析参数
	pageNo, err := h.GetInt("page_no")
	if err != nil {
		pageNo = common.PAGE_NO_DEFAULT
	}
	pageSize, err := h.GetInt("page_size")
	if err != nil {
		pageSize = common.PAGE_SIZE_DEFAULT
	}
	offset, limit := common.GetOffsetLimit(pageNo, pageSize)

	var htmlData htmlvisitor.HTMLHomeData
	h.Data["Website"] = "beego.me"
	h.Data["Email"] = "astaxie@gmail.com"

	// 首页文章列表
	a := new(db.Article)
	cnt, err := a.Count(common.ARTICLE_STATUS_PUBLISH)
	if err != nil {
		logMsg := fmt.Sprintf("home article count failed: %s", err.Error())
		logs.Error(logMsg)
		htmlData.ErrorInfo = logMsg
		h.Data["HTMLHomeData"] = htmlData
		h.TplName = "visit/visit_home.html"
		return
	}

	articles, err := a.Select(offset, limit, common.ARTICLE_STATUS_PUBLISH)
	if err != nil {
		logMsg := fmt.Sprintf("article select(offset=%d, limit=%d) failed: %s", offset, limit, err.Error())
		logs.Error(logMsg)
		htmlData.ErrorInfo = logMsg
		h.Data["HTMLHomeData"] = htmlData
		h.TplName = "visit/visit_home.html"
		return
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
	htmlData.Articles = arDesc
	htmlData.PageInfo = common.PageUtil(int(cnt), pageNo, pageSize)

	h.Data["HTMLHomeData"] = htmlData
	h.TplName = "visit/visit_home.html"
	return
}


func (h *HomeController) generalExcerpt(a *db.Article) string {
	_, content := common.GetHeaderAndContent([]byte(a.Content))
	return common.GeneralExcerpt(content)
}