package model

import (
	"database/sql/driver"
	"encoding/json"
)

type GetFundFlowMessage struct {
	Data FundFlowData `json:"data"`
	Code int          `json:"code"`
}

type FundFlowData struct {
	TodayFundFlow   TodayFundFlow   `json:"todayFundFlow"`
	TodayFundTrend  TodayFundTrend  `json:"todayFundTrend"`
	HistoryFundFlow HistoryFundFlow `json:"historyFundFlow"`
}

type TodayFundFlow struct {
	MainNetIn string `json:"mainNetIn"`
	MainIn    string `json:"mainIn"`
	MainOut   string `json:"mainOut"`
	RetailIn  string `json:"retailIn"`
	RetailOut string `json:"retailOut"`
}

type TodayFundTrend struct {
}

type HistoryFundFlow struct {
	OneDayKlineList []OneDayKlineListData `json:"oneDayKlineList"`
}

type OneDayKlineListData struct {
	Date      string `json:"date"`
	MainNetIn string `json:"mainNetIn"`
	AvgIn     string `json:"avgIn"`
}

type FundFlow struct {
	MainNetIn     string
	AvgIn         string
	TodayFundFlow TodayFundFlow
}

func (this *FundFlow) Scan(value interface{}) error {
	*this = FundFlow{}
	if value == nil {
		return nil
	}
	return json.Unmarshal(value.([]byte), this)
}

func (this FundFlow) Value() (driver.Value, error) {
	return json.Marshal(this)
}

type DetailMessage struct {
	Data DetailData `json:"data"`
	Code int        `json:"code"`
}

type DetailData struct {
	Summary Summary      `json:"summary"`
	Detail  StringSlices `json:"detail"`
}

type Summary struct {
	Date   string `json:"date"`
	T      string `json:"time"`
	Volume string `json:"volume"`
}

type SummaryVolume struct {
	BVolume int64 `json:"b"`
	SVolume int64 `json:"s"`
	MVolume int64 `json:"m"`
}

type SummaryVolumeKv map[int]SummaryVolume

func (this *SummaryVolumeKv) Scan(value interface{}) error {
	*this = SummaryVolumeKv{}
	if value == nil {
		return nil
	}
	return json.Unmarshal(value.([]byte), this)
}

func (this SummaryVolumeKv) Value() (driver.Value, error) {
	return json.Marshal(this)
}
