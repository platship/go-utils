/*
 * @Author: coller coller@88.com
 * @Date: 2024-04-21 15:44:25
 * @LastEditors: coller coller@88.com
 * @LastEditTime: 2024-04-21 15:48:12
 * @FilePath: /go-utils/htmlx/htmlx.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package htmlx

import (
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
