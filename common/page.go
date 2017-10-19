package common

/**
 缺省的page的信息
 */
const (
    PAGE_NO_DEFAULT = 1             // 缺省的页码
    PAGE_SIZE_DEFAULT = 1          // 缺省页条目size
)

/**
 基于传入的值，算出offset和limit
 */
func GetOffsetLimit(pageNo int, pageSize int) (int, int) {
    var offset int

    if pageNo <= 0 {
        offset = 0
    } else {
        offset = (pageNo - 1) * pageSize
    }

    return offset, pageSize
}

type Page struct {
    PageNo     int                  // 页号
    PageSize   int                  // 页中的data size
    TotalPage  int                  // 总页
    TotalCount int                  // 总的数据
    PrePage    bool
    PrePageNo  int
    NextPage   bool
    NextPageNo int
}

/**
 生成页结构数据
        count           数据条数
        pageNo          当前页号
        pageSize        页中的data size
        list            数据list
 */
func PageUtil(count int, pageNo int, pageSize int) Page {
    // tp: 计算出总的页数
    tp := count / pageSize
    if count % pageSize > 0 {
        tp = count / pageSize + 1
    }
    if pageNo <= 0 {
        pageNo = 1
    }

    page := Page {
        PageNo: pageNo,
        PageSize: pageSize,
        TotalPage: tp,
        TotalCount: count,
    }
    if pageNo > 1 {
        page.PrePage = true
        page.PrePageNo = pageNo - 1
    }
    if pageNo < page.TotalPage {
        page.NextPage = true
        page.NextPageNo = pageNo + 1
    }

    return page
}