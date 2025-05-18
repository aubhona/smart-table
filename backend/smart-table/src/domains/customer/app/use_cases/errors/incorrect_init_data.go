package apperrors

import (
	"fmt"
)

type IncorrectInitDataError struct {
	InitData string
}

func (err IncorrectInitDataError) Error() string {
	return fmt.Sprintf("Incorrect initData %v", err.InitData)
}
