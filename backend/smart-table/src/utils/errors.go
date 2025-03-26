package utils

func IsTheSameErrorType[T any](err error) bool {
	_, ok := err.(T)

	return ok
}
