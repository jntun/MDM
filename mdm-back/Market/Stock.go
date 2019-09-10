package market

// Stock represents the stocks that make up a market
type Stock struct {
	Name   string  `json:"name"`
	Symbol string  `json:"symbol"`
	Price  float32 `json:"price"`
}
