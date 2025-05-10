package errors

import "fmt"

type UnsuccessMenuDishFetch struct {
	InnerError error
}

func (u UnsuccessMenuDishFetch) Error() string {
	return fmt.Sprintf("error while fetching menu dish list: %s", u.InnerError.Error())
}
