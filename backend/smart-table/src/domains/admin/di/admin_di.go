package di

import (
	"github.com/smart-table/src/dependencies"
	appQueries "github.com/smart-table/src/domains/admin/app/queries"
	appServices "github.com/smart-table/src/domains/admin/app/services"
	app "github.com/smart-table/src/domains/admin/app/use_cases"
	"github.com/smart-table/src/domains/admin/domain"
	domainServices "github.com/smart-table/src/domains/admin/domain/services"
	"github.com/smart-table/src/domains/admin/infra/pg"
	infraQueries "github.com/smart-table/src/domains/admin/infra/queries"
	"go.uber.org/dig"
)

func addRepositories(container *dig.Container) error {
	err := container.Provide(func(deps *dependencies.Dependencies) domain.UserRepository {
		return pg.NewUserRepository(deps.DBConnPool)
	})
	if err != nil {
		return err
	}

	err = container.Provide(func(deps *dependencies.Dependencies) domain.RestaurantRepository {
		return pg.NewRestaurantRepository(deps.DBConnPool)
	})
	if err != nil {
		return err
	}

	err = container.Provide(func(deps *dependencies.Dependencies) domain.PlaceRepository {
		return pg.NewPlaceRepository(deps.DBConnPool)
	})
	if err != nil {
		return err
	}

	return nil
}

func addServices(container *dig.Container) error {
	err := container.Provide(
		func(
			stCustomerQueryServiceImpl *infraQueries.SmartTableCustomerQueryServiceImpl,
		) appQueries.SmartTableCustomerQueryService {
			return stCustomerQueryServiceImpl
		})
	if err != nil {
		return err
	}

	err = container.Provide(infraQueries.NewSmartTableQueryServiceImpl)
	if err != nil {
		return err
	}

	err = container.Provide(appQueries.NewS3QueryService)
	if err != nil {
		return err
	}

	err = container.Provide(appServices.NewHashService)
	if err != nil {
		return err
	}

	err = container.Provide(appServices.NewJwtService)
	if err != nil {
		return err
	}

	err = container.Provide(appServices.NewPlaceTableService)
	if err != nil {
		return err
	}

	return nil
}

func addHandlers(container *dig.Container) error { //nolint
	err := container.Provide(app.NewUserSingUpCommandHandler)
	if err != nil {
		return err
	}

	err = container.Provide(app.NewUserSingInCommandHandler)
	if err != nil {
		return err
	}

	err = container.Provide(app.NewRestaurantCreateCommandHandler)
	if err != nil {
		return err
	}

	err = container.Provide(app.NewRestaurantListCommandHandler)
	if err != nil {
		return err
	}

	err = container.Provide(app.NewPlaceCreateCommandHandler)
	if err != nil {
		return err
	}

	err = container.Provide(app.NewPlaceListCommandHandler)
	if err != nil {
		return err
	}

	err = container.Provide(app.NewDishCreateCommandHandler)
	if err != nil {
		return err
	}

	err = container.Provide(app.NewEmployeeAddCommandHandler)
	if err != nil {
		return err
	}

	err = container.Provide(app.NewDishListCommandHandler)
	if err != nil {
		return err
	}

	err = container.Provide(app.NewEmployeeListCommandHandler)
	if err != nil {
		return err
	}

	err = container.Provide(app.NewMenuDishCreateCommandHandler)
	if err != nil {
		return err
	}

	err = container.Provide(app.NewMenuDishListCommandHandler)
	if err != nil {
		return err
	}

	err = container.Provide(app.NewTableDeepLinksListCommandHandler)
	if err != nil {
		return err
	}

	err = container.Provide(app.NewTableIDValidateCommandHandler)
	if err != nil {
		return err
	}

	err = container.Provide(app.NewMenuDishStateCommandHandler)
	if err != nil {
		return err
	}

	err = container.Provide(app.NewOrderListCommandHandler)
	if err != nil {
		return err
	}

	err = container.Provide(app.NewOrderInfoCommandHandler)
	if err != nil {
		return err
	}

	return nil
}

func AddDeps(container *dig.Container) error {
	err := container.Provide(domainServices.NewUUIDGenerator)
	if err != nil {
		return err
	}

	err = addRepositories(container)
	if err != nil {
		return err
	}

	err = addServices(container)
	if err != nil {
		return err
	}

	err = addHandlers(container)
	if err != nil {
		return err
	}

	return nil
}
