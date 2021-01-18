package custom_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/bingoohuang/gogotcha/jso/custom"
	"github.com/bingoohuang/gogotcha/lang"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Document struct {
	Name  string    `json:"name"`
	Stamp time.Time `json:"stamp"`
}

func (d Document) MarshalJSON() ([]byte, error) {
	// To alias the original type.
	// This alias will have all the same fields, but none of the methods (MarshalJSON/UnmarshalJSON).
	type Alias Document
	return json.Marshal(&struct {
		*Alias
		Stamp string `json:"stamp"`
	}{
		Alias: (*Alias)(&d),
		Stamp: d.Stamp.Format("2006-01-02 15:04:05.000"),
	})
}

func (d *Document) UnmarshalJSON(data []byte) error {
	// To alias the original type.
	// This alias will have all the same fields, but none of the methods (MarshalJSON/UnmarshalJSON).
	type Alias Document
	aux := &struct {
		*Alias
		Stamp string `json:"stamp"`
	}{
		Alias: (*Alias)(d),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	d.Stamp, _ = time.ParseInLocation("2006-01-02 15:04:05.000", aux.Stamp, time.Local)
	return nil
}

func (p Person) MarshalJSON() ([]byte, error) {
	// To alias the original type.
	// This alias will have all the same fields, but none of the methods (MarshalJSON/UnmarshalJSON).
	type Alias Person

	v := struct {
		Name string `json:"name"`
		Alias
	}{
		Alias: Alias(p),
		Name:  "Fixed:" + p.Name,
	}

	printStructMeta(p)
	printStructMeta(&p)
	printStructMeta(v.Alias)
	printStructMeta(&v.Alias)

	// Person Field:(0)(Name)
	// Person Field:(1)(Age)
	// Person Method:(0)(MarshalJSON)
	// *Person Field:(0)(Name)
	// *Person Field:(1)(Age)
	// *Person Method:(0)(MarshalJSON)
	// *Person Method:(1)(UnmarshalJSON)
	// Alias Field:(0)(Name)
	// Alias Field:(1)(Age)
	// *Alias Field:(0)(Name)
	// *Alias Field:(1)(Age)

	return json.Marshal(v)
}

func (p *Person) UnmarshalJSON(data []byte) error {
	// To alias the original type.
	// This alias will have all the same fields, but none of the methods (MarshalJSON/UnmarshalJSON).
	type Alias Person
	aux := &struct {
		Name string `json:"name"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	p.Name = strings.TrimPrefix(aux.Name, "Fixed:")
	return nil
}

func TestCustom2(t *testing.T) {
	// http://choly.ca/post/go-json-marshalling/
	p := Person{Name: "bingoohuang", Age: 100}
	v, err := json.Marshal(p)
	assert.Nil(t, err)
	assert.Equal(t, `{"name":"Fixed:bingoohuang","age":100}`, string(v))

	var p2 Person
	assert.Nil(t, json.Unmarshal(v, &p2))
	assert.Equal(t, p, p2)

	ti := time.Now().Truncate(time.Millisecond)
	d := Document{Name: "bingoohuang", Stamp: ti}
	v2, err2 := json.Marshal(d)
	assert.Nil(t, err2)
	assert.Equal(t, `{"name":"bingoohuang","stamp":"`+ti.Format("2006-01-02 15:04:05.000")+`"}`, string(v2))

	var d2 Document
	assert.Nil(t, json.Unmarshal(v2, &d2))
	assert.Equal(t, d, d2)
}

func printStructMeta(i interface{}) {
	v, ok := i.(reflect.Value)
	if !ok {
		v = reflect.ValueOf(i)
	}

	typ := v.Type().Name()
	fv := v

	if v.Kind() == reflect.Ptr {
		fv = v.Elem()
		typ = "*" + fv.Type().Name()
	}

	if fv.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < fv.NumField(); i++ {
		fmt.Printf("%s Field:(%d)(%s)\n", typ, i, fv.Type().Field(i).Name)
	}

	for i := 0; i < v.NumMethod(); i++ {
		fmt.Printf("%s Method:(%d)(%s)\n", typ, i, v.Type().Method(i).Name)
	}
}

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
