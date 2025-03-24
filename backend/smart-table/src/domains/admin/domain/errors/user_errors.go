package errors

import "fmt"

// Ошибка для случая, когда пользователь с данным логином уже существует
type UserAlreadyExistsError struct {
	Login string
}

func (e UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("user with login '%s' already exists", e.Login)
}

// Ошибка для случая, когда пользователь с данной электронной почтой уже существует
type EmailAlreadyExistsError struct {
	Email string
}

func (e EmailAlreadyExistsError) Error() string {
	return fmt.Sprintf("user with email '%s' already exists", e.Email)
}
