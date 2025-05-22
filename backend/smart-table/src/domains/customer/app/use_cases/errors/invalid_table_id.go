package apperrors

import "fmt"

type InvalidTableID struct {
	TableID string
}

func (e InvalidTableID) Error() string {
	return fmt.Sprintf("invalid table_id='%s'", e.TableID)
}
