package goroutine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurGoroutineID(t *testing.T) {
	it := assert.New(t)

	ch := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			goroutineID := CurGoroutineID()

			t.Logf("goroutineID:%d\n", goroutineID)

			it.True(goroutineID > 0)

			ch <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-ch
	}
}
