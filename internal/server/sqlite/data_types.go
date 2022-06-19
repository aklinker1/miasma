package sqlite

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
)

// blob -> []string

type sqliteStringArray []string

func (array sqliteStringArray) Value() (driver.Value, error) {
	return json.Marshal(array)
}

func (array *sqliteStringArray) Scan(src interface{}) error {
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

type sqliteInt32Array []int32

func (array sqliteInt32Array) Value() (driver.Value, error) {
	return json.Marshal(array)
}

func (array *sqliteInt32Array) Scan(src interface{}) error {
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

type sqliteBoundVolumeArray []internal.BoundVolume

func (array sqliteBoundVolumeArray) Value() (driver.Value, error) {
	return json.Marshal(array)
}

func (array *sqliteBoundVolumeArray) Scan(src interface{}) error {
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

type sqliteAppRoutingBlob internal.AppRouting

func (blob *sqliteAppRoutingBlob) Value() (driver.Value, error) {
	if blob == nil {
		return nil, nil
	}
	return json.Marshal(blob)
}

func (blob *sqliteAppRoutingBlob) Scan(src interface{}) error {
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
