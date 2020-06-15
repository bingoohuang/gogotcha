package lang_test

import (
	"testing"

	"github.com/bingoohuang/gogotcha/lang"
	"github.com/stretchr/testify/assert"
)

func TestIsChinese(t *testing.T) {
	assert.True(t, lang.HasChinese("This is China,这是中国"))
}

func BenchmarkIsChinese(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lang.HasChinese("This is China,这是中国")
	}
}
