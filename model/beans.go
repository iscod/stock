package model

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
)

type StringSlice []string

func (this StringSlice) MarshalJSON() ([]byte, error) {
	return []byte("[" + strings.Join(this, ",") + "]"), nil
}

func (this *StringSlice) UnmarshalJSON(data []byte) error {
	if len(data) <= 2 {
		*this = make([]string, 0)
		return nil
	}
	*this = strings.Split(string(data[1:len(data)-1]), ",")
	return nil
}

func (this *StringSlice) Scan(value interface{}) error {
	*this = make([]string, 0)
	if value == nil {
		return nil
	}
	return this.UnmarshalJSON(value.([]byte))
}

type StringSlices []StringSlice

func (this *StringSlices) Scan(value interface{}) error {
	*this = StringSlices{}
	if value == nil {
		return nil
	}
	return json.Unmarshal(value.([]byte), this)
}

func (this StringSlices) Value() (driver.Value, error) {
	return json.Marshal(this)
}
