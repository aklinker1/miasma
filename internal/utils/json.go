package utils

import "encoding/json"

func ToJSON(object any) []byte {
	res, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}
	return res
}
