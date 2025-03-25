package di

import (
	"github.com/smart-table/src/dependencies"
	appServices "github.com/smart-table/src/domains/customer/app/services"
	app "github.com/smart-table/src/domains/customer/app/use_cases"
	"github.com/smart-table/src/domains/customer/domain"
	domainServices "github.com/smart-table/src/domains/customer/domain/services"
	"github.com/smart-table/src/domains/customer/infra/pg"
	"go.uber.org/dig"
)

func AddDeps(container *dig.Container) error {
	err := container.Provide(func(deps *dependencies.Dependencies) domain.OrderRepository {
		return pg.NewOrderRepository(deps.DBConnPool)
	})
	if err != nil {
		return err
	}

	err = container.Provide(func(deps *dependencies.Dependencies) domain.CustomerRepository {
		return pg.NewCustomerRepository(deps.DBConnPool)
	})
	if err != nil {
		return err
	}

	err = container.Provide(appServices.NewRoomCodeService)
	if err != nil {
		return err
	}

	err = container.Provide(domainServices.NewUUIDGenerator)
	if err != nil {
		return err
	}

	err = container.Provide(app.NewOrderCreateCommandHandler)
	if err != nil {
		return err
	}

	return nil
}
