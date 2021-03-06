package v1_test

import (
	"testing"

	v1 "github.com/bingoohuang/gogotcha/jso/v1"
)

// nolint:gochecknoglobals
var awesomeJSON = []byte(`{
  "id": "123456789",
  "message": "Total awesomeness",
  "score": 9.99,
  "confirmed": true
}`)

func TestAwesomeToJSON(t *testing.T) {
	awesome := v1.NewAwesome("123456789", "Total awesomeness", 9.99, true)

	testJSON, err := awesome.ToJSON(true)

	if err != nil {
		t.Error("Failed to create json from awesome")
	}

	if string(testJSON) != string(awesomeJSON) {
		t.Errorf("JSON output\n%s\nis not as expected\n%s", testJSON, awesomeJSON)
	}
}

func TestAwesomeFromJSON(t *testing.T) {
	awesome := v1.NewAwesomeFromJSON(awesomeJSON)

	if awesome == nil {
		t.Error("Unmarshalling json into awesome failed")
		return
	}

	if awesome.ID != "123456789" {
		t.Error("Awesome ID does not match expected value")
	}

	if awesome.Message != "Total awesomeness" {
		t.Error("Awesome Message does not match expected value")
	}

	if awesome.Score != 9.99 {
		t.Error("Awesome Id does not match expected value")
	}

	if !awesome.Confirmed {
		t.Error("Awesome Confirmed does not match expected value")
	}
}

func BenchmarkAwesomeFromJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v1.NewAwesomeFromJSON(awesomeJSON)
	}
}

func BenchmarkAwesomeToJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		awesome := v1.NewAwesome("123456789", "Total awesomeness", 9.99, true)
		_, _ = awesome.ToJSON(false)
	}
}

func BenchmarkAwesomeToJSONPretty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		awesome := v1.NewAwesome("123456789", "Total awesomeness", 9.99, true)
		_, _ = awesome.ToJSON(true)
	}
}
