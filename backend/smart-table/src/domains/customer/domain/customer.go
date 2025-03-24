package domain

import (
	"time"

	"github.com/google/uuid"
	domain "github.com/smart-table/src/domains/customer/domain/services"
	"github.com/smart-table/src/utils"
)

type Customer struct {
	uuid       uuid.UUID
	tgID       string
	tgLogin    string
	avatarLink string
	chatID     string
	createdAt  time.Time
	updatedAt  time.Time
}

func NewCustomer(
	tgID string,
	tgLogin string,
	avatarLink string,
	chatID string,
	uuidGenerator domain.UUIDGenerator,
) utils.SharedRef[Customer] {
	customer := Customer{
		uuid:       uuidGenerator.GenerateShardedUUID(uuidGenerator.GetShardID()),
		tgID:       tgID,
		tgLogin:    tgLogin,
		avatarLink: avatarLink,
		chatID:     chatID,
		createdAt:  time.Now(),
		updatedAt:  time.Now(),
	}

	customerRef, _ := utils.NewSharedRef(&customer)

	return customerRef
}

func RestoreCustomer(
	id uuid.UUID,
	tgID string,
	tgLogin string,
	avatarLink string,
	chatID string,
	createdAt time.Time,
	updatedAt time.Time,
) utils.SharedRef[Customer] {
	customer := Customer{
		uuid:       id,
		tgID:       tgID,
		tgLogin:    tgLogin,
		avatarLink: avatarLink,
		chatID:     chatID,
		createdAt:  createdAt,
		updatedAt:  updatedAt,
	}

	customerRef, _ := utils.NewSharedRef(&customer)

	return customerRef
}

func (c *Customer) GetUUID() uuid.UUID      { return c.uuid }
func (c *Customer) GetTgID() string         { return c.tgID }
func (c *Customer) GetTgLogin() string      { return c.tgLogin }
func (c *Customer) GetAvatarLink() string   { return c.avatarLink }
func (c *Customer) GetChatID() string       { return c.chatID }
func (c *Customer) GetCreatedAt() time.Time { return c.createdAt }
func (c *Customer) GetUpdatedAt() time.Time { return c.updatedAt }
