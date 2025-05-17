package domainerrors

import "fmt"

type IncorrectDeleteItemsCount struct {
	Count int
}

func (e IncorrectDeleteItemsCount) Error() string {
	return fmt.Sprintf("incorrect delete items count %d", e.Count)
}
