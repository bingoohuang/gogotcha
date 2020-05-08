package lang

import (
	"github.com/bingoohuang/strcase"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

// nolint gochecknoinits
func init() {
	extra.SetNamingStrategy(func(n string) string { return strcase.ToCamelLower(n) })
}

// nolint gochecknoinits
var (
	Unmarshal     = jsoniter.Unmarshal
	Marshal       = jsoniter.Marshal
	MarshalIndent = jsoniter.MarshalIndent
)
