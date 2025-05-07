package app

import (
	"github.com/google/uuid"
	"github.com/smart-table/src/domains/customer/domain"
	domainServices "github.com/smart-table/src/domains/customer/domain/services"
	"github.com/smart-table/src/utils"
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
	customer, err := handler.customerRepository.FindCustomerByTgID(createCommand.TgID)

	if err != nil {
		return CustomerAuthorizeCommandHandlerResult{}, err
	}

	if customer.Get().GetTgLogin() != createCommand.TgLogin {
		customer.Get().SetTgLogin(createCommand.TgLogin)

		tx, err := handler.customerRepository.Begin()
		if err != nil {
			return CustomerAuthorizeCommandHandlerResult{}, err
		}

		defer utils.Rollback(handler.customerRepository, tx)

		err = handler.customerRepository.SaveAndUpdate(tx, customer)
		if err != nil {
			return CustomerAuthorizeCommandHandlerResult{}, err
		}

		err = handler.customerRepository.Commit(tx)
		if err != nil {
			return CustomerAuthorizeCommandHandlerResult{}, err
		}
	}

	return CustomerAuthorizeCommandHandlerResult{CustomerUUID: customer.Get().GetUUID()}, err
}
