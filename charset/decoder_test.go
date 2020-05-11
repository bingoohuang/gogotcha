package charset

import (
	"fmt"
	"testing"

	"github.com/bingoohuang/gou/file"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html/charset"
)

func TestGBK(t *testing.T) {
	s, err := file.ToString("testdata/gbk.txt")
	assert.Nil(t, err)

	encoding, name, certain := charset.DetermineEncoding([]byte(s), "")
	fmt.Println(encoding, name, certain)

	decoded, err := NewDecoder("GBK").Decode(s)
	assert.Nil(t, err)
	assert.Equal(t, "中华人民共和国合同法", decoded)

	s, err = file.ToString("testdata/utf8.txt")
	assert.Nil(t, err)

	encoding, name, certain = charset.DetermineEncoding([]byte(s), "")
	fmt.Println(encoding, name, certain)

	decoded, err = NewDecoder("UTF8").Decode(s)
	assert.Nil(t, err)
	assert.Equal(t, "中华人民共和国合同法", decoded)

	s, err = file.ToString("testdata/GB18030.txt")
	assert.Nil(t, err)

	encoding, name, certain = charset.DetermineEncoding([]byte(s), "")
	fmt.Println(encoding, name, certain)

	decoded, err = NewDecoder("GB18030").Decode(s)
	assert.Nil(t, err)
	assert.Equal(t, "中华人民共和国合同法", decoded)
}
