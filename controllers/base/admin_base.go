package base


/**
 admin管理的Controller的公共基础类
 */
type AdminCommonCtr struct {
    CommonCtr
}

func (this *AdminCommonCtr) Prepare() {
    this.CommonBase()

    this.Layout = "admin/admin_layout.html"
}
