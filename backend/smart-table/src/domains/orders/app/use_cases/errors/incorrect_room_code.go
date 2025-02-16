package app_errors

import (
	"fmt"
	"github.com/es-debug/backend-academy-2024-go-template/src/utils"
)

type IncorrectRoomCodeError struct {
	RoomCode utils.Optional[string]
}

func (err IncorrectRoomCodeError) Error() string {
	return fmt.Sprintf("Incorrect room code %v", err.RoomCode.ToPointer())
}
