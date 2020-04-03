package jso

import (
	"encoding/json"
	"fmt"

	v1 "github.com/bingoohuang/gogotcha/jso/v1"
	v2 "github.com/bingoohuang/gogotcha/jso/v2"
)

func v1demo() {
	awesome := v1.NewAwesome("123456789", "Total awesomeness", 9.99, true)
	awesomeJSON, _ := awesome.ToJSON(false)
	fmt.Printf("%s\n", awesomeJSON)
	moreAwesomeJSON, _ := awesome.ToJSON(true)
	fmt.Printf("%s\n", moreAwesomeJSON)
}

func v2demo() {
	awesome := v2.NewAwesome("123456789", "Total awesomeness", 9.99, true)
	awesomeJSON, _ := awesome.ToJSON(false)
	fmt.Printf("%s\n", awesomeJSON)
	moreAwesomeJSON, _ := awesome.ToJSON(true)
	fmt.Printf("%s\n", moreAwesomeJSON)
}

// https://stackoverflow.com/questions/35691811/golang-unmarshal-json-map-array
// https://play.golang.org/p/UPoFxorqWl

func parseAny() {
	b := []byte(`[{"email":"example@test.com"}]`)
	c := []byte(`{"email":"example@test.com"}`)

	var m interface{}

	json.Unmarshal(b, &m)
	switch v := m.(type) {
	case []interface{}:
		fmt.Println("this is b", v)
	default:
		fmt.Println("No type found")
	}

	// output:
	// this is b [map[email:example@test.com]]

	json.Unmarshal(c, &m)
	switch v := m.(type) {
	case map[string]interface{}:
		fmt.Println("this is c", v)
	default:
		fmt.Println("No type found")
	}

	// output:
	// this is c map[email:example@test.com]
}
