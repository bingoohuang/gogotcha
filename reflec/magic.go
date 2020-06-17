package reflec

import (
	"reflect"
	"sync"
	"unsafe"
)

type People struct {
	Age  int
	Name string
}

func New() *People {
	return &People{
		Age:  30,
		Name: "patrickchen",
	}
}

func NewUseReflect() interface{} {
	var p People
	t := reflect.TypeOf(p)
	v := reflect.New(t)
	f0 := v.Elem().Field(0)
	f0.Set(reflect.ValueOf(30))
	f1 := v.Elem().Field(1)
	f1.Set(reflect.ValueOf("patrickchen"))
	return v.Interface()
}

var (
	offset1 uintptr
	offset2 uintptr
	offset3 uintptr
	t       = reflect.TypeOf((*PeopleV2)(nil)).Elem()
)

func init() {
	offset1 = t.Field(1).Offset
	offset2 = t.Field(2).Offset
	offset3 = t.Field(3).Offset
}

type PeopleV2 struct {
	Age   int
	Name  string
	Test1 string
	Test2 string
}

type emptyInterface struct {
	typ  *struct{}
	word unsafe.Pointer
}

func NewV2() interface{} {
	return &PeopleV2{
		Age:   18,
		Name:  "shiina",
		Test1: "test1",
		Test2: "test2",
	}
}

func NewUseReflectV2() interface{} {
	var p PeopleV2
	t := reflect.TypeOf(p)
	v := reflect.New(t)
	f0 := v.Elem().Field(0)
	f0.Set(reflect.ValueOf(30))
	f1 := v.Elem().Field(1)
	f1.Set(reflect.ValueOf("patrickchen"))
	f2 := v.Elem().Field(2)
	f2.Set(reflect.ValueOf("test1"))
	f3 := v.Elem().Field(3)
	f3.Set(reflect.ValueOf("test2"))
	return v.Interface()
}

func NewQuickReflectV2() interface{} {
	v := reflect.New(t)

	p := v.Interface()
	ptr0 := uintptr((*emptyInterface)(unsafe.Pointer(&p)).word)
	ptr1 := ptr0 + offset1
	ptr2 := ptr0 + offset2
	ptr3 := ptr0 + offset3
	*((*int)(unsafe.Pointer(ptr0))) = 30
	*((*string)(unsafe.Pointer(ptr1))) = "patrickchen"
	*((*string)(unsafe.Pointer(ptr2))) = "test1"
	*((*string)(unsafe.Pointer(ptr3))) = "test2"
	return p
}

var (
	/**
	  ...........
	  **/
	Pool sync.Pool
)

func init() {
	Pool.New = func() interface{} {
		return reflect.New(t)
	}
	for i := 0; i < 100; i++ {
		Pool.Put(reflect.New(t).Elem())
	}
}
func NewQuickReflectWithPool() interface{} {
	p := Pool.Get()
	ptr0 := uintptr((*emptyInterface)(unsafe.Pointer(&p)).word)
	ptr1 := ptr0 + offset1
	ptr2 := ptr0 + offset2
	ptr3 := ptr0 + offset3
	*((*int)(unsafe.Pointer(ptr0))) = 18
	*((*string)(unsafe.Pointer(ptr1))) = "shiina"
	*((*string)(unsafe.Pointer(ptr2))) = "test1"
	*((*string)(unsafe.Pointer(ptr3))) = "test2"
	return p
}

type PeopleV3 struct {
}

func NewV3() interface{} {
	return &PeopleV3{}
}

func NewUseReflectV3() interface{} {
	var p PeopleV3
	t := reflect.TypeOf(p)
	v := reflect.New(t)
	return v.Interface()
}
