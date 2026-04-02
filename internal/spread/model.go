package spread

import "time"

type Spread struct {
	Symbol    string    `json:"symbol"`
	Spread    float64   `json:"spread"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SetSpreadRequest struct {
	Spread float64 `json:"spread"`
}

var Symbols = []string{
	"EURUSD",
	"EURCAD",
	"USDJPY",
	"BTCUSD",
	"XAUUSD",
}