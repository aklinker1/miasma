package shared

import (
	"encoding/json"

	"github.com/aklinker1/miasma/internal/shared/log"
)

func ConvertUInt32ArrayToInt64Array(input []uint32) []int64 {
	output := []int64{}
	for _, inputItem := range input {
		output = append(output, int64(inputItem))
	}
	return output
}

func ConvertInt64ArrayToUInt32Array(input []int64) []uint32 {
	output := []uint32{}
	for _, inputItem := range input {
		output = append(output, uint32(inputItem))
	}
	return output
}

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
