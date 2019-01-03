package xueqiu

import "net/http"
import "encoding/json"

const host string = "https://xueqiu.com"
const stockhost string = "https://stock.xueqiu.com"

func GetCookie() ([]*http.Cookie, error) {
	resp, err := http.Get("http://xueqiu.com/s/SZ000651")
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return resp.Cookies(), err
}

func Get(url string) (*json.Decoder, error) {
	curl := host + url
	cookie, err := GetCookie()

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", curl, nil)
	for i := 0; i < len(cookie); i++ {
		req.AddCookie(cookie[i])
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)

	return dec, nil
}
