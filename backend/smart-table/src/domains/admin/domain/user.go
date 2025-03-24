package domain

import (
	"time"

	"github.com/google/uuid"
	domainServices "github.com/smart-table/src/domains/admin/domain/services"
	"github.com/smart-table/src/utils"
)

type User struct {
	uuid         uuid.UUID
	login        string
	tgID         string
	tgLogin      string
	chatID       string
	firstName    string
	lastName     string
	passwordHash string
	createdAt    time.Time
	updatedAt    time.Time
}

func NewUser(
	login string,
	tgID string,
	tgLogin string,
	chatID string,
	firstName string,
	lastName string,
	passwordHash string,
	uuidGenerator *domainServices.UUIDGenerator,
) utils.SharedRef[User] {
	user := User{}

	user.login = login
	user.tgID = tgID
	user.tgLogin = tgLogin
	user.chatID = chatID
	user.firstName = firstName
	user.lastName = lastName
	user.passwordHash = passwordHash
	user.createdAt = time.Now()
	user.updatedAt = time.Now()

	shardID := uuidGenerator.GetShardID()
	userUUID := uuidGenerator.GenerateShardedUUID(shardID)

	user.uuid = userUUID

	userRef, _ := utils.NewSharedRef(&user)

	return userRef
}

func RestoreUser(
	uuid uuid.UUID,
	login string,
	tgID string,
	tgLogin string,
	chatID string,
	firstName string,
	lastName string,
	passwordHash string,
	createdAt time.Time,
	updatedAt time.Time,
) utils.SharedRef[User] {
	user := User{
		uuid:         uuid,
		login:        login,
		tgID:         tgID,
		tgLogin:      tgLogin,
		chatID:       chatID,
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
func (u *User) GettgID() string         { return u.tgID }
func (u *User) GetTgLogin() string      { return u.tgLogin }
func (u *User) GetchatID() string       { return u.chatID }
func (u *User) GetFirstName() string    { return u.firstName }
func (u *User) GetLastName() string     { return u.lastName }
func (u *User) GetPasswordHash() string { return u.passwordHash }
func (u *User) GetCreatedAt() time.Time { return u.createdAt }
func (u *User) GetUpdatedAt() time.Time { return u.updatedAt }
