package model

type QuoteMessage struct {
	Data             QuoteDate
	ErrorDescription string `json:"Error_description"`
}

// message data struct
type QuoteDate struct {
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
	Id            int `json:"-"`
	Symbol        string
	CommentCount  int64           `json:"-"`
	CommentCount1 int64           `json:"-"`
	CommentCount2 int64           `json:"-"`
	CommentCount3 int64           `json:"-"`
	Volume        int64           `json:"-"`
	SummaryVolume SummaryVolumeKv `json:"-"`
	FundFlow      FundFlow        `json:"-"`
	Code          string          `gorm:"-"` //股票码
	// Type       int
	Name      string  //名字
	Open      float32 //开盘价
	Current   float32 //当前价格
	AvgPrice  float32 `json:"Avg_price"` //均价
	Low       float32 //当天最低
	High      float32 //当天最高价
	LastClose float32 `gorm:"-"` //昨收盘价
	High52w   float32 `gorm:"-"` //52周最高
	Low52w    float32 `gorm:"-"` //52周最低
	LimitDown float32 `gorm:"-"` //跌停价格
	LimitUp   float32 `gorm:"-"` //涨停价格

	// Tick_size             float32
	Amplitude float32 `gorm:"-"` //振幅
	// Current_year_percent  float32
	Market_capital      float32 `gorm:"-"` //总市值
	StringMarketCapital float32 `gorm:"-"` //流通值
	// Issue_date            int
	// Sub_type              string
	// Currency              string
	// Lot_size              int
	// // Profit string //2.240157620494E10
	// Timestamp int    //时间戳
	Chg     float32 `gorm:"-"` //涨跌额
	Percent float32 `gorm:"-"` //涨跌幅度
	// // Profit_four string //2.805930302493E10
	// Volume          int    //成交量
	// Volume_ratio    string ////量比
	// Profit_forecast string
	// Turnover_rate   string
	Navps          float32 `gorm:"-"` //每股净资产
	Dividend_yield float32 `gorm:"-"` //股息收益率-股息率
	Eps            float32 `gorm:"-"` //每股收益
	Pb             float32 `gorm:"-"` //市净率
	Pe_lyr         float32 `gorm:"-"` //市盈率(静)
	Pe_ttm         float32 `gorm:"-"` //市盈率(TTM)
	// Exchange        string //证交所代码,SZ,SH
	// Pe_forecast     string
	Time      int64  `gorm:"-"` //时间戳
	ExecAt    string `json:"-"` //创建时间
	CreatedAt int64  `json:"-"` //创建时间
	// Total_shares    int    //总股本
	// string_shares   string //流通股本
	// status          int
}
