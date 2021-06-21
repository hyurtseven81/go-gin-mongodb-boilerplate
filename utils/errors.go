package utils

import "fmt"

type DataPadError struct {
	StatusCode int

	Err error
}

func (r *DataPadError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}
