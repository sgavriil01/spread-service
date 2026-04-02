package spread

import (
	"testing"
	"time"
)

func resetSpreads() {
	spreads = map[string]Spread{}
}

func TestIsValidSymbol(t *testing.T) {
	if !IsValidSymbol("EURUSD") {
		t.Fatalf("expected EURUSD to be valid")
	}

	if IsValidSymbol("INVALID") {
		t.Fatalf("expected INVALID to be invalid")
	}
}

func TestSetAndGetSpread(t *testing.T) {
	resetSpreads()

	saved, ok := SetSpread("EURUSD", 0.0001)
	if !ok {
		t.Fatalf("expected SetSpread to succeed")
	}

	if saved.Symbol != "EURUSD" {
		t.Fatalf("unexpected symbol: %s", saved.Symbol)
	}

	if saved.Spread != 0.0001 {
		t.Fatalf("unexpected spread: %f", saved.Spread)
	}

	if saved.UpdatedAt.IsZero() {
		t.Fatalf("expected UpdatedAt to be set")
	}

	if saved.UpdatedAt.Location() != time.UTC {
		t.Fatalf("expected UpdatedAt to be in UTC")
	}

	got, found := GetSpread("EURUSD")
	if !found {
		t.Fatalf("expected GetSpread to find value")
	}

	if got != saved {
		t.Fatalf("expected stored spread to match saved spread")
	}
}

func TestSetSpreadInvalidSymbol(t *testing.T) {
	resetSpreads()

	_, ok := SetSpread("INVALID", 0.0001)
	if ok {
		t.Fatalf("expected SetSpread to fail for invalid symbol")
	}
}

func TestGetSpreadNotFoundCases(t *testing.T) {
	resetSpreads()

	if _, found := GetSpread("EURUSD"); found {
		t.Fatalf("expected not found for never-set valid symbol")
	}

	if _, found := GetSpread("INVALID"); found {
		t.Fatalf("expected not found for invalid symbol")
	}
}
