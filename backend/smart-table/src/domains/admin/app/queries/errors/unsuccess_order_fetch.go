package errors

import "fmt"

type UnsuccessOrderFetch struct {
	InnerError error
}

func (u UnsuccessOrderFetch) Error() string {
	return fmt.Sprintf("error while fetching order list: %s", u.InnerError.Error())
}
