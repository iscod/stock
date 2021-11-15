package model

type ChartMessage struct {
	Data Chart
}

type Chart struct {
	//Chart 价格图表结构体
	LastClose float32 `json:"last_close"`
	Items     []Item
}

type Item struct {
	//Chart 价格图表item结构体
	Current   float32 //当前价
	Volume    int     //成交量
	AvgPrice  float32 `json:"avg_price"` //均价
	Chg       float32 //涨跌额
	Percent   float32 //涨跌幅度
	Timestamp int64   //时间戳
	Amount    float64
}
