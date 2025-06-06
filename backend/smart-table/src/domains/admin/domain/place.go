package domain

import (
	"fmt"
	"time"

	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	domainServices "github.com/smart-table/src/domains/admin/domain/services"
	"github.com/smart-table/src/utils"
	"golang.org/x/exp/slices"
)

type Place struct {
	uuid        uuid.UUID
	restaurant  utils.SharedRef[Restaurant]
	employees   []utils.SharedRef[Employee]
	menuDishes  []utils.SharedRef[MenuDish]
	address     string
	tableCount  int
	openingTime time.Time
	closingTime time.Time
	createdAt   time.Time
	updatedAt   time.Time

	deletedMenuDishUUIDs []uuid.UUID
}

func NewPlace(
	restaurant utils.SharedRef[Restaurant],
	address string,
	tableCount int,
	openingTime,
	closingTime time.Time,
	uuidGenerator *domainServices.UUIDGenerator,
) (utils.SharedRef[Place], error) {
	if tableCount <= 0 {
		return utils.SharedRef[Place]{}, domainErrors.InvalidTableCount{TableCount: tableCount}
	}

	place := Place{
		restaurant:  restaurant,
		employees:   make([]utils.SharedRef[Employee], 0),
		menuDishes:  make([]utils.SharedRef[MenuDish], 0),
		address:     address,
		tableCount:  tableCount,
		openingTime: openingTime,
		closingTime: closingTime,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),

		deletedMenuDishUUIDs: make([]uuid.UUID, 0),
	}

	shardID := uuidGenerator.GetShardID()
	place.uuid = uuidGenerator.GenerateShardedUUID(shardID)

	placeRef, _ := utils.NewSharedRef(&place)

	return placeRef, nil
}

func RestorePlace(
	id uuid.UUID,
	restaurant utils.SharedRef[Restaurant],
	employees []utils.SharedRef[Employee],
	menuDishes []utils.SharedRef[MenuDish],
	address string,
	tableCount int,
	openingTime,
	closingTime,
	createdAt,
	updatedAt time.Time,
) utils.SharedRef[Place] {
	place := Place{
		uuid:        id,
		restaurant:  restaurant,
		employees:   employees,
		menuDishes:  menuDishes,
		address:     address,
		tableCount:  tableCount,
		openingTime: openingTime,
		closingTime: closingTime,
		createdAt:   createdAt,
		updatedAt:   updatedAt,

		deletedMenuDishUUIDs: make([]uuid.UUID, 0),
	}

	placeRef, _ := utils.NewSharedRef(&place)

	return placeRef
}

func (p *Place) AddEmployee(
	user utils.SharedRef[User],
	role string,
) {
	employee := NewEmployee(user, p.uuid, role)
	p.employees = append(p.employees, employee)
}

func (p *Place) AddMenuDish(
	dishUUID uuid.UUID,
	price decimal.Decimal,
	exist bool,
	uuidGenerator *domainServices.UUIDGenerator,
) (utils.SharedRef[MenuDish], error) {
	dish, ok := lo.Find(p.restaurant.Get().dishes, func(dish utils.SharedRef[Dish]) bool {
		return dish.Get().uuid == dishUUID
	})
	if !ok {
		return utils.SharedRef[MenuDish]{}, domainErrors.DishNotFound{UUID: dishUUID}
	}

	menuDish := NewMenuDish(p.uuid, dish, price, exist, uuidGenerator)
	p.menuDishes = append(p.menuDishes, menuDish)

	return menuDish, nil
}

func (p *Place) DeleteMenuDish(menuDishUUID uuid.UUID) error {
	for i, menuDish := range p.menuDishes {
		if menuDish.Get().GetUUID() != menuDishUUID {
			continue
		}

		newMenuDishList := make([]utils.SharedRef[MenuDish], 0, len(p.menuDishes)-1)
		newMenuDishList = append(newMenuDishList, p.menuDishes[:i]...)
		newMenuDishList = append(newMenuDishList, p.menuDishes[i+1:]...)

		p.menuDishes = newMenuDishList
		p.deletedMenuDishUUIDs = append(p.deletedMenuDishUUIDs, menuDishUUID)

		return nil
	}

	return domainErrors.MenuDishNotFound{UUID: menuDishUUID}
}

func (p *Place) ContainsEmployee(employeeUserUUID uuid.UUID) bool {
	return employeeUserUUID == p.GetRestaurant().Get().GetOwner().Get().GetUUID() ||
		slices.ContainsFunc(p.GetEmployees(), func(employee utils.SharedRef[Employee]) bool {
			return employee.Get().GetUser().Get().GetUUID() == employeeUserUUID
		})
}

func (p *Place) ValidateTableNumber(tableNumber int) bool {
	return 0 < tableNumber && tableNumber <= p.GetTableCount()
}

func (p *Place) GetTableIDs() []string {
	tableIDs := make([]string, 0, p.GetTableCount())

	for i := 1; i <= p.GetTableCount(); i++ {
		tableIDs = append(tableIDs, fmt.Sprintf("%s_%d", p.GetUUID(), i))
	}

	return tableIDs
}

func (p *Place) GetUUID() uuid.UUID                         { return p.uuid }
func (p *Place) GetRestaurant() utils.SharedRef[Restaurant] { return p.restaurant }
func (p *Place) GetEmployees() []utils.SharedRef[Employee]  { return p.employees }
func (p *Place) GetMenuDishes() []utils.SharedRef[MenuDish] { return p.menuDishes }

func (p *Place) GetMenuDishByUUID(menuDishUUID uuid.UUID) utils.Optional[utils.SharedRef[MenuDish]] {
	menuDish, found := lo.Find(p.menuDishes, func(d utils.SharedRef[MenuDish]) bool {
		return d.Get().uuid == menuDishUUID
	})
	if !found {
		return utils.EmptyOptional[utils.SharedRef[MenuDish]]()
	}

	return utils.NewOptional(menuDish)
}

func (p *Place) GetAddress() string        { return p.address }
func (p *Place) GetTableCount() int        { return p.tableCount }
func (p *Place) GetOpeningTime() time.Time { return p.openingTime }
func (p *Place) GetClosingTime() time.Time { return p.closingTime }
func (p *Place) GetCreatedAt() time.Time   { return p.createdAt }
func (p *Place) GetUpdatedAt() time.Time   { return p.updatedAt }

func (p *Place) GetDeletedMenuDishUUIDs() []uuid.UUID { return p.deletedMenuDishUUIDs }
