package v2

import (
	"github.com/bingoohuang/gogotcha/lang"
)

// NewAwesome creates a new Instance of Awesome
func NewAwesome(id string, message string, score float64, confirmed bool) *Awesome {
	return &Awesome{
		id,
		message,
		score,
		confirmed,
	}
}

type Awesome struct {
	ID        string
	Message   string
	Score     float64
	Confirmed bool
}

// NewAwesomeFromJSON creates a new Instance of Awesome from JSON
func NewAwesomeFromJSON(jsonData []byte) (awesome *Awesome) {
	if err := lang.Unmarshal(jsonData, &awesome); err != nil {
		return nil
	}

	return awesome
}

// ToJSON marshalls Awesome to JSON
func (a *Awesome) ToJSON(pretty bool) ([]byte, error) {
	if pretty {
		return lang.MarshalIndent(a, "", "  ")
	}

	return lang.Marshal(a)
}
