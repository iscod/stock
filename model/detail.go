package model

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
