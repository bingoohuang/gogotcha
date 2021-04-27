package reflec_test

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

func ReflectField(f interface{}, fieldNames string) reflect.Value {
	v := reflect.ValueOf(f)
	for _, name := range strings.Split(fieldNames, ".") {
		if v.IsValid() && v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		if !v.IsValid() {
			return v
		}
		v = v.FieldByName(name)
	}

	return v
}

func fileFD(f *os.File) int {
	file := reflect.ValueOf(f).Elem().FieldByName("file").Elem()
	pfdVal := file.FieldByName("pfd")
	return int(pfdVal.FieldByName("Sysfd").Int())
}

func TestFd(t *testing.T) {
	temp, _ := os.CreateTemp("", "")
	defer os.Remove(temp.Name())

	f, _ := os.Open(temp.Name())
	defer f.Close()

	fmt.Printf("file descriptor is %d,%d,%d\n", f.Fd(), fileFD(f),
		ReflectField(f, "file.pfd.Sysfd").Int())
}
