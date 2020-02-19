package main

import (
	"encoding/json"
	"fmt"
)

// https://stackoverflow.com/questions/35691811/golang-unmarshal-json-map-array
// https://play.golang.org/p/UPoFxorqWl

func main() {
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
