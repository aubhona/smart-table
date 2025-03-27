package errors

type IncorrectPassword struct{}

func (e IncorrectPassword) Error() string {
	return "Incorrect password"
}
