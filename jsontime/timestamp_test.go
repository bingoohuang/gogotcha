package jsontime_test

import (
	"encoding/json"
	"github.com/bingoohuang/gogotcha/jsontime"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUnmashalMsg(t *testing.T) {
	j := `{
		"A": "123",
		"B": "2020-03-18 10:51:54.198",
		"C": "2020-03-18 10:51:54,198",
		"d": "2020-03-18T10:51:54.198000Z"
	}`
	msg, err := ParseMsg(j)
	assert.Nil(t, err)
	assert.Equal(t, jsontime.Timestamp(time.Unix(0, 123000)), msg.A)
	p, _ := time.Parse("2006-01-02 15:04:05.000", "2020-03-18 10:51:54.198")
	assert.Equal(t, jsontime.Timestamp(p), msg.B)
	assert.Equal(t, jsontime.Timestamp(p), msg.C)
	assert.Equal(t, jsontime.Timestamp(p), msg.D)
	assert.Equal(t, time.Time(msg.D).Format("20060102150405"), "20200318105154")

}

type Msg struct {
	A jsontime.Timestamp
	B jsontime.Timestamp
	C jsontime.Timestamp
	D jsontime.Timestamp `json:"d"`
}

func ParseMsg(s string) (msg Msg, err error) {
	err = json.Unmarshal([]byte(s), &msg)
	return
}
