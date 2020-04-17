package hashid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashID(t *testing.T) {
	h := NewHashID("bingoo", 10)

	s := h.Encrypt(1001)
	assert.Equal(t, "vaYWPK1LM8", s)
	assert.Equal(t, []int{1001}, h.Decrypt(s))
}
