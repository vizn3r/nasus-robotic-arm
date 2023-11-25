package util

import (
	"encoding/json"
	"os"
)

// This is for easier JSON manipulation

func ToJSON(data any, dest string) error {
	b, err := json.Marshal(data);
	if err != nil {
		return err
	}
	os.WriteFile(dest, b, 0777)
	return nil
}

func ParseJSON(dest any, source string) error {
	b, err := os.ReadFile(source)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, dest)
	if err != nil {
		return err
	}
	return nil
}