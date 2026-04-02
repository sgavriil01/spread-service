package spread

import (
	"slices"
	"time"
)

var spreads = map[string]Spread{}

func IsValidSymbol(symbol string) bool {
	return slices.Contains(Symbols, symbol)
}

func GetSpread(symbol string) (Spread, bool) {
	if !IsValidSymbol(symbol) {
		return Spread{}, false
	}

	spread, exists := spreads[symbol]

	if !exists {
		return Spread{}, false
	}

	return spread, true
}

func SetSpread(symbol string, spreadValue float64) (Spread, bool){
	if !IsValidSymbol(symbol) {
		return Spread{}, false
	}

	spread := Spread{
		Symbol: symbol,
		Spread: spreadValue,
		UpdatedAt: time.Now().UTC(),
	}

	spreads[symbol] = spread
	return spread, true
}