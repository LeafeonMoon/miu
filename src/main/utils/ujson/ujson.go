package ujson

import (
	"encoding/json"
)

func ToJson(val interface{}) (string, error) {
	jsonBytes, err := json.Marshal(val)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), err
}

func ToObject(jsonStr string, object interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), &object)
	if err != nil {
		return err
	}
	return nil
}
