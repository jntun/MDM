package market

import (
	"encoding/json"
	"io/ioutil"
)

const directory = "./stockmath/files/"
const requestFile = directory + "request.json"
const resposneFile = requestFile + "response.json"

type PriceUpdate struct {
	Ticker   string   `json:"ticker"`
	LogEvent LogEvent `json:"event,omitempty"`
	Price    float32  `json:"priceUpdate,omitempty"`
}

type LogEvent string
type LogEntry []PriceUpdate

type FileHandler struct {
	RequestLog  map[int]*LogEntry
	ResponseLog map[int]*LogEntry
	current     int
}

func (handler *FileHandler) Save(market *Market) error {

	logEntry := handler.mapToLogEntry(market)
	handler.RequestLog[handler.current] = logEntry

	jsonData, err := json.MarshalIndent(logEntry, "", "")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(requestFile, jsonData, 0644)

	handler.current += 1
	return nil
}

func (handler *FileHandler) Load() {
}

func (handler *FileHandler) mapToLogEntry(market *Market) *LogEntry {
	log := LogEntry{}
	for _, stock := range market.Stocks {
		log = append(log, handler.mapToPriceUpdate(*stock))
	}

	return &log
}

func (handler *FileHandler) mapToPriceUpdate(stock Stock) PriceUpdate {
	return PriceUpdate{Ticker: stock.Ticker, LogEvent: "null"}
}

func NewFileHandler() *FileHandler {
	requestLog := make(map[int]*LogEntry)
	responseLog := make(map[int]*LogEntry)
	return &FileHandler{RequestLog: requestLog, ResponseLog: responseLog, current: 0}
}
