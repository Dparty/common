package data

import (
	"database/sql/driver"
	"encoding/json"
)

type List[T any] []T

func (s List[T]) Append(t T) List[T] {
	return append(s, t)
}

func (s List[T]) Len() int {
	return len(s)
}

func (List[T]) GormDataType() string {
	return "JSON"
}

func (s *List[T]) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s List[T]) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}
