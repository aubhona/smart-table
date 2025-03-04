package domain

import (
	"time"

	domain "github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/domain/services"
	"github.com/es-debug/backend-academy-2024-go-template/src/utils"
	"github.com/google/uuid"
)

type User struct {
	uuid         uuid.UUID
	login        string
	tgId         string
	tgLogin      string
	chatId       string
	firstName    string
	lastName     string
	passwordHash string
	createdAt    time.Time
	updatedAt    time.Time
}

func NewUser(
	login string,
	tgId string,
	tgLogin string,
	chatId string,
	firstName string,
	lastName string,
	passwordHash string,
	uuidGenerator *domain.UUIDGenerator,
) utils.SharedRef[User] {
	user := User{}

	user.login = login
	user.tgId = tgId
	user.tgLogin = tgLogin
	user.chatId = chatId
	user.firstName = firstName
	user.lastName = lastName
	user.passwordHash = passwordHash
	user.createdAt = time.Now()
	user.updatedAt = time.Now()

	shardId := uuidGenerator.GetShardID()
	userUuid := uuidGenerator.GenerateShardedUUID(shardId)

	user.uuid = userUuid

	userRef, _ := utils.NewSharedRef(&user)

	return userRef
}

func RestoreUser(
	uuid uuid.UUID,
	login string,
	tgId string,
	tgLogin string,
	chatId string,
	firstName string,
	lastName string,
	passwordHash string,
	createdAt time.Time,
	updatedAt time.Time,
) utils.SharedRef[User] {
	user := User{
		uuid:         uuid,
		login:        login,
		tgId:         tgId,
		tgLogin:      tgLogin,
		chatId:       chatId,
		firstName:    firstName,
		lastName:     lastName,
		passwordHash: passwordHash,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}

	userRef, _ := utils.NewSharedRef(&user)

	return userRef
}

func (u *User) GetUUID() uuid.UUID      { return u.uuid }
func (u *User) GetLogin() string        { return u.login }
func (u *User) GetTgID() string         { return u.tgId }
func (u *User) GetTgLogin() string      { return u.tgLogin }
func (u *User) GetChatID() string       { return u.chatId }
func (u *User) GetFirstName() string    { return u.firstName }
func (u *User) GetLastName() string     { return u.lastName }
func (u *User) GetPasswordHash() string { return u.passwordHash }
func (u *User) GetCreatedAt() time.Time { return u.createdAt }
func (u *User) GetUpdatedAt() time.Time { return u.updatedAt }
