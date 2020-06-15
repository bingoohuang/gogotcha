// Package internstrings interns strings.
// Interning is best effort only.
// Interned strings may be removed automatically
// at any time without notification.
// All functions may be called concurrently
// with themselves and each other.
package internstrings

import "sync"

// nolint:gochecknoglobals
var (
	pool = sync.Pool{
		New: func() interface{} {
			return make(map[string]string)
		},
	}
)

// String returns s, interned.
func String(s string) string {
	m := pool.Get().(map[string]string)
	if c, ok := m[s]; ok {
		pool.Put(m)

		return c
	}

	m[s] = s
	pool.Put(m)

	return s
}

// Bytes returns b converted to a string, interned.
func Bytes(b []byte) string {
	m := pool.Get().(map[string]string)
	if c, ok := m[string(b)]; ok {
		pool.Put(m)
		return c
	}

	s := string(b)
	m[s] = s
	pool.Put(m)

	return s
}
