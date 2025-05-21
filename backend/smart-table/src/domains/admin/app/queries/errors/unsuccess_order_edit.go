package errors

import "fmt"

type UnsuccessOrderEdit struct {
	InnerError error
}

func (u UnsuccessOrderEdit) Error() string {
	return fmt.Sprintf("error while fetching order list: %s", u.InnerError.Error())
}
