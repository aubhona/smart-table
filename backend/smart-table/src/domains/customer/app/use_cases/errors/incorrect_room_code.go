package apperrors

import (
	"fmt"

	"github.com/smart-table/src/utils"
)

type IncorrectRoomCodeError struct {
	RoomCode utils.Optional[string]
}

func (err IncorrectRoomCodeError) Error() string {
	return fmt.Sprintf("Incorrect room code %v", err.RoomCode.ToPointer())
}
