package model

import (
	"database/sql/driver"
	"encoding/json"
)

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
