package jsonflatten_test

import (
	"testing"

	"github.com/bingoohuang/gogotcha/jsonflatten"
	"github.com/stretchr/testify/assert"
)

func TestFlatten(t *testing.T) {
	o, err := jsonflatten.FlattenJSON(`{"name":"bingoohuang", "address":{"city":"San Francisco", "postcode":123}}`)
	assert.Nil(t, err)

	assert.Equal(t, map[string]interface{}{
		"name":     "bingoohuang",
		"city":     "San Francisco",
		"postcode": float64(123),
	}, o)

	// nolint:lll
	o, err = jsonflatten.FlattenJSON(`{"name":"bingoohuang", "address": {"detail":{"city":"San Francisco", "postcode":123}}}`)
	assert.Nil(t, err)

	assert.Equal(t, map[string]interface{}{
		"name":     "bingoohuang",
		"city":     "San Francisco",
		"postcode": float64(123),
	}, o)

	o, err = jsonflatten.FlattenJSON(`{"name":"bingoohuang", "address":{"city":"San Francisco", "postcode":123}}`,
		jsonflatten.AllowFn(func(keys []string) bool {
			if len(keys) == 2 && keys[0] == "address" {
				return keys[1] == "city"
			}

			return true
		}))
	assert.Nil(t, err)

	assert.Equal(t, map[string]interface{}{
		"name": "bingoohuang",
		"city": "San Francisco",
	}, o)
}
