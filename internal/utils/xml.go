package utils

import (
	"encoding/xml"
)

func Marshal(v interface{}) ([]byte, error) {
	return xml.MarshalIndent(v, "", "  ")
}

func Unmarshal(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}
