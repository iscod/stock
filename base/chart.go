package base

import (
	"github.com/iscod/stock/model"
	"io"
)
import "net/http"
import "encoding/json"

const host string = "https://xueqiu.com"
const stockhost string = "https://stock.xueqiu.com"

var charturi = "/v5/stock/chart/minute.json?period=1d&symbol="

func Getchartmessage(symbol string) (*model.ChartMessage, error) {
	var chartmessage model.ChartMessage
	curl := stockhost + charturi + symbol
	cookie, err := GetCookie()

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", curl, nil)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(cookie); i++ {
		req.AddCookie(cookie[i])
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	var m model.ChartMessage
	for {
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		chartmessage = m
	}

	return &chartmessage, err
}

//价格获取接口 get quote
func GetChart(symbol string) (*model.Chart, error) {
	message, err := Getchartmessage(symbol)

	if err != nil {
		return nil, err
	}

	return &message.Data, nil
}
