package lang

import (
	"github.com/bingoohuang/strcase"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

func init() {
	extra.SetNamingStrategy(func(n string) string { return strcase.ToCamelLower(n) })
}

var Unmarshal = jsoniter.Unmarshal
var Marshal = jsoniter.Marshal
var MarshalIndent = jsoniter.MarshalIndent
