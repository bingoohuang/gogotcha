# gostruct

A Runtime Struct Builder for Go https://github.com/itsubaki/gostruct

## Example

```go
func Example() {
	person := gostruct.New().
		AddString("Name").
		AddInt64("Age").
		Build()

	p := person.New()
	p.SetString("Name", "gopher")
	p.SetInt64("Age", 11)

	fmt.Printf(" %T:  %+v\n", p.Interface(), p.Interface())
	fmt.Printf("%T: %+v\n", p.Addr(), p.Addr())
	jsonBytes, _ := json.Marshal(p.Interface())
	fmt.Printf("JSON: %s\n", jsonBytes)

	// Output:
	//  struct { Name string; Age int64 }:  {Name:gopher Age:11}
	// *struct { Name string; Age int64 }: &{Name:gopher Age:11}
	// JSON: {"Name":"gopher","Age":11}
}
```
