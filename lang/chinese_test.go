package lang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var str = "This is China,这是中国"

func TestIsChinese(t *testing.T) {
	assert.True(t, HasChinese(str))
}

func BenchmarkIsChinese(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HasChinese(str)
	}
}
