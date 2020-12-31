package lang_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/bingoohuang/gogotcha/lang"
	"github.com/stretchr/testify/assert"
)

// nolint:gochecknoglobals
var awesomeJSON = []byte(`{
  "id": "123456789",
  "message": "Total awesomeness",
  "score": 9.99,
  "confirmed": true
}`)

func TestAwesomeToJSON(t *testing.T) {
	awesome := Awesome{"123456789", "Total awesomeness", 9.99, true}

	testJSON, err := lang.MarshalIndent(awesome, "", "  ")

	assert.Nil(t, err)
	assert.Equal(t, testJSON, awesomeJSON)
}

func TestAwesomeFromJSON(t *testing.T) {
	var awesome Awesome

	assert.Nil(t, lang.Unmarshal(awesomeJSON, &awesome))
	assert.Equal(t, Awesome{"123456789", "Total awesomeness", 9.99, true}, awesome)
}

type Awesome struct {
	ID        string
	Message   string
	Score     float64
	Confirmed bool
}

func ExampleJSONMarshal() {
	data, _ := json.Marshal(context.WithValue(context.Background(), "a", "b"))
	fmt.Println(string(data))

	ti := struct {
		time.Time
		N int
	}{
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		5,
	}

	m, _ := json.Marshal(ti)
	fmt.Println(string(m))

	// Output:
	// {"Context":0}
	// "2020-12-20T00:00:00Z"
}

type T struct{ x int }
type T2 struct{ x int }
type T3 struct{ x int }
type T4 struct{ x int }
type T5 struct{ x int }

func (t T2) String() string              { return "boo" }
func (t T3) Format(s fmt.State, _ rune)  { fmt.Fprint(s, "baa") }
func (t T4) String() string              { return "boo" }
func (t T4) Format(s fmt.State, _ rune)  { fmt.Fprint(s, "baa") }
func (t T5) String() string              { return "boo" }
func (t *T5) Format(s fmt.State, _ rune) { fmt.Fprint(s, "baa") }

func ExampleFormat() {
	v := T{123}
	fmt.Println(v)

	v2 := T2{123}
	fmt.Println(v2)

	v3 := T3{123}
	fmt.Println(v3)

	v4 := T4{123}
	fmt.Println(v4)

	v5 := T5{123}
	fmt.Println(v5)

	// Output:
	// {123}
	// boo
	// baa
	// baa
	// boo
}

type Ty struct{ x int }

type Ti interface {
	Say() string
}

func (t Ty) Say() string { return "Hi" }

func ExampleSay() {
	var a Ti = Ty{123}
	a.Say()
	// Output:
}
