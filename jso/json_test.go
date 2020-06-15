package jso_test

import (
	"encoding/json"
	"fmt"

	v1 "github.com/bingoohuang/gogotcha/jso/v1"
	v2 "github.com/bingoohuang/gogotcha/jso/v2"
)

// nolint:govet
func ExampleNewAwesomeV1() {
	awesome := v1.NewAwesome("123456789", "Total awesomeness", 9.99, true)

	awesomeJSON, _ := awesome.ToJSON(false)
	fmt.Println(string(awesomeJSON))

	moreAwesomeJSON, _ := awesome.ToJSON(true)
	fmt.Println(string(moreAwesomeJSON))

	// Output:
	// {"id":"123456789","message":"Total awesomeness","score":9.99,"confirmed":true}
	// {
	//   "id": "123456789",
	//   "message": "Total awesomeness",
	//   "score": 9.99,
	//   "confirmed": true
	// }
}

// nolint:govet
func ExampleNewAwesomeV2() {
	awesome := v2.NewAwesome("123456789", "Total awesomeness", 9.99, true)

	awesomeJSON, _ := awesome.ToJSON(false)
	fmt.Println(string(awesomeJSON))

	moreAwesomeJSON, _ := awesome.ToJSON(true)
	fmt.Println(string(moreAwesomeJSON))

	// Output:
	// {"id":"123456789","message":"Total awesomeness","score":9.99,"confirmed":true}
	// {
	//   "id": "123456789",
	//   "message": "Total awesomeness",
	//   "score": 9.99,
	//   "confirmed": true
	// }
}

// https://stackoverflow.com/questions/35691811/golang-unmarshal-json-map-array
// https://play.golang.org/p/UPoFxorqWl

func ExampleUnmarshal() {
	b := []byte(`[{"email":"example@test.com"}]`)
	c := []byte(`{"email":"example@test.com"}`)

	var m interface{}

	_ = json.Unmarshal(b, &m)

	switch v := m.(type) {
	case []interface{}:
		fmt.Println("this is b", v)
	default:
		fmt.Println("No type found")
	}

	// output:
	// this is b [map[email:example@test.com]]

	_ = json.Unmarshal(c, &m)

	switch v := m.(type) {
	case map[string]interface{}:
		fmt.Println("this is c", v)
	default:
		fmt.Println("No type found")
	}

	// Output:
	// this is b [map[email:example@test.com]]
	// this is c map[email:example@test.com]
}
