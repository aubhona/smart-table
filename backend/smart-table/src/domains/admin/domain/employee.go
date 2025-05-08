package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/smart-table/src/utils"
)

type Employee struct {
	user      utils.SharedRef[User]
	placeUUID uuid.UUID
	role      string
	active    bool
	createdAt time.Time
	updatedAt time.Time
}

func NewEmployee(
	user utils.SharedRef[User],
	placeUUID uuid.UUID,
	role string,
) utils.SharedRef[Employee] {
	employee := Employee{
		user:      user,
		placeUUID: placeUUID,
		role:      role,
		active:    true,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	employeeRef, _ := utils.NewSharedRef(&employee)

	return employeeRef
}

func RestoreEmployee(
	user utils.SharedRef[User],
	placeUUID uuid.UUID,
	role string,
	active bool,
	createdAt time.Time,
	updatedAt time.Time,
) utils.SharedRef[Employee] {
	employee := Employee{
		user:      user,
		placeUUID: placeUUID,
		role:      role,
		active:    active,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}

	employeeRef, _ := utils.NewSharedRef(&employee)

	return employeeRef
}

func (e *Employee) GetUser() utils.SharedRef[User] { return e.user }
func (e *Employee) GetPlaceUUID() uuid.UUID        { return e.placeUUID }
func (e *Employee) GetRole() string                { return e.role }
func (e *Employee) GetActive() bool                { return e.active }
func (e *Employee) GetCreatedAt() time.Time        { return e.createdAt }
func (e *Employee) GetUpdatedAt() time.Time        { return e.updatedAt }
