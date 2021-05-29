package custom_formats

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

type UInt32Array []uint32

func (blob *UInt32Array) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Scanned value was not a byte array: %v", value)
	}
	strValues := strings.Split(string(bytes), "|")
	nums := make([]uint32, len(strValues))
	for i, str := range strValues {
		num, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		nums[i] = uint32(num)
	}
	*blob = nums
	return nil
}

func (blob UInt32Array) Value() (driver.Value, error) {
	if len(blob) == 0 {
		return nil, nil
	}
	stringValues := make([]string, len(blob))
	for i, number := range blob {
		stringValues[i] = fmt.Sprint(number)
	}
	return []byte(strings.Join(stringValues, "|")), nil
}
