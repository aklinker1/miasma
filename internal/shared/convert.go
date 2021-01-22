package shared

func ConvertUInt32ArrayToInt64Array(input []uint32) (output []int64) {
	for _, inputItem := range input {
		output = append(output, int64(inputItem))
	}
	return output
}

func ConvertInt64ArrayToUInt32Array(input []int64) (output []uint32) {
	for _, inputItem := range input {
		output = append(output, uint32(inputItem))
	}
	return output
}
