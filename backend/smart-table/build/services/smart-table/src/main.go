package main

import (
	"fmt"
	"log"

	"github.com/smart-table/src/servers"

	"github.com/smart-table/src/config"

	"github.com/smart-table/src/dependencies"
	"github.com/smart-table/src/domains/customer/di"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	deps := dependencies.InitDependencies(cfg)
	logger := deps.Logger

	container, err := di.BuildContainer(deps)
	if err != nil {
		logger.Fatal(err.Error())
	}

	router := servers.GetRouter(container, deps)

	err = router.Run(fmt.Sprintf(":%d", deps.Config.App.Port))
	if err != nil {
		logger.Fatal(err.Error())
	}
}
