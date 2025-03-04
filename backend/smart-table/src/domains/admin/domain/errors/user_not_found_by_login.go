package domain_errors

import "fmt"

type UserNotFoundByLogin struct {
	Login string
}

func (c UserNotFoundByLogin) Error() string {
	return fmt.Sprintf("User not found by login: %s", c.Login)
}
