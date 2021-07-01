package custom_formats

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type StringArray []string

func (blob *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Scanned value was not a byte array: %v", value)
	}
	*blob = strings.Split(string(bytes), "|")
	return nil
}

func (blob StringArray) Value() (driver.Value, error) {
	if len(blob) == 0 {
		return nil, nil
	}
	return []byte(strings.Join(blob, "|")), nil
}
