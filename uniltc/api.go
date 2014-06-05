package uniltc

// order type
const (
	Bid = "bid"
	Ask = "ask"
)

// candle stick type
const (
	C1m = "1m"
	C5m = "5m"
	C1h = "1h"
	C1d = "1d"
)

// trading pairs
const (
	Pair_LTC_BTC = "LTC_BTC"
	Pair_LTC_USD = "LTC_USD"
	Pair_LTC_EUR = "LTC_EUR"
)

var (
	Stub = new(stub)
)

// divisions
const (
	Div_Total  = 1e8
	Div_Volume = 1e4
	Div_Rate   = 1e4
)

// Noted:
// Timestamp is the ticker last updated
// Last, High, Low, Buy, Sell, Volume, Average are equal to the real value * division
type ticker struct {
	Timestamp                                   int64
	Last, High, Low, Buy, Sell, Volume, Average float64
}

type history struct {
	Timestamp           int64
	Type                string
	Rate, Volume, Total float64
}

type order struct {
	OrderId             int64
	OrderType           string
	Rate, Volume, Total float64
}

type depth struct {
	Timestamp int64
	OrderBook map[string][]order
}

type candleStick struct {
	Timestamp                                                 int64
	Open, High, Low, Close, BidVolume, AskVolume, TotalVolume float64
}

type stub struct {
	// public api
	Ticker        func(pair string) (ticker, error)
	LastTrade     func(pair string) (history, error)
	OrderBook     func(pair string) (depth, error)
	OrderBookTopN func(pair string, n int) (depth, error)
	History       func(pair string, n int) ([]history, error)
	HistorySince  func(pair string, since int64) ([]history, error)
	CandleStick   func(pair string, candleStickType string) ([]candleStick, error)

	// private api
	ApiNewBidOrder    func(api_key string, res string, pair string, rate int64, amount int64) ([]order, interface{}, error)
	ApiNewAskOrder    func(api_key string, res string, pair string, rate int64, amount int64) ([]order, interface{}, error)
	ApiCancelBidOrder func(api_key string, res string, pair string, orderId int64) error
	ApiCancelAskOrder func(api_key string, res string, pair string, orderId int64) error
	ApiCurrentOrders  func(api_key string, res string, pair string) (map[string][]order, error)
	ApiHistoryTrades  func(api_key string, res string, pair string) ([]history, error)
}
