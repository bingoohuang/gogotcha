package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsChinese(t *testing.T) {
	assert.True(t, HasChinese("This is China,这是中国"))
}

func BenchmarkIsChinese(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HasChinese("This is China,这是中国")
	}
}
