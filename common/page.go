package common

type Page struct {
    PageNo     int                  // 页号
    PageSize   int                  // 页中的data size
    TotalPage  int                  // 总页
    TotalCount int                  // 总的数据
    FirstPage  bool                 // 第一页
    LastPage   bool                 // 最后一页

    List       interface{}          // 数据list
    Offset     int                  // 偏移量
}

/**
 生成页结构数据
        count           数据条数
        pageNo          当前页号
        pageSize        页中的data size
        list            数据list
 */
func PageUtil(count int, pageNo int, pageSize int, list interface{}) Page {
    // tp: 计算出总的页数
    tp := count / pageSize
    if count % pageSize > 0 {
        tp = count / pageSize + 1
    }

    page := Page {
        PageNo: pageNo,
        PageSize: pageSize,
        TotalPage: tp,
        TotalCount: count,
        FirstPage: pageNo == 1,
        LastPage: pageNo == tp,
        List: list,
        Offset: (pageNo - 1) * pageSize,
    }
    return page
}