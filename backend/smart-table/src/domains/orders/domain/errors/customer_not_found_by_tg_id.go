package domain_errors

import "fmt"

type CustomerNotFoundByTgID struct {
	TgId string
}

func (c CustomerNotFoundByTgID) Error() string {
	return fmt.Sprintf("Customer not found by TgId: %s", c.TgId)
}
