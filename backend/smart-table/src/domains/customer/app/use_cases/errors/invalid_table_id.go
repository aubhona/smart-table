package apperrors

import "fmt"

type InvalidTableID struct {
	TableID string
}

func (e InvalidTableID) Error() string {
	return fmt.Sprintf("invalid table id '%s'", e.TableID)
}
