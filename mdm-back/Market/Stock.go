package market

import "fmt"

// Stock represents the stocks that make up a market
type Stock struct {
	Name   string  `json:"name"`
	Ticker string  `json:"ticker"`
	Price  float32 `json:"price"`
	Volume int     `json:"volume"`
}

func (stock *Stock) CanBuy(volume int, balance float32) bool {
	if stock.Volume >= 0 {
		if stock.Price*float32(volume) < balance {
			return true
		}
	}
	return false
}

func (stock *Stock) Tick() {
	// TODO
	stock.Price += float32(3)
}

func (stock Stock) String() string {
	return fmt.Sprintf("%s: $%v - %v", stock.Ticker, stock.Price, stock.Volume)
}

func NewStock(name string, ticker string, startPrice ...float32) *Stock {
	stock := &Stock{Name: name, Ticker: ticker, Volume: 10000}

	if startPrice[0] != 0 {
		stock.Price = startPrice[0]
	}

	return stock
}
