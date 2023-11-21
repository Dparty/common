package data

import (
	"database/sql/driver"
	"encoding/json"
)

type Pair[L, R any] struct {
	L L `json:"left"`
	R R `json:"right"`
}

func (Pair[L, R]) GormDataType() string {
	return "JSON"
}

func (s *Pair[L, R]) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s Pair[L, R]) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}
