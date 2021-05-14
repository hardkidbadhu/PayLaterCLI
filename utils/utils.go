package utils

import (
	"encoding/json"
	"fmt"
)

type Utils interface {
	StructToMap(data interface{}) (map[string]interface{}, error)
}

type utils struct{}

func (u utils) StructToMap(data interface{}) (map[string]interface{}, error) {
	mp := map[string]interface{}{}

	byts, err := json.Marshal(data)
	if err != nil {
		fmt.Println("oh! we are working on it please try with valid input")
	}

	if err := json.Unmarshal(byts, &mp); err != nil {
		fmt.Println("oh! we are working on it please try with valid input")
	}

	return mp, nil
}

func NewUtils() Utils {
	return &utils{}
}