package common

import (
    "regexp"
    "github.com/russross/blackfriday"
)

/**
 blackfriday配置文件
 */
const (
    BlackfridayCommonHtmlFlags = 0 |
        blackfriday.HTML_TOC |
        blackfriday.HTML_USE_XHTML |
        blackfriday.HTML_USE_SMARTYPANTS |
        blackfriday.HTML_SMARTYPANTS_FRACTIONS |
        blackfriday.HTML_SMARTYPANTS_DASHES |
        blackfriday.HTML_SMARTYPANTS_LATEX_DASHES |
        blackfriday.HTML_NOFOLLOW_LINKS

    BlackfridayCommonExtensions = 0 |
        blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
        blackfriday.EXTENSION_TABLES |
        blackfriday.EXTENSION_FENCED_CODE |
        blackfriday.EXTENSION_AUTOLINK |
        blackfriday.EXTENSION_STRIKETHROUGH |
        blackfriday.EXTENSION_SPACE_HEADERS |
        blackfriday.EXTENSION_HEADER_IDS |
        blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
        blackfriday.EXTENSION_DEFINITION_LISTS
)

/**
 拆分content为header和content
 */
func GetHeaderAndContent(md []byte) (header string, content string) {
    // markdown转换为html的文档，涉及header/content
    renderer := blackfriday.HtmlRenderer(BlackfridayCommonHtmlFlags, "", "")
    byteContent := blackfriday.Markdown(md, renderer, BlackfridayCommonExtensions)
    regH := regexp.MustCompile("</nav>")
    index := regH.FindIndex(byteContent)
    if index != nil {
        header = string(byteContent[0: index[1]])
        content = string(byteContent[index[1]: ])
    } else {
        header = ""
        content = string(byteContent)
    }

    return
}

/**
 markdown转换为content的string返回
 */
func GetContent(md []byte) (content string) {
    // markdown转换为html的文档
    renderer := blackfriday.HtmlRenderer(BlackfridayCommonHtmlFlags, "", "")
    byteContent := blackfriday.Markdown(md, renderer, BlackfridayCommonExtensions)
    return string(byteContent)
}
