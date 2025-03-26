package apperrors

import (
	"fmt"
)

type CustomerAlreadyExist struct {
	TgID string
}

func (err CustomerAlreadyExist) Error() string {
	return fmt.Sprintf("Customer with tg id %s already exist", err.TgID)
}
