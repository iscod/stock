package base

import (
	"github.com/iscod/stock/model"
	"io"
	"strings"
)
import "net/http"
import "encoding/json"


var fund_flow = "https://gu.qq.com/proxy/cgi/cgi-bin/fundflow/hsfundtab?code="

//价格获取接口 get DetailMessage
func GetFundFlowMessage(symbol string) (*model.GetFundFlowMessage, error) {
	var message model.GetFundFlowMessage
	url := fund_flow + strings.ToLower(symbol)
	cookie, err := GetCookieQQ()

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	for i := 0; i < len(cookie); i++ {
		req.AddCookie(cookie[i])
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)
	defer resp.Body.Close()
	var m model.GetFundFlowMessage
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