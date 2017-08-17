package manage

import (
    "blog/controllers/base"
    "blog/common"
    "github.com/astaxie/beego/logs"
    "fmt"
)

type DraftManageController struct {
    base.AdminCommonCtr
}

func (d *DraftManageController) Get ()  {
    // 基础调用函数
    d.AdminBase()

    a := ArticleTitleInfo{}
    drafts, err := a.GetTitleInfos(common.ARTICLE_STATUS_DRAFT)
    if err != nil {
        logs.Error(fmt.Sprintf("draft GetTitleInfos failed: %s", err.Error()))
        d.TplName = "admin/drafts.html"
        return
    }

    d.Data["List"] = drafts
    d.TplName = "admin/drafts.html"
    return
}