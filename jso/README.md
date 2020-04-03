```bash
➜  gogotcha git:(master) ✗ go test -bench=. ./jso/...                                                                                                                                                        [Fri Apr  3 15:59:20 2020]
?       github.com/bingoohuang/gogotcha/jso     [no test files]
goos: darwin
goarch: amd64
pkg: github.com/bingoohuang/gogotcha/jso/v1
BenchmarkAwesomeFromJSON-12               805695              1467 ns/op
BenchmarkAwesomeToJSON-12                2749590               439 ns/op
BenchmarkAwesomeToJSONPretty-12           775660              1391 ns/op
PASS
ok      github.com/bingoohuang/gogotcha/jso/v1  4.866s
goos: darwin
goarch: amd64
pkg: github.com/bingoohuang/gogotcha/jso/v2
BenchmarkAwesomeFromJSON-12              2773173               434 ns/op
BenchmarkAwesomeToJSON-12                2545711               473 ns/op
BenchmarkAwesomeToJSONPretty-12          1836224               655 ns/op
PASS
ok      github.com/bingoohuang/gogotcha/jso/v2  5.201s

```