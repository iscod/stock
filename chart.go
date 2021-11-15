package main

import "io"
import "net/http"
import "encoding/json"

var charturi = "/v5/stock/chart/minute.json?period=1d&symbol="

func Getchartmessage(symbol string) (ChartMessage, error) {
	var chartmessage ChartMessage
	curl := stockhost + charturi + symbol
	cookie, err := GetCookie()

	if err != nil {
		return chartmessage, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", curl, nil)
	for i := 0; i < len(cookie); i++ {
		req.AddCookie(cookie[i])
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return chartmessage, err
	}

	dec := json.NewDecoder(resp.Body)

	if err != nil {
		return chartmessage, err
	}

	for {
		var m ChartMessage
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			return chartmessage, err
		}

		chartmessage = m
	}

	return chartmessage, err
}

//价格获取接口 get quote
func Getchart(symbol string) (Chart, error) {
	chartmessage, err := Getchartmessage(symbol)

	if err != nil {
		return chartmessage.Data, err
	}

	return chartmessage.Data, nil
}
