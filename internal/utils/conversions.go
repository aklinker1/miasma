package utils

import (
	"encoding/json"
)

func ToSQLiteBlob(obj any) []byte {
	if obj == nil {
		return nil
	}
	bytes, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return bytes
}

func FromSQLiteBlob[T any](blob []byte, receiver T) error {
	return json.Unmarshal(blob, &receiver)
}
