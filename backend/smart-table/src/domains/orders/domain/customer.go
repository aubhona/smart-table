package domain

import (
	"time"

	domain "github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/domain/services"
	"github.com/es-debug/backend-academy-2024-go-template/src/utils"
	"github.com/google/uuid"
)

type Customer struct {
	uuid       uuid.UUID
	tgId       string
	tgLogin    string
	avatarLink string
	chatId     string
	createdAt  time.Time
	updatedAt  time.Time
}

func NewCustomer(tgId string, tgLogin string, avatarLink string, chatId string, uuidGenerator domain.UUIDGenerator) utils.SharedRef[Customer] {
	customer := Customer{
		uuid:       uuidGenerator.GenerateShardedUUID(uuidGenerator.GetShardID()),
		tgId:       tgId,
		tgLogin:    tgLogin,
		avatarLink: avatarLink,
		chatId:     chatId,
		createdAt:  time.Now(),
		updatedAt:  time.Now(),
	}

	customerRef, _ := utils.NewSharedRef(&customer)

	return customerRef
}

func RestoreCustomer(
	uuid uuid.UUID,
	tgId string,
	tgLogin string,
	avatarLink string,
	chatId string,
	createdAt time.Time,
	updatedAt time.Time,
) utils.SharedRef[Customer] {
	customer := Customer{
		uuid:       uuid,
		tgId:       tgId,
		tgLogin:    tgLogin,
		avatarLink: avatarLink,
		chatId:     chatId,
		createdAt:  createdAt,
		updatedAt:  updatedAt,
	}

	customerRef, _ := utils.NewSharedRef(&customer)

	return customerRef
}

func (c *Customer) GetUUID() uuid.UUID      { return c.uuid }
func (c *Customer) GetTgId() string         { return c.tgId }
func (c *Customer) GetTgLogin() string      { return c.tgLogin }
func (c *Customer) GetAvatarLink() string   { return c.avatarLink }
func (c *Customer) GetChatId() string       { return c.chatId }
func (c *Customer) GetCreatedAt() time.Time { return c.createdAt }
func (c *Customer) GetUpdatedAt() time.Time { return c.updatedAt }
