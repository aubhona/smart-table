package servers

import (
	"github.com/smart-table/src/dependencies"
	"github.com/smart-table/src/utils"
	views "github.com/smart-table/src/views/bot"
	"go.uber.org/dig"
	"gopkg.in/telebot.v4"
	"gopkg.in/telebot.v4/middleware"
)

func NewBot(container *dig.Container, deps *dependencies.Dependencies) (*telebot.Bot, error) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:     deps.Config.Bot.Token,
		Poller:    &telebot.LongPoller{Timeout: deps.Config.Bot.PollerTimeout},
		ParseMode: telebot.ModeMarkdown,
	})

	if err != nil {
		return nil, err
	}

	bot.Use(middleware.Recover())
	bot.Use(botLogger(deps.Logger))
	bot.Use(func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			c.Set(utils.DiContainerName, container)
			c.Set(utils.DependenciesName, deps)

			return next(c)
		}
	})

	botHandler := views.BotUpdatesHandler{}

	bot.Handle(telebot.OnText, botHandler.HandleOnTextUpdates)

	return bot, nil
}
