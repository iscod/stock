// XueQiu comment message get and save

package xueqiu

import "io"
import "net/http"
import "encoding/json"

var quteuri = "/v5/stock/quote.json?extend=detail&symbol="

//价格获取接口 get QuoteMessage
func GetQuoteMessage(symbol string) (QuoteMessage, error) {
	var quotemessage QuoteMessage
	url := stockhost + quteuri + symbol
	cookie, err := GetCookie()

	if err != nil {
		return quotemessage, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	for i := 0; i < len(cookie); i++ {
		req.AddCookie(cookie[i])
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return quotemessage, err
	}

	dec := json.NewDecoder(resp.Body)

	if err != nil {
		return quotemessage, err
	}

	for {
		var m QuoteMessage
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			return quotemessage, err
		}

		quotemessage = m
	}

	return quotemessage, nil
}

//价格获取接口 get quote
func Getquote(symbol string) (Quote, error) {
	quotemessage, err := GetQuoteMessage(symbol)

	if err != nil {
		return quotemessage.Data.Quote, err
	}

	return quotemessage.Data.Quote, nil
}

//价格获取接口 get quote
func Getmarket(symbol string) (Market, error) {
	quotemessage, err := GetQuoteMessage(symbol)

	if err != nil {
		return quotemessage.Data.Market, err
	}

	return quotemessage.Data.Market, nil
}

//A QuoteMessage for quote.json
//    See https://stock.xueqiu.com/v5/stock/quote.json?symbol=SZ000651&extend=detail
type QuoteMessage struct {
	Data              Quotedate
	Error_description string
}

// message data struct
type Quotedate struct {
	Market Market
	Quote  Quote
}

//Market info struct
type Market struct {
	Status_id int
	Region    string
	Status    string
	Time_zone string
}

//Quote info struct
type Quote struct {
	Symbol string
	Code   string //股票码
	// Type       int
	Name       string  //名字
	Open       float32 //开盘价
	Current    float32 //当前价格
	Avg_price  float32 //均价
	Low        float32 //当天最低
	High       float32 //当天最高价
	Last_close float32 //昨收盘价
	High52w    float32 //52周最高
	Low52w     float32 //52周最低
	Limit_down float32 //跌停价格
	Limit_up   float32 //涨停价格

	// Tick_size             float32
	Amplitude float32 //振幅
	// Current_year_percent  float32
	Market_capital        float32 //总市值
	string_market_capital float32 //流通值
	// Issue_date            int
	// Sub_type              string
	// Currency              string
	// Lot_size              int
	// // Profit string //2.240157620494E10
	// Timestamp int    //时间戳
	Chg     float32 //涨跌额
	Percent float32 //涨跌幅度
	// // Profit_four string //2.805930302493E10
	// Volume          int    //成交量
	// Volume_ratio    string ////量比
	// Profit_forecast string
	// Turnover_rate   string
	Navps          float32 //每股净资产
	Dividend_yield float32 //股息收益率-股息率
	Eps            float32 //每股收益
	Pb             float32 //市净率
	Pe_lyr         float32 //市盈率(静)
	Pe_ttm         float32 //市盈率(TTM)
	// Exchange        string //证交所代码,SZ,SH
	// Pe_forecast     string
	Time int //时间戳
	// Total_shares    int    //总股本
	// string_shares   string //流通股本
	// status          int
}
