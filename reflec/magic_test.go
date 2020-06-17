package reflec_test

import (
	"fmt"
	"github.com/bingoohuang/gogotcha/reflec"
	"reflect"
	"testing"
	"unsafe"
)

func BenchmarkNew(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reflec.New()
	}
}

func BenchmarkNewUseReflect(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reflec.NewUseReflect()
	}
}

func BenchmarkNewV2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reflec.NewV2()
	}
}

func BenchmarkNewUseReflectV2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reflec.NewUseReflectV2()
	}
}

func BenchmarkNewQuickReflectV2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reflec.NewQuickReflectV2()
	}
}

func BenchmarkNewUseReflectV2WithPool(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		obj := reflec.NewQuickReflectWithPool()
		reflec.Pool.Put(obj)
	}
}

func BenchmarkNewV3(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reflec.NewV3()
	}
}

func BenchmarkNewUseReflectV3(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reflec.NewUseReflectV3()
	}
}

type emptyInterface struct {
	typ  *struct{}
	word unsafe.Pointer
}

func ExampleStruct() {
	var in interface{}
	in = reflec.PeopleV2{
		Age:   30,
		Name:  "patrickchen",
		Test1: "test1",
		Test2: "test2",
	}

	t2 := uintptr(((*emptyInterface)(unsafe.Pointer(&in))).word)
	*(*int)(unsafe.Pointer(t2)) = 111
	fmt.Println(in)
	// Output: {111 patrickchen test1 test2}
}

func ExampleStructV2() {
	var in interface{}
	in = reflec.PeopleV2{
		Age:   30,
		Name:  "patrickchen",
		Test1: "test1",
		Test2: "test2",
	}

	typeP := reflect.TypeOf(in)
	offset1 := typeP.Field(1).Offset // get the offset by reflection
	offset2 := typeP.Field(2).Offset
	offset3 := typeP.Field(3).Offset

	t2 := uintptr(((*emptyInterface)(unsafe.Pointer(&in))).word)

	*(*int)(unsafe.Pointer(t2)) = 111 //get the first member variable address
	*(*string)(unsafe.Pointer(t2 + offset1)) = "hello"
	*(*string)(unsafe.Pointer(t2 + offset2)) = "hello1"
	*(*string)(unsafe.Pointer(t2 + offset3)) = "hello2"
	fmt.Println(in)
	//Output: {111 hello hello1 hello2}
}

type Test1 struct {
	a int32
	b []byte
}
type Test2 struct {
	b int16
	a string
}

func ExampleStruct2() {
	t1 := Test1{
		a: 1,
		b: []byte("asdasd"),
	}

	t2 := *(*Test2)(unsafe.Pointer(&t1))
	fmt.Println(t2)
	//Output: {1 asdasd}
}
