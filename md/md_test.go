package md

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/gomarkdown/markdown"
)

func TestMdHtml(t *testing.T) {
	md := []byte(`
## 运动员名单

国家 | 姓名
---|---
中国|孙杨
美国|菲利普斯


	`)
	output := markdown.ToHTML(md, nil, nil)
	fmt.Println(string(output))
	ioutil.WriteFile("md.html", output, 0777)
}
