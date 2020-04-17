package jsonflatten

import "encoding/json"

type AllowFn func(keys []string) bool

type Filter interface {
	Allow(keys []string) bool
}

func (a AllowFn) Allow(keys []string) bool {
	return a(keys)
}

// FlattenJSON takes a JSON s and returns a new one where nested maps/slice[0] flatten.
func FlattenJSON(s string, filters ...Filter) (map[string]interface{}, error) {
	var m map[string]interface{}

	if err := json.Unmarshal([]byte(s), &m); err != nil {
		return nil, err
	}

	return flatten([]string{}, m, make(map[string]interface{}), filters), nil
}

// Flatten takes a map and returns a new one where nested maps/slice[0] flatten.
func Flatten(m map[string]interface{}, filters ...Filter) map[string]interface{} {
	return flatten([]string{}, m, make(map[string]interface{}), filters)
}

func flatten(parentKeys []string, i, o map[string]interface{}, filters []Filter) map[string]interface{} {
	for k, v := range i {
		ps := append(parentKeys, k)
		if !allowKeys(ps, filters) {
			continue
		}

		switch child := v.(type) {
		case map[string]interface{}:
			flatten(ps, child, o, filters)
		case []interface{}:
			if len(child) > 0 {
				flatten(ps, child[0].(map[string]interface{}), o, filters)
			}
		default:
			o[k] = v
		}
	}

	return o
}

func allowKeys(keys []string, filters []Filter) bool {
	if len(filters) == 0 {
		return true
	}

	for _, filter := range filters {
		if filter.Allow(keys) {
			return true
		}
	}

	return false
}
