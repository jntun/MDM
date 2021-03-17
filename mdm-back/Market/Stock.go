package market

import (
    "fmt"
    "github.com/Nastyyy/mdm-back/config"
)

// Stock represents the stocks that make up a market
type Stock struct {
	Name   string  `json:"name"`
	Ticker string  `json:"ticker"`
	Price  float32 `json:"price"`
	Volume int     `json:"volume"`
        pfunc *PriceFunction
}

func (stock *Stock) CanBuy(volume int, balance float32) bool {
	if stock.Volume >= 0 && stock.Price*float32(volume) <= balance {
            return true
	}
	return false
}

func (stock *Stock) Tick(tick uint64) {
	// TODO
        err := stock.updatePrice(tick)
        if err != nil {
            config.StockLog(fmt.Sprintf("Error updating stock: %s", stock.Ticker))
        }
        config.StockLog(fmt.Sprintf("[%s] \n%s", stock.Ticker, stock.pfunc))
        //config.VerboseLog(fmt.Sprintf("[STCK] Price update for %s", stock.Ticker))
}

func (stock Stock) String() string {
	return fmt.Sprintf("%s: $%v - %v", stock.Ticker, stock.Price, stock.Volume)
}

func (stock *Stock) updatePrice(tick uint64) error {
        newPrice := stock.pfunc.NextPrice(float64(tick))
        config.VerboseLog(fmt.Sprintf("[NextPrice][%s] %f",  stock.Ticker, newPrice))
        stock.Price = newPrice
        return nil
}

func NewStock(name string, ticker string, startPrice ...float32) *Stock {
        stock := &Stock{Name: name, Ticker: ticker, Volume: 10000, pfunc: GeneratePriceFunc()}

	if startPrice[0] != 0 {
		stock.Price = startPrice[0]
	}

	return stock
}
