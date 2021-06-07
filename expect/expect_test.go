package expect

import (
	"testing"
)

func TestGetTSimple(t *testing.T) {
	if tt := getT(); t != tt {
		t.Fatalf("expected: %p, got %p", t, tt)
	}
}

func indirectGetT() *testing.T { return getT() }

func TestGetTIndirect(t *testing.T) {
	if tt := indirectGetT(); t != tt {
		t.Fatalf("expected: %p, got %p", t, tt)
	}
}

func getTWithParam(t *testing.T) *testing.T { return getT() }

func TestGetTWithTestingArgument(t *testing.T) {
	if tt := getTWithParam(t); t != tt {
		t.Fatalf("expected: %p, got %p", t, tt)
	}
}
