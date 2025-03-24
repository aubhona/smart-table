// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type SmartTableAdminUser struct {
	Uuid         uuid.UUID
	Login        string
	TgID         string
	TgLogin      string
	ChatID       string
	FirstName    string
	LastName     string
	PasswordHash string
	CreatedAt    pgtype.Timestamptz
	UpdatedAt    pgtype.Timestamptz
}

type SmartTableCustomerCustomer struct {
	Uuid       uuid.UUID
	TgID       string
	TgLogin    string
	AvatarLink string
	ChatID     string
	CreatedAt  pgtype.Timestamptz
	UpdatedAt  pgtype.Timestamptz
}

type SmartTableCustomerDish struct {
	Uuid        uuid.UUID
	Name        string
	Description string
	Weight      int32
	PictureLink string
	RestUuid    uuid.UUID
	Category    string
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

type SmartTableCustomerItem struct {
	Uuid         uuid.UUID
	OrderUuid    uuid.UUID
	Comment      pgtype.Text
	Status       string
	Resolution   pgtype.Text
	Name         string
	Description  string
	PictureLink  string
	Weight       int32
	Category     string
	Price        pgtype.Numeric
	CustomerUuid uuid.UUID
	DishUuid     uuid.UUID
	IsDraft      bool
	CreatedAt    pgtype.Timestamptz
	UpdatedAt    pgtype.Timestamptz
}

type SmartTableCustomerMenuDish struct {
	Uuid      uuid.UUID
	DishUuid  uuid.UUID
	PlaceUuid uuid.UUID
	Price     pgtype.Numeric
	Exist     bool
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type SmartTableCustomerOrder struct {
	Uuid          uuid.UUID
	RoomCode      string
	TableID       string
	CustomersUuid []uuid.UUID
	HostUserUuid  uuid.UUID
	Status        string
	Resolution    pgtype.Text
	CreatedAt     pgtype.Timestamptz
	UpdatedAt     pgtype.Timestamptz
}

type SmartTableCustomerPlace struct {
	Uuid        uuid.UUID
	RestUuid    uuid.UUID
	Address     string
	OpeningTime pgtype.Time
	ClosingTime pgtype.Time
	TableCount  int32
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

type SmartTableCustomerRestaurant struct {
	Uuid      uuid.UUID
	Name      string
	OwnerUuid uuid.UUID
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type SmartTableCustomerStaff struct {
	UserUuid  uuid.UUID
	PlaceUuid uuid.UUID
	Role      string
	Active    bool
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type SmartTableCustomerUser struct {
	Uuid         uuid.UUID
	AvatarLink   string
	Login        string
	TgLogin      string
	Name         string
	Phone        string
	PasswordHash string
	ChatID       string
	CreatedAt    pgtype.Timestamptz
	UpdatedAt    pgtype.Timestamptz
}
