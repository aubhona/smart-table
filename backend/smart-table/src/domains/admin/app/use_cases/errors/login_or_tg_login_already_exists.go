package errors

import (
	"fmt"
)

type LoginOrTgLoginAlreadyExists struct {
	Login   string
	TgLogin string
}

func (e LoginOrTgLoginAlreadyExists) Error() string {
	return fmt.Sprintf("user with login '%s' or tg_login '%s' already exists", e.Login, e.TgLogin)
}
