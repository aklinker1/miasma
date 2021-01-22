package shared

func ConvertUInt32ArrayToInt64Array(input []uint32) (output []int64) {
	for _, inputItem := range input {
		output = append(output, int64(inputItem))
	}
	return output
}
