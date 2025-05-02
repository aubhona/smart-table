package errors

import (
	"fmt"
	"strconv"
)

type InvalidTableCount struct {
	TableCount int
}

func (e InvalidTableCount) Error() string {
	return fmt.Sprintf("Invalid table count: %s", strconv.Itoa(e.TableCount))
}
