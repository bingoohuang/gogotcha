package main

import (
	"fmt"

	"github.com/goccy/go-yaml"

	"github.com/elastic/go-ucfg"
	yu "github.com/elastic/go-ucfg/yaml"
)

func main() {
	demo1()
	demo2()
	demo3()
	demo4()
	demo5()
}

func demo1() {
	var v struct {
		A int
		B string
	}

	v.A = 1
	v.B = "hello"
	bytes, _ := yaml.Marshal(v)

	fmt.Println(string(bytes)) // "a: 1\nb: hello\n"
}

func demo2() {
	yml := `
%YAML 1.2
---
a: 1
b: c
`

	var v struct {
		A int
		B string
	}

	if err := yaml.Unmarshal([]byte(yml), &v); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", v) // {A:1 B:c}
}

func demo3() {
	yml := `---
foo: 1
bar: c
`
	// To control marshal/unmarshal behavior, you can use the yaml tag
	var v struct {
		A int    `yaml:"foo"`
		B string `yaml:"bar"`
	}

	if err := yaml.Unmarshal([]byte(yml), &v); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", v) // {A:1 B:c}
}

func demo4() {
	yml := `---
foo: 1
bar: c
`
	// For convenience, we also accept the json tag. Note that not all options from the json tag
	// will have significance when parsing YAML documents. If both tags exist, yaml tag will take precedence.

	var v struct {
		A int    `json:"foo"`
		B string `json:"bar"`
	}

	if err := yaml.Unmarshal([]byte(yml), &v); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", v) // {A:1 B:c}
}

// HealthHTTPConfig defines config structure for HTTP health check.
type HealthHTTPConfig struct {
	Hosts []string              `config:"hosts"`
	Check HealthHTTPConfigCheck `config:"check"`
}

// HealthHTTPConfigCheck defines config check sub structure for HTTP health check.
type HealthHTTPConfigCheck struct {
	Response HealthHTTPConfigCheckResponse `config:"response"`
}

// HealthHTTPConfigCheckResponse defines config check sub structure for HTTP health check.
type HealthHTTPConfigCheckResponse struct {
	Status int `config:"status"`
}

func parseHealthHTTPConfig(yml string) (v HealthHTTPConfig, err error) {
	yc, err := yu.NewConfig([]byte(yml), ucfg.PathSep("."))
	if err != nil {
		return v, err
	}

	if err := yc.Unpack(&v); err != nil {
		return v, err
	}

	return v, err
}

func demo5() {
	yml := `
hosts: ["http://localhost:80/service/status"]
check.response.status: 200
`
	v, err := parseHealthHTTPConfig(yml)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", v) // {Hosts:[http://localhost:80/service/status] Check:{Response:{Status:200}}}
}
