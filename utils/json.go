package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type JSON map[string]interface{}

func (pc *JSON) Scan(val interface{}) error {
	switch v := val.(type) {	// interface.(type) 判断类型
	case []byte:
		json.Unmarshal(v, &pc)
		return nil
	case string:
		json.Unmarshal([]byte(v), &pc)
		return nil
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (pc *JSON) Value() (driver.Value, error) {
	return json.Marshal(pc)
}