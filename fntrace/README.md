# fntrace

func tracing demo.

1. `go install -tags trace ./...`
1. `fntrace-demo`

```go
package main

import "github.com/bingoohuang/fntrace"

func main() {
    a()
}

func a() {
    defer fntrace.Trace()()
    b()
}

func b() {
    defer fntrace.Trace()()
    c()
}

func c() {
    defer fntrace.Trace()()
    d()
}

func d() {
    defer fntrace.Trace()()
}
```

```bash
g[01]:->main.a
g[01]:-->main.b
g[01]:--->main.c
g[01]:---->main.d
g[01]:<----main.d
g[01]:<---main.c
g[01]:<--main.b
g[01]:<-main.a
```

## thanks

1. [Go函数调用链跟踪的一种实现思路](https://tonybai.com/2020/12/10/a-kind-of-thinking-about-how-to-trace-function-call-chain/)
1. [bigwhite / functrace](https://github.com/bigwhite/functrace)
