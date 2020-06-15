package custom_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bingoohuang/gogotcha/jso/custom"
	"github.com/bingoohuang/gogotcha/lang"
)

func TestCustom(t *testing.T) {
	data := []byte(`{"id": "foo"}`)
	item := custom.Item{}
	err := json.Unmarshal(data, &item)

	fmt.Println("err: ", err)
	fmt.Println("item: ", item)

	item = custom.Item{}
	err = lang.Unmarshal(data, &item)
	fmt.Println("err: ", err)
	fmt.Println("item: ", item)
}
