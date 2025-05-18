package app //nolint

import (
	"sync"

	appQueries "github.com/smart-table/src/domains/customer/app/queries"
	app "github.com/smart-table/src/domains/customer/app/services"
	appErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	"github.com/smart-table/src/domains/customer/domain"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
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

	waitGroup := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	for _, customer := range order.Get().GetCustomers() {
		waitGroup.Add(1)

		go func(customer utils.SharedRef[domain.Customer]) {
			defer waitGroup.Done()
			mutex.Lock()
			defer mutex.Unlock()

			sendErr := handler.botQueryService.SendMessage(customer.Get().GetChatID(), tip)
			if sendErr != nil {
				logging.GetLogger().Error("failed to send tip", zap.Error(sendErr))
			}
		}(customer)
	}

	waitGroup.Wait()

	return nil
}
