package lang

import (
	"github.com/bingoohuang/strcase"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

// nolint:gochecknoinits
func init() {
	extra.SetNamingStrategy(strcase.ToCamelLower)
}

// nolint:gochecknoglobals
var (
	Unmarshal     = jsoniter.Unmarshal
	Marshal       = jsoniter.Marshal
	MarshalIndent = jsoniter.MarshalIndent
)
