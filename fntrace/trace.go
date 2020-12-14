// +build trace

package fntrace

import (
	"fmt"
	"github.com/bingoohuang/fntrace/pkg/gid"
	"runtime"
	"strings"
	"sync"
)

var m = MakeGidIndent()

type GidIndent struct {
	cache map[uint64]int
	mu    sync.Mutex
}

func MakeGidIndent() *GidIndent {
	return &GidIndent{
		cache: make(map[uint64]int),
	}
}

func (g *GidIndent) GetAdd(id uint64, delta int) int {
	g.mu.Lock()
	v := g.cache[id]
	g.cache[id] = v + delta
	g.mu.Unlock()

	return v
}

type trace struct {
	id     uint64
	name   string
	indent int
}

func createTrace(id uint64) trace {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		panic("not found caller")
	}

	t := trace{
		id:     id,
		name:   runtime.FuncForPC(pc).Name(),
		indent: m.GetAdd(id, 1) + 1,
	}

	fmt.Printf("g[%02d]:%s>%s\n", t.id, strings.Repeat("-", t.indent), t.name)
	return t
}

func (t trace) printEnd() {
	m.GetAdd(t.id, -1)
	fmt.Printf("g[%02d]:<%s%s\n", t.id, strings.Repeat("-", t.indent), t.name)

}

func Trace() func() {
	return createTrace(gid.GetGID()).printEnd
}
