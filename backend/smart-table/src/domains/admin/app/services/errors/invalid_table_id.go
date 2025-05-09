package errors

import "fmt"

type InvalidTableID struct {
	TableID string
}

func (c InvalidTableID) Error() string {
	return fmt.Sprintf("invalid table id %s", c.TableID)
}
