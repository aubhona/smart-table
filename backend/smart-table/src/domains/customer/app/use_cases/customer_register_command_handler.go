package app

import (
	"github.com/smart-table/src/utils"

	"github.com/google/uuid"
	apperrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	"github.com/smart-table/src/domains/customer/domain"
	domainerrors "github.com/smart-table/src/domains/customer/domain/errors"
	domainServices "github.com/smart-table/src/domains/customer/domain/services"
)

type CustomerRegisterCommandHandlerResult struct {
	CustomerUUID uuid.UUID
}

type CustomerRegisterCommandHandler struct {
	customerRepository domain.CustomerRepository
	uuidGenerator      *domainServices.UUIDGenerator
}

func NewCustomerRegisterCommandHandler(
	customerRepository domain.CustomerRepository,
	uuidGenerator *domainServices.UUIDGenerator,
) *CustomerRegisterCommandHandler {
	return &CustomerRegisterCommandHandler{
		customerRepository,
		uuidGenerator,
	}
}

func (handler *CustomerRegisterCommandHandler) Handle(
	createCommand *CustomerRegisterCommand) (CustomerRegisterCommandHandlerResult, error) {
	customer, err := handler.customerRepository.FindCustomerByTgID(createCommand.TgID)

	if err == nil {
		return CustomerRegisterCommandHandlerResult{
			CustomerUUID: customer.Get().GetUUID(),
		}, &apperrors.CustomerAlreadyExist{TgID: createCommand.TgID}
	}

	if !utils.IsTheSameErrorType[domainerrors.CustomerNotFoundByTgID](err) {
		return CustomerRegisterCommandHandlerResult{}, err
	}

	customer = domain.NewCustomer(
		createCommand.TgID, createCommand.TgLogin, "TODO", createCommand.ChatID, *handler.uuidGenerator)

	tx, err := handler.customerRepository.Begin()
	if err != nil {
		return CustomerRegisterCommandHandlerResult{}, err
	}

	defer utils.Rollback(handler.customerRepository, tx)

	err = handler.customerRepository.SaveAndUpdate(tx, customer)

	if err != nil {
		return CustomerRegisterCommandHandlerResult{}, err
	}

	err = handler.customerRepository.Commit(tx)
	if err != nil {
		return CustomerRegisterCommandHandlerResult{}, err
	}

	return CustomerRegisterCommandHandlerResult{CustomerUUID: customer.Get().GetUUID()}, err
}
