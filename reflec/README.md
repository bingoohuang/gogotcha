# Go reflect works like a magic

[Go reflect works like a magic](https://medium.com/swlh/go-reflect-works-like-a-magic-575cb0cbecc2)

```bash
➜  reflec git:(master) ✗ go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out                                                                                                       [Wed Jun 17 09:50:10 2020]
goos: darwin
goarch: amd64
pkg: github.com/bingoohuang/gogotcha/reflec
BenchmarkNew-12                         1000000000               0.521 ns/op           0 B/op          0 allocs/op
BenchmarkNewUseReflect-12                7900816               138 ns/op              64 B/op          2 allocs/op
BenchmarkNewV2-12                       1000000000               0.256 ns/op           0 B/op          0 allocs/op
BenchmarkNewUseReflectV2-12              5872969               202 ns/op             128 B/op          2 allocs/op
BenchmarkNewQuickReflectV2-12           20279065                59.3 ns/op            64 B/op          1 allocs/op
BenchmarkNewUseReflectV2WithPool-12     67272578                17.4 ns/op             0 B/op          0 allocs/op
BenchmarkNewV3-12                       1000000000               0.511 ns/op           0 B/op          0 allocs/op
BenchmarkNewUseReflectV3-12             62567138                20.0 ns/op             0 B/op          0 allocs/op
➜  reflec git:(master) ✗ go tool pprof ./profile.out                                                                                                                                                         [Wed Jun 17 09:50:50 2020]
Type: cpu
Time: Jun 17, 2020 at 9:50am (CST)
Duration: 7.93s, Total samples = 7.78s (98.10%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) list NewUseReflectV2
Total: 7.78s
ROUTINE ======================== github.com/bingoohuang/gogotcha/reflec.NewUseReflectV2 in /Users/bingoobjca/github/gogotcha/reflec/a.go
      40ms      780ms (flat, cum) 10.03% of Total
         .          .     63:   }
         .          .     64:}
         .          .     65:
         .          .     66:func NewUseReflectV2() interface{} {
         .          .     67:   var p PeopleV2
      20ms      180ms     68:   t := reflect.TypeOf(p)
         .      250ms     69:   v := reflect.New(t)
         .       60ms     70:   f0 := v.Elem().Field(0)
         .       60ms     71:   f0.Set(reflect.ValueOf(30))
      10ms       30ms     72:   f1 := v.Elem().Field(1)
         .       40ms     73:   f1.Set(reflect.ValueOf("patrickchen"))
         .       10ms     74:   f2 := v.Elem().Field(2)
         .       20ms     75:   f2.Set(reflect.ValueOf("test1"))
      10ms       10ms     76:   f3 := v.Elem().Field(3)
         .      100ms     77:   f3.Set(reflect.ValueOf("test2"))
         .       20ms     78:   return v.Interface()
         .          .     79:}
         .          .     80:
         .          .     81:func NewQuickReflectV2() interface{} {
         .          .     82:   v := reflect.New(t)
         .          .     83:
ROUTINE ======================== github.com/bingoohuang/gogotcha/reflec_test.BenchmarkNewUseReflectV2 in /Users/bingoobjca/github/gogotcha/reflec/a_test.go
         0      780ms (flat, cum) 10.03% of Total
         .          .     34:
         .          .     35:func BenchmarkNewUseReflectV2(b *testing.B) {
         .          .     36:   b.ReportAllocs()
         .          .     37:   b.ResetTimer()
         .          .     38:   for i := 0; i < b.N; i++ {
         .      780ms     39:           reflec.NewUseReflectV2()
         .          .     40:   }
         .          .     41:}
         .          .     42:
         .          .     43:func BenchmarkNewQuickReflectV2(b *testing.B) {
         .          .     44:   b.ReportAllocs()
ROUTINE ======================== github.com/bingoohuang/gogotcha/reflec_test.BenchmarkNewUseReflectV2WithPool in /Users/bingoobjca/github/gogotcha/reflec/a_test.go
      30ms      1.04s (flat, cum) 13.37% of Total
         .          .     49:}
         .          .     50:
         .          .     51:func BenchmarkNewUseReflectV2WithPool(b *testing.B) {
         .          .     52:   b.ReportAllocs()
         .          .     53:   b.ResetTimer()
      10ms       10ms     54:   for i := 0; i < b.N; i++ {
      10ms      580ms     55:           obj := reflec.NewQuickReflectWithPool()
      10ms      450ms     56:           reflec.Pool.Put(obj)
         .          .     57:   }
         .          .     58:}
         .          .     59:
         .          .     60:func BenchmarkNewV3(b *testing.B) {
         .          .     61:   b.ReportAllocs()
(pprof) list NewQuickReflectV2
Total: 7.78s
ROUTINE ======================== github.com/bingoohuang/gogotcha/reflec.NewQuickReflectV2 in /Users/bingoobjca/github/gogotcha/reflec/a.go
      90ms      700ms (flat, cum)  9.00% of Total
         .          .     76:   f3 := v.Elem().Field(3)
         .          .     77:   f3.Set(reflect.ValueOf("test2"))
         .          .     78:   return v.Interface()
         .          .     79:}
         .          .     80:
      20ms       20ms     81:func NewQuickReflectV2() interface{} {
      20ms      580ms     82:   v := reflect.New(t)
         .          .     83:
      20ms       70ms     84:   p := v.Interface()
      10ms       10ms     85:   ptr0 := uintptr((*emptyInterface)(unsafe.Pointer(&p)).word)
      10ms       10ms     86:   ptr1 := ptr0 + offset1
         .          .     87:   ptr2 := ptr0 + offset2
         .          .     88:   ptr3 := ptr0 + offset3
         .          .     89:   *((*int)(unsafe.Pointer(ptr0))) = 30
      10ms       10ms     90:   *((*string)(unsafe.Pointer(ptr1))) = "patrickchen"
         .          .     91:   *((*string)(unsafe.Pointer(ptr2))) = "test1"
         .          .     92:   *((*string)(unsafe.Pointer(ptr3))) = "test2"
         .          .     93:   return p
         .          .     94:}
         .          .     95:
ROUTINE ======================== github.com/bingoohuang/gogotcha/reflec_test.BenchmarkNewQuickReflectV2 in /Users/bingoobjca/github/gogotcha/reflec/a_test.go
         0      700ms (flat, cum)  9.00% of Total
         .          .     42:
         .          .     43:func BenchmarkNewQuickReflectV2(b *testing.B) {
         .          .     44:   b.ReportAllocs()
         .          .     45:   b.ResetTimer()
         .          .     46:   for i := 0; i < b.N; i++ {
         .      700ms     47:           reflec.NewQuickReflectV2()
         .          .     48:   }
         .          .     49:}
         .          .     50:
         .          .     51:func BenchmarkNewUseReflectV2WithPool(b *testing.B) {
         .          .     52:   b.ReportAllocs()
(pprof) list NewQuickReflectWithPool
Total: 7.78s
ROUTINE ======================== github.com/bingoohuang/gogotcha/reflec.NewQuickReflectWithPool in /Users/bingoobjca/github/gogotcha/reflec/a.go
     220ms      570ms (flat, cum)  7.33% of Total
         .          .    105:   }
         .          .    106:   for i := 0; i < 100; i++ {
         .          .    107:           Pool.Put(reflect.New(t).Elem())
         .          .    108:   }
         .          .    109:}
      20ms       20ms    110:func NewQuickReflectWithPool() interface{} {
      50ms      400ms    111:   p := Pool.Get()
         .          .    112:   ptr0 := uintptr((*emptyInterface)(unsafe.Pointer(&p)).word)
      20ms       20ms    113:   ptr1 := ptr0 + offset1
      30ms       30ms    114:   ptr2 := ptr0 + offset2
      10ms       10ms    115:   ptr3 := ptr0 + offset3
         .          .    116:   *((*int)(unsafe.Pointer(ptr0))) = 18
      20ms       20ms    117:   *((*string)(unsafe.Pointer(ptr1))) = "shiina"
      40ms       40ms    118:   *((*string)(unsafe.Pointer(ptr2))) = "test1"
      30ms       30ms    119:   *((*string)(unsafe.Pointer(ptr3))) = "test2"
         .          .    120:   return p
         .          .    121:}
         .          .    122:
         .          .    123:type PeopleV3 struct {
         .          .    124:}
```