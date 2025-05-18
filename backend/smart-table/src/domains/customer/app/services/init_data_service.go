package app

import (
	"time"

	"github.com/smart-table/src/dependencies"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

type InitDataService struct {
	botToken        string
	tokenExpiration time.Duration
}

func NewInitDataService(deps *dependencies.Dependencies) *InitDataService {
	return &InitDataService{
		botToken:        deps.Config.Bot.Token,
		tokenExpiration: deps.Config.App.Customer.InitData.Expiration,
	}
}

func (ids *InitDataService) VerifyInitData(initData string) bool {
	return initdata.Validate(initData, ids.botToken, ids.tokenExpiration) == nil
}
