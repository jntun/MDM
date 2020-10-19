package market

// Stock represents the stocks that make up a market
type Stock struct {
	Name   string  `json:"name"`
	Ticker string  `json:"ticker"`
	Price  float32 `json:"price"`
	Volume int32   `json:"volume"`
}

func (stock *Stock) priceUpdate() {
	// TODO
	stock.Price += 1
}

func NewStock(name string, ticker string, startPrice ...float32) *Stock {
	stock := &Stock{Name: name, Ticker: ticker}

	if startPrice[0] != 0 {
		stock.Price = startPrice[0]
	}

	return stock
}
