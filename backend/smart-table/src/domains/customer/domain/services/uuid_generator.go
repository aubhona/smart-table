package domain

import (
	"github.com/google/uuid"
)

type UUIDGenerator struct{}

func NewUUIDGenerator() *UUIDGenerator {
	return &UUIDGenerator{}
}

func (generator *UUIDGenerator) Generate() uuid.UUID {
	return uuid.New()
}

func (generator *UUIDGenerator) GenerateShardedUUID(shardID uint8) uuid.UUID {
	if shardID > 15 {
		panic("Shard ID must be between 0 and 15")
	}

	u := uuid.New()
	u[0] = (shardID << 4) | (u[0] & 0x0F)

	return u
}

func (generator *UUIDGenerator) ExtractShardID(u uuid.UUID) uint8 {
	return u[0] >> 4
}

func (generator *UUIDGenerator) GetShardID() uint8 {
	return 1
}
