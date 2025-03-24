package app

import (
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/google/uuid"
	"github.com/smart-table/src/domains/customer/domain"
	"github.com/smart-table/src/utils"
)

type RoomCodeService struct {
	hasher hash.Hash64
}

func NewRoomCodeService() *RoomCodeService {
	return &RoomCodeService{fnv.New64a()}
}

func (r *RoomCodeService) CreateRoomCode(tableID string, hostUserUUID uuid.UUID) (string, error) {
	r.hasher.Reset()

	_, err := r.hasher.Write([]byte(tableID))
	if err != nil {
		return "", err
	}

	_, err = r.hasher.Write([]byte(hostUserUUID.String()))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", r.hasher.Sum64()%1000), nil
}

func (r *RoomCodeService) VerifyRoomCode(order utils.SharedRef[domain.Order], roomCode string) bool {
	return order.Get().GetRoomCode() == roomCode
}
