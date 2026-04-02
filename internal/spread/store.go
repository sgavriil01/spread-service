package spread

import "slices"

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
