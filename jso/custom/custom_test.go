package custom

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bingoohuang/gogotcha/lang"
)

func TestCustom(t *testing.T) {
	data := []byte(`{"id": "foo"}`)
	item := Item{}
	err := json.Unmarshal(data, &item)

	fmt.Println("err: ", err)
	fmt.Println("item: ", item)

	item = Item{}
	err = lang.Unmarshal(data, &item)
	fmt.Println("err: ", err)
	fmt.Println("item: ", item)
}
