package lang

import "unicode"

func HasChinese(str string) bool {
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			return true
		}
	}

	return false
}
