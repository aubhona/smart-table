package di

import (
	"github.com/es-debug/backend-academy-2024-go-template/src/custom"
	app_services "github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/app/services"
	app "github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/app/use_cases"
	"github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/domain"
	domain_services "github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/domain/services"
	"github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/infra/pg"
	"go.uber.org/dig"
)

func BuildContainer(deps *custom.Dependencies) (*dig.Container, error) {
	container := dig.New()

	err := container.Provide(func() *custom.Dependencies {
		return deps
	})
	if err != nil {
		return nil, err
	}

	err = container.Provide(func(deps *custom.Dependencies) domain.OrderRepository {
		return pg.NewOrderRepository(deps.DbConnPool)
	})
	if err != nil {
		return nil, err
	}

	err = container.Provide(func(deps *custom.Dependencies) domain.CustomerRepository {
		return pg.NewCustomerRepository(deps.DbConnPool)
	})
	if err != nil {
		return nil, err
	}

	err = container.Provide(app_services.NewRoomCodeService)
	if err != nil {
		return nil, err
	}

	err = container.Provide(domain_services.NewUUIDGenerator)
	if err != nil {
		return nil, err
	}

	err = container.Provide(app.NewOrderCreateCommandHandler)
	if err != nil {
		return nil, err
	}

	return container, nil
}
