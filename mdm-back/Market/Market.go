package market

import (
	"fmt"
	"os"
)

type Market struct {
	Name   string
	Ticker string
	Stocks []*Stock
}

func (market *Market) Update() {
	for _, stock := range market.Stocks {
		stock.Tick()
	}
}

func (market *Market) GetStock(ticker string) (stock *Stock, err error) {
	for _, stock := range market.Stocks {
		if stock.Ticker == ticker {
			return stock, nil
		}
	}

	return nil, nil
}

func (market *Market) AddStock(stock *Stock) {
	market.Stocks = append(market.Stocks, stock)
}

func (market Market) String() string {
	ret := "| "
	for _, stock := range market.Stocks {
		ret += fmt.Sprintf("%s: %v | ", stock.Ticker, stock.Price)
	}
	return ret
}

func NewMarket(stocks ...*Stock) *Market {
	market := &Market{Stocks: stocks}

	if os.Getenv("DEBUG") == "true" {
		market.AddStock(NewStock("Apple Inc.", "AAPL", 100.0))
		market.AddStock(NewStock("Advanced Micro Devices", "AMD", 83.17))
		market.AddStock(NewStock("Nvidia Corportaion", "NVDA", 552.46))
	}

	return market
}
