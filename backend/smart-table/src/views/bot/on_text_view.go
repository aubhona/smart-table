package views

import (
	"fmt"
	"strconv"

	"github.com/smart-table/src/dependencies"

	"github.com/smart-table/src/domains/bot/presentation"
	app "github.com/smart-table/src/domains/customer/app/use_cases"
	appErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	"gopkg.in/telebot.v4"
)

func (b *BotUpdatesHandler) HandleOnTextUpdates(context telebot.Context) error {
	//nolint
	if context.Text() != "/start" { // TODO: wrap it to common handler to handler all commands
		ctxErr := context.Send(presentation.UnknownCommand)
		return ctxErr
	}

	deps := context.Get(utils.DependenciesName).(*dependencies.Dependencies)
	handler, err := utils.GetFromTelebotContainer[*app.CustomerRegisterCommandHandler](context)

	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))

		ctxErr := context.Send(presentation.UnknownError)
		if ctxErr != nil {
			return ctxErr
		}

		return err
	}

	result, err := handler.Handle(&app.CustomerRegisterCommand{
		TgID:    strconv.FormatInt(context.Sender().ID, 10),
		TgLogin: context.Sender().Username,
		ChatID:  strconv.FormatInt(context.Chat().ID, 10),
	})
	if err != nil && !utils.IsTheSameErrorType[*appErrors.CustomerAlreadyExist](err) {
		logging.GetLogger().Error(fmt.Sprintf("Error while handling customer register: %v", err))

		ctxErr := context.Send(presentation.UnknownError)
		if ctxErr != nil {
			return ctxErr
		}

		return err
	}

	ctxErr := context.Send(presentation.StartMessage, &telebot.SendOptions{
		ReplyMarkup: &telebot.ReplyMarkup{
			InlineKeyboard: [][]telebot.InlineButton{
				{
					telebot.InlineButton{
						Unique: result.CustomerUUID.String(),
						Text:   presentation.OpenWebApp,
						WebApp: &telebot.WebApp{
							URL: deps.Config.Bot.WebAppURL,
						},
					},
				},
			},
		},
	})

	return ctxErr
}
