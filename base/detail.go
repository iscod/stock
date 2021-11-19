package base

import (
	"github.com/iscod/goStock/model"
	"io"
	"strings"
)
import "net/http"
import "encoding/json"

const stockhostqq string = "https://gu.qq.com"

var detail = "https://gu.qq.com/proxy/cgi/cgi-bin/yidong/getDadan?code="

func GetCookieQQ() ([]*http.Cookie, error) {
	resp, err := http.Get("https://gu.qq.com/sz000651/gp/dadan")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp.Cookies(), err
}

//价格获取接口 get DetailMessage
func GetDetailMessage(symbol string) (*model.DetailMessage, error) {
	var message model.DetailMessage
	url := detail + strings.ToLower(symbol)
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
	var m model.DetailMessage
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
func GetDetail(symbol string) (*model.DetailData, error) {
	message, err := GetDetailMessage(symbol)

	if err != nil {
		return nil, err
	}

	return &message.Data, nil
}
