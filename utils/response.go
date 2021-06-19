package utils

type Data struct {
	Items interface{} `json:"items"`
	Count int64       `json:"count"`
}

type DataResult struct {
	Data Data `json:"data"`
}

func NewDataresult(items interface{}, count int64) DataResult {
	data := Data{
		Items: items,
		Count: count,
	}
	dataResult := DataResult{
		Data: data,
	}

	return dataResult
}
