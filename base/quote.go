package base

import (
	"github.com/iscod/stock/model"
	"io"
)
import "net/http"
import "encoding/json"

var quteuri = "/v5/stock/quote.json?extend=detail&symbol="

//A QuoteMessage for quote.json
//    See https://stock.xueqiu.com/v5/stock/quote.json?symbol=SZ000651&extend=detail

//价格获取接口 get QuoteMessage
func GetQuoteMessage(symbol string) (*model.QuoteMessage, error) {
	var message model.QuoteMessage
	url := stockhost + quteuri + symbol
	cookie, err := GetCookie()

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	for i := 0; i < len(cookie); i++ {
		req.AddCookie(cookie[i])
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)

	var m model.QuoteMessage
	for {
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		message = m
	}

	return &message, nil
}

//价格获取接口 get quote
func GetQuote(symbol string) (*model.Quote, error) {
	message, err := GetQuoteMessage(symbol)

	if err != nil {
		return nil, err
	}

	return &message.Data.Quote, nil
}

//价格获取接口 get quote
func GetMarket(symbol string) (*model.Market, error) {
	message, err := GetQuoteMessage(symbol)

	if err != nil {
		return nil, err
	}

	return &message.Data.Market, nil
}
