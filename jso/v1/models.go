package v1

import (
	"encoding/json"
)

type Awesome struct {
	ID        string  `json:"id"`
	Message   string  `json:"message"`
	Score     float64 `json:"score"`
	Confirmed bool    `json:"confirmed"`
}

// NewAwesome creates a new Instance of Awesome.
func NewAwesome(id string, message string, score float64, confirmed bool) *Awesome {
	return &Awesome{
		id,
		message,
		score,
		confirmed,
	}
}

// NewAwesomeFromJSON creates a new Instance of Awesome from JSON.
func NewAwesomeFromJSON(jsonData []byte) (awesome *Awesome) {
	if err := json.Unmarshal(jsonData, &awesome); err != nil {
		return nil
	}

	return awesome
}

// ToJSON marshalls Awesome to JSON.
func (a *Awesome) ToJSON(pretty bool) ([]byte, error) {
	if pretty {
		return json.MarshalIndent(a, "", "  ")
	}

	return json.Marshal(a)
}
