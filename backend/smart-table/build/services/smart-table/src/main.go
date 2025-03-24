package main

import (
	"github.com/smart-table/src/dependencies"
	"github.com/smart-table/src/domains/customer/di"
)

func main() {
	deps := dependencies.InitDependencies()
	logger := deps.Logger

	container, err := di.BuildContainer(deps)
	if err != nil {
		logger.Fatal(err.Error())
	}

	router := di.GetRouter(container, deps)

	err = router.Run(deps.Config.App.Port)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
