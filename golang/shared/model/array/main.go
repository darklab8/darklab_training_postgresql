package array

// from https://github.com/FlowerLab/sorbifolia/blob/main/gorm-datatype/array_string.go

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type Array []string

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *Array) Scan(value interface{}) error {
	bytes, ok := value.([]string)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	*j = bytes
	return nil
}

// Value return json value, implement driver.Valuer interface
func (j Array) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return fmt.Sprintf("{%s}", strings.Join([]string(j), ",")), nil
}
