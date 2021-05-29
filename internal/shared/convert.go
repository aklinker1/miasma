package shared

import (
	"encoding/json"

	"github.com/aklinker1/miasma/internal/shared/log"
)

func StringPtr(input string) *string {
	return &input
}
func BoolPtr(input bool) *bool {
	return &input
}

func StructToMap(input interface{}) (m map[string]interface{}) {
	str, err := json.Marshal(input)
	if err != nil {
		log.W("Failed during marshal step of StructToMap: %v", err)
	}
	err = json.Unmarshal(str, &m)
	if err != nil {
		log.W("Failed during unmarshal step of StructToMap: %v", err)
	}
	return m
}
