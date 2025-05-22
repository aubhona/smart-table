package apperrors

import "fmt"

type IncorrectTableID struct {
	RequestTableID string
	ActualTableID  string
}

func (e IncorrectTableID) Error() string {
	return fmt.Sprintf("incorrect  request table_id='%s', actual table_id='%s'", e.RequestTableID, e.ActualTableID)
}
