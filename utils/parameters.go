package utils

import (
	"encoding/json"
	"log"
)

type Query struct {
	Filter string `form:"filter"`
	Select string `form:"select"`
	Sort   string `form:"sort"`
	Skip   int    `form:"skip"`
	Limit  int    `form:"limit"`
}

func toJSON(data string) map[string]interface{} {
	if data == "" {
		return nil
	}
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(data), &raw)

	if err != nil {
		log.Printf("Error while parsing '%s': %v", data, err)
	}

	return raw
}

func (q Query) GetSelect() interface{} {
	return toJSON(q.Select)
}

func (q Query) GetFilter() interface{} {
	return toJSON(q.Filter)
}

func (q Query) GetSort() interface{} {
	return toJSON(q.Sort)
}
