package htmlx

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// TextFromHTML 获取html的纯文本
func TextFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	return doc.Text()
}

// 去除字符串中的html标签
func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "")
	src = strings.Replace(src, "&nbsp;", "", 1) // 去除空格符
	return strings.TrimSpace(src)
}

func TrimLineBreaks(str string) string {
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\"", "")
	str = strings.Replace(str, "&nbsp;", "", 1) // 去除空格符
	return str
}
