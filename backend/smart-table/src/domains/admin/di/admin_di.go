package di

import (
	"github.com/smart-table/src/dependencies"
	appServices "github.com/smart-table/src/domains/admin/app/services"
	app "github.com/smart-table/src/domains/admin/app/use_cases"
	"github.com/smart-table/src/domains/admin/domain"
	domainServices "github.com/smart-table/src/domains/admin/domain/services"
	"github.com/smart-table/src/domains/admin/infra/pg"
	"go.uber.org/dig"
)

func AddDeps(container *dig.Container) error {
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

	err = container.Provide(appServices.NewHashService)
	if err != nil {
		return err
	}

	err = container.Provide(appServices.NewJwtService)
	if err != nil {
		return err
	}

	err = container.Provide(domainServices.NewUUIDGenerator)
	if err != nil {
		return err
	}

	err = container.Provide(app.NewUserSingUpCommandHandler)
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

	return nil
}
