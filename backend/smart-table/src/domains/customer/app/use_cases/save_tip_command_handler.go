package app //nolint

import (
	appQueries "github.com/smart-table/src/domains/customer/app/queries"
	app "github.com/smart-table/src/domains/customer/app/services"
	appErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	"github.com/smart-table/src/domains/customer/domain"
	"github.com/smart-table/src/logging"
	"go.uber.org/zap"
)

type SaveTipCommandHandler struct {
	orderRepository domain.OrderRepository
	botQueryService *appQueries.BotQueryService
	tipService      *app.TipService
}

func NewSaveTipCommandHandler(
	orderRepository domain.OrderRepository,
	botQueryService *appQueries.BotQueryService,
	tipService *app.TipService,
) *SaveTipCommandHandler {
	return &SaveTipCommandHandler{
		orderRepository: orderRepository,
		botQueryService: botQueryService,
		tipService:      tipService,
	}
}

func (handler *SaveTipCommandHandler) Handle(
	command *SaveTipCommand,
) error {
	order, err := handler.orderRepository.FindOrder(command.OrderUUID)
	if err != nil {
		return err
	}

	if !order.Get().ContainsCustomer(command.CustomerUUID) {
		return appErrors.OrderAccessDenied{OrderUUID: command.OrderUUID, CustomerUUID: command.CustomerUUID}
	}

	tip := handler.tipService.CreateStringTip(order)

	customer := order.Get().GetCustomerByUUID(command.CustomerUUID)

	err = handler.botQueryService.SendMessage(customer.Value().Get().GetChatID(), tip)
	if err != nil {
		logging.GetLogger().Error("failed to send tip", zap.Error(err))
		return err
	}

	return err
}
