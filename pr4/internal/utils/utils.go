package utils

import (
	"encoding/json"
	"errors"
)

func ToPtr[T any](val T) *T {
	result := new(T)
	*result = val
	return result
}

type Range[T any] struct {
	From, To *T
}

func (r *Range[T]) UnmarshalJSON(data []byte) error {
	var x []*T
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if len(x) != 2 {
		return errors.New("incorrect format of range")
	}
	r.From, r.To = x[0], x[1]
	return nil
}
