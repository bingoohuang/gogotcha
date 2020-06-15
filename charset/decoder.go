package charset

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"

	"io/ioutil"
	"strings"
)

// Decoder contains charset encoding to decode non-utf8 string.
type Decoder struct {
	decoder *encoding.Decoder
}

// NewDecoder create a Decoder by charset name.
func NewDecoder(charset string) *Decoder {
	switch strings.ToUpper(charset) {
	case "GB18030":
		return &Decoder{decoder: simplifiedchinese.GB18030.NewDecoder()}
	case "GBK":
		return &Decoder{decoder: simplifiedchinese.GBK.NewDecoder()}
	case "HZGB2312":
		return &Decoder{decoder: simplifiedchinese.HZGB2312.NewDecoder()}
	case "BIG5":
		return &Decoder{decoder: traditionalchinese.Big5.NewDecoder()}
	default:
		return &Decoder{}
	}
}

// Decode decodes an non-utf8 string.
func (c *Decoder) Decode(s string) (string, error) {
	if c.decoder == nil {
		return s, nil
	}

	all, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(s), c.decoder))
	if err != nil {
		return s, err
	}

	return string(all), nil
}
