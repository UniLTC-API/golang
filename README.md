UniLTC Api Golang SDK
======

#### How to install 

```go get -u github.com/UniLTC-API/golang/uniltc```

#### How to use

```go
package main

import (
	"log"

	"github.com/UniLTC-API/golang/uniltc"
	"github.com/hprose/hprose-go/hprose"
)

// cresdential
const (
	key = "your-key"
	sec = "your-sec"
)

// host
const (
	// possible value
	// http://usapi.uniltc.com/	http://londonapi.uniltc.com/
	// https://usapi.uniltc.com/	https://londonapi.uniltc.com/
	// tcp://usapi.uniltc.com:9090/ tcp://londonapi.uniltc.com:9090/
	// tcp://usapi.uniltc.com:9091/ tcp://londonapi.uniltc.com:9091/	(tcp with SSL secure)
	host = "tcp://londonapi.uniltc.com:9090/"
)

var rpcClient = hprose.NewClient(host)

func init() {
	rpcClient.UseService(uniltc.Stub)
}

// public api
func ticker() {
	// ticker
	val, err := uniltc.Stub.Ticker(uniltc.Pair_LTC_USD)
	if err != nil {
		log.Fatalf("fetch ticker error %v", err)
	}
	log.Printf("Last: %f", val.Last/uniltc.Div_Rate)
	log.Printf("High: %f", val.High/uniltc.Div_Rate)
	log.Printf("Low: %f", val.Low/uniltc.Div_Rate)
	log.Printf("Buy: %f", val.Buy/uniltc.Div_Rate)
	log.Printf("Sell: %f", val.Sell/uniltc.Div_Rate)
	log.Printf("Volume: %f", val.Volume/uniltc.Div_Volume)
	log.Printf("24-hour-average: %f", val.Average/uniltc.Div_Rate)
	log.Printf("=====================")
	log.Printf("Done ticker testing")
	log.Println()
	log.Println()
}

func orderBook() {
	// orderbook
	ob, err := uniltc.Stub.OrderBook(uniltc.Pair_LTC_USD)
	if err != nil {
		log.Fatalf("fetch depth error %v", err)
	}
	log.Printf("Bid Orders:")
	for _, v := range ob.OrderBook[uniltc.Bid] {
		log.Printf("Rate: %f	Volume %f", v.Rate/uniltc.Div_Rate, v.Volume/uniltc.Div_Volume)
	}
	log.Printf("Ask Orders:")
	for _, v := range ob.OrderBook[uniltc.Ask] {
		log.Printf("Rate: %f	Volume %f", v.Rate/uniltc.Div_Rate, v.Volume/uniltc.Div_Volume)
	}
	log.Printf("====================")
	log.Printf("Done depth testing")
	log.Println()
	log.Println()
}

func lastTrade() {
	// last trade
	lt, err := uniltc.Stub.LastTrade(uniltc.Pair_LTC_USD)
	if err != nil {
		log.Fatalf("fetch last trade error %v", err)
	}
	log.Printf("Last trade timestamp: %d", lt.Timestamp)
	log.Printf("Last trade type: %s", lt.Type)
	log.Printf("Last trade rate: %f", lt.Rate/uniltc.Div_Rate)
	log.Printf("Last trade volume:  %f", lt.Volume/uniltc.Div_Volume)
	log.Printf("Last trade total: %f", lt.Total/uniltc.Div_Total)
	log.Printf("====================")
	log.Printf("Done last trade testing")
	log.Println()
	log.Println()
}

func history() {
	// history
	h, err := uniltc.Stub.History(uniltc.Pair_LTC_USD, 10)
	if err != nil {
		log.Fatalf("fetch history error %v", err)
	}
	log.Printf("The last 10 trades")
	for _, v := range h {
		log.Printf("timestamp: %d   rate: %f   volume: %f", v.Timestamp, v.Rate/uniltc.Div_Rate, v.Volume/uniltc.Div_Volume)
	}
	log.Printf("====================")
	log.Printf("Done history testing")
	log.Println()
	log.Println()
}

func historySince() {
	// history since
	hs, err := uniltc.Stub.HistorySince(uniltc.Pair_LTC_USD, -1)
	if err != nil {
		log.Fatalf("fetch history since timestamp error %v", err)
	}
	log.Printf("History since the begining")
	for _, v := range hs {
		log.Printf("timestamp: %d   rate: %f   volume: %f", v.Timestamp, v.Rate/uniltc.Div_Rate, v.Volume/uniltc.Div_Volume)
	}
	log.Printf("====================")
	log.Printf("Done history since testing")
	log.Println()
	log.Println()
}

func candleStick() {
	// 1-minute candle stick
	cs, err := uniltc.Stub.CandleStick(uniltc.Pair_LTC_USD, uniltc.C1m)
	if err != nil {
		log.Fatalf("fetch candle stick error %v", err)
	}
	log.Printf("Candle stick")
	for _, v := range cs {
		log.Printf("timestamp: %d   open: %f   high: %f   low: %f   close: %f   bid-volume: %f   ask-volume: %f   total-volume: %f", v.Timestamp, v.Open/uniltc.Div_Rate, v.High/uniltc.Div_Rate, v.Low/uniltc.Div_Rate, v.Close/uniltc.Div_Rate, v.BidVolume/uniltc.Div_Volume, v.AskVolume/uniltc.Div_Volume, v.TotalVolume/uniltc.Div_Volume)
	}
	log.Printf("====================")
	log.Printf("Done candle stick testing")
	log.Println()
	log.Println()
}

// private api
func newBidOrder() {

}

func main() {
	ticker()
	orderBook()
	lastTrade()
	history()
	historySince()
	candleStick()
}
```
