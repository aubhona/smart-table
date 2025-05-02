package errors

import "fmt"

type UserNotFoundByLogin struct {
	Login string
}

func (e UserNotFoundByLogin) Error() string {
	return fmt.Sprintf("User not found by login: %s", e.Login)
}
