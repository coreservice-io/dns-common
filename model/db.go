package model

import (
	"database/sql/driver"
	"errors"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type StringArray []string

func (sa *StringArray) Scan(value interface{}) error {
	bytesValue, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal StringArray value:", value))
	}
	return json.Unmarshal(bytesValue, sa)
}

// Value return json value, implement driver.Valuer interface
func (sa StringArray) Value() (driver.Value, error) {
	result, err := json.Marshal(sa)
	return string(result), err
}
