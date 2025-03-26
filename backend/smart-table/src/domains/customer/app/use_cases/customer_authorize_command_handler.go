package app

import (
	"context"

	"github.com/google/uuid"
	"github.com/smart-table/src/domains/customer/domain"
	domainServices "github.com/smart-table/src/domains/customer/domain/services"
)

type CustomerAuthorizeCommandHandlerResult struct {
	CustomerUUID uuid.UUID
}

type CustomerAuthorizeCommandHandler struct {
	customerRepository domain.CustomerRepository
	uuidGenerator      *domainServices.UUIDGenerator
}

func NewCustomerAuthorizeCommandHandler(
	customerRepository domain.CustomerRepository,
	uuidGenerator *domainServices.UUIDGenerator,
) *CustomerAuthorizeCommandHandler {
	return &CustomerAuthorizeCommandHandler{
		customerRepository,
		uuidGenerator,
	}
}

func (handler *CustomerAuthorizeCommandHandler) Handle(
	createCommand *CustomerAuthorizeCommand) (CustomerAuthorizeCommandHandlerResult, error) {
	ctx := context.Background()
	customer, err := handler.customerRepository.FindCustomerByTgID(ctx, createCommand.TgID)

	if err != nil {
		return CustomerAuthorizeCommandHandlerResult{}, err
	}

	if customer.Get().GetTgLogin() != createCommand.TgLogin || customer.Get().GetChatID() != createCommand.ChatID {
		customer.Get().SetTgLogin(createCommand.TgLogin)
		customer.Get().SetChatID(createCommand.ChatID)

		err := handler.customerRepository.SaveAndUpdate(ctx, customer)
		if err != nil {
			return CustomerAuthorizeCommandHandlerResult{}, err
		}
	}

	return CustomerAuthorizeCommandHandlerResult{CustomerUUID: customer.Get().GetUUID()}, err
}
