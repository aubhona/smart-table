package errors

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type InvalidTableNumber struct {
	TableNumber int
	PlaceUUID   uuid.UUID
}

func (e InvalidTableNumber) Error() string {
	return fmt.Sprintf("invalid table number '%s' in place with uuid '%s'", strconv.Itoa(e.TableNumber), e.PlaceUUID)
}
