package pinyinx

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var (
	pyTableLen = len(pyValue)
)
var Ftpinyin = new(Pinyin)

type Pinyin struct {
	//分隔符
	Split string
	//是否首字母大写
	Upper bool
}

func Convert(s, split string) string {
	var pyStr string
	gbk, err := utf8ToGbk([]byte(s))
	if err != nil {
		return pyStr
	}

	pyArr := gbkToPinyin(string(gbk))
	pyStr = strings.Join(pyArr, split)

	return pyStr
}

func gbkToPinyin(gbk string) []string {
	var pyStr []string

	for i := 0; i < len(gbk); i++ {
		p := int(gbk[i])
		if p > 0 && p < 160 {
			pyStr = append(pyStr, string(gbk[i]))
		} else {
			i++
			q := int(gbk[i])
			p = p*256 + q - 65536

			py := pinyinSearch(p)
			if py != "" {
				// if c.Upper == false {
				// 	py = strings.ToLower(py)
				// }
				pyStr = append(pyStr, py)
			}
		}
	}

	return pyStr
}

func utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Search(p int) string {
	if v, ok := tableMap[p]; ok {
		return v
	} else {
		for i := pyTableLen - 1; i >= 0; i-- {
			if pyValue[i] <= p {
				return pyName[i]
			}
		}
	}
	return ""
}
