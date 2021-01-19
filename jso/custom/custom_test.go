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
	Name string `json:"name,abc"`
	Age  int    `json:"age"`
}

type Document struct {
	Name  string    `json:"name"`
	Stamp time.Time `json:"stamp,yyyy-MM-dd HH:mm:ss.SSS"`
}

func (d Document) MarshalJSON() ([]byte, error) {
	// To alias the original type.
	// This alias will have all the same fields, but none of the methods (MarshalJSON/UnmarshalJSON).
	type Alias Document

	alias := struct {
		Alias
		Stamp string `json:"stamp"`
	}{
		Alias: (Alias)(d),
		Stamp: d.Stamp.Format("2006-01-02 15:04:05.000"),
	}

	printStructMeta(d)
	printStructMeta(&d)
	printStructMeta(alias)
	printStructMeta(&alias)

	return json.Marshal(alias)
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
		Alias
		Name string `json:"name"`
	}{
		Alias: Alias(p),
		Name:  "Fixed:" + p.Name,
	}

	printStructMeta(p)
	printStructMeta(&p)
	printStructMeta(v)
	printStructMeta(&v)

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
		*Alias
		Name string `json:"name"`
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
	assert.Equal(t, `{"age":100,"name":"Fixed:bingoohuang"}`, string(v))

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

	typ := v.Type().String()
	fv := v

	if v.Kind() == reflect.Ptr {
		fv = v.Elem()
	}

	if fv.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < fv.NumField(); i++ {
		f := fv.Type().Field(i)
		jsonTag := f.Tag.Get("json")
		anonymous := ""
		if f.Anonymous {
			anonymous = "Anonymous "
		}
		fmt.Printf("%s %sField:(%d)(%s),jsonTag:(%s)\n", typ, anonymous, i, f.Name, jsonTag)
	}

	for i := 0; i < v.NumMethod(); i++ {
		m := v.Type().Method(i)
		fmt.Printf("%s Method:(%d)(%s)\n", typ, i, m.Name)
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
