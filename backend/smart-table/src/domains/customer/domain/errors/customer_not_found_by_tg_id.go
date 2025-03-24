package domainerrors

import "fmt"

type CustomerNotFoundByTgID struct {
	TgID string
}

func (c CustomerNotFoundByTgID) Error() string {
	return fmt.Sprintf("Customer not found by TgId: %s", c.TgID)
}
