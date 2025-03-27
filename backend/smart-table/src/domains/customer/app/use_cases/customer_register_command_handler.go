package app

import (
	"context"
	"github.com/jackc/pgx/v5"

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
	ctx := context.Background()
	_, err := handler.customerRepository.FindCustomerByTgID(context.Background(), createCommand.TgID)
	if err == nil {
		return CustomerRegisterCommandHandlerResult{}, &apperrors.CustomerAlreadyExist{TgID: createCommand.TgID}
	}

	if !utils.IsTheSameErrorType[domainerrors.CustomerNotFoundByTgID](err) {
		return CustomerRegisterCommandHandlerResult{}, err
	}

	customer := domain.NewCustomer(
		createCommand.TgID, createCommand.TgLogin, "TODO", createCommand.ChatID, *handler.uuidGenerator)

	tx, err := handler.customerRepository.Begin(ctx)
	if err != nil {
		return CustomerRegisterCommandHandlerResult{}, err
	}

	defer func(customerRepository domain.CustomerRepository, ctx context.Context, tx pgx.Tx) {
		_ = customerRepository.Commit(ctx, tx)
	}(handler.customerRepository, ctx, tx)

	err = handler.customerRepository.SaveAndUpdate(ctx, tx, customer)

	if err != nil {
		return CustomerRegisterCommandHandlerResult{}, err
	}

	return CustomerRegisterCommandHandlerResult{CustomerUUID: customer.Get().GetUUID()}, err
}
