package errors

type InvalidToken struct{}

func (c InvalidToken) Error() string {
	return "Token is invalid"
}
