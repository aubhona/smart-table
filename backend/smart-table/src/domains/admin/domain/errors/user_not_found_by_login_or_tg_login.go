package errors

import "fmt"

type UserNotFoundByLoginOrTgLogin struct {
	Login   string
	TgLogin string
}

func (e UserNotFoundByLoginOrTgLogin) Error() string {
	return fmt.Sprintf("User not found by login: %s, tg login: %s", e.Login, e.TgLogin)
}
