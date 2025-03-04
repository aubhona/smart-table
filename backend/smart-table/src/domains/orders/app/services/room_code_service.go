package app

import (
	"fmt"
	"github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/domain"
	"github.com/es-debug/backend-academy-2024-go-template/src/utils"
	"github.com/google/uuid"
	"hash"
	"hash/fnv"
	"strings"
)

type RoomCodeService struct {
	hasher hash.Hash64
}

func NewRoomCodeService() *RoomCodeService {
	return &RoomCodeService{fnv.New64a()}
}

func (r *RoomCodeService) CreateRoomCode(tableId string, hostUserUuid uuid.UUID) (string, error) {
	r.hasher.Reset()

	_, err := r.hasher.Write([]byte(tableId))
	if err != nil {
		return "", err
	}

	_, err = r.hasher.Write([]byte(hostUserUuid.String()))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", r.hasher.Sum64()%1000), nil
}

func (r *RoomCodeService) VerifyRoomCode(order utils.SharedRef[domain.Order], roomCode string) bool {
	return strings.Compare(order.Get().GetRoomCode(), roomCode) == 0
}
