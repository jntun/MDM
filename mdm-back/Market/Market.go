package market

import (
	"fmt"

        "github.com/Nastyyy/mdm-back/config"
)

type Market struct {
	Name   string
	Ticker string
	Stocks []*Stock
	file   *FileHandler
}

func (market *Market) Update(tick uint64) {
	market.file.Save(market)

	for _, stock := range market.Stocks {
		stock.Tick(tick)
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
	market := &Market{Stocks: stocks, file: NewFileHandler()}

	if config.DEBUG_STOCK {
		market.AddStock(NewStock("Apple Inc.", "AAPL", 100.0))
		//market.AddStock(NewStock("Advanced Micro Devices", "AMD", 83.17))
		//market.AddStock(NewStock("Nvidia Corportaion", "NVDA", 552.46))
		//market.AddStock(NewStock("Tesla Inc", "TSLA", 555.46))
		//market.AddStock(NewStock("Nikola Corportaion", "NKLA", 34.86))
		//market.AddStock(NewStock("Dell Technologies", "DELL", 70.81))
	}

	return market
}
