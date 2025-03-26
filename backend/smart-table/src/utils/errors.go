package utils

func IsTheSameErrorType[T any](err error, targetErr T) bool {
	_, ok := err.(T)

	return ok
}
