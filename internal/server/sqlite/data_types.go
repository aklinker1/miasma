package sqlite

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
)

// blob -> []string

type StringArray []string

func (array StringArray) Value() (driver.Value, error) {
	return json.Marshal(array)
}

func (array *StringArray) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	if bytes, ok := src.([]byte); ok {
		json.Unmarshal(bytes, array)
		return nil
	}

	return &server.Error{
		Code:    server.EINTERNAL,
		Message: fmt.Sprintf("Failed to scan %+v (%T) into a []string", src, src),
		Op:      "StringArray.Scan",
	}
}

// blob -> []int32

type Int32Array []int32

func (array Int32Array) Value() (driver.Value, error) {
	return json.Marshal(array)
}

func (array *Int32Array) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	if bytes, ok := src.([]byte); ok {
		json.Unmarshal(bytes, array)
		return nil
	}

	return &server.Error{
		Code:    server.EINTERNAL,
		Message: fmt.Sprintf("Failed to scan %+v (%T) into a []int32", src, src),
		Op:      "Int32Array.Scan",
	}
}

// blob -> []internal.BoundVolume

type BoundVolumeArray []internal.BoundVolume

func (array BoundVolumeArray) Value() (driver.Value, error) {
	return json.Marshal(array)
}

func (array *BoundVolumeArray) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	if bytes, ok := src.([]byte); ok {
		json.Unmarshal(bytes, array)
		return nil
	}

	return &server.Error{
		Code:    server.EINTERNAL,
		Message: fmt.Sprintf("Failed to scan %+v (%T) into a []internal.BoundVolume", src, src),
		Op:      "BoundVolumeArray.Scan",
	}
}

// blob -> internal.AppRouting

type AppRoutingBlob internal.AppRouting

func (blob *AppRoutingBlob) Value() (driver.Value, error) {
	if blob == nil {
		return nil, nil
	}
	return json.Marshal(blob)
}

func (blob *AppRoutingBlob) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	if bytes, ok := src.([]byte); ok {
		json.Unmarshal(bytes, blob)
		return nil
	}

	return &server.Error{
		Code:    server.EINTERNAL,
		Message: fmt.Sprintf("Failed to scan %+v (%T) into a internal.AppRouting", src, src),
		Op:      "AppRoutingBlob.Scan",
	}
}
