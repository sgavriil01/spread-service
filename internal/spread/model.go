package spread

import "time"

type Spread struct {
	Symbol    string    `json:"symbol"`
	Spread    float64   `json:"spread"`
	UpdatedAt time.Time `json:"updated_at"`
}

var Symbols = []string{
	"EURUSD",
	"EURCAD",
	"USDJPY",
	"BTCUSD",
	"XAUUSD",
}