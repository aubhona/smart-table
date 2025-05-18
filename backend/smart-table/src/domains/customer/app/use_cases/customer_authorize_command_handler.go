package app

import (
	"github.com/google/uuid"
	appServices "github.com/smart-table/src/domains/customer/app/services"
	appErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	"github.com/smart-table/src/domains/customer/domain"
	domainServices "github.com/smart-table/src/domains/customer/domain/services"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	"go.uber.org/zap"
)

type CustomerAuthorizeCommandHandlerResult struct {
	CustomerUUID uuid.UUID
	JwtToken     string
}

type CustomerAuthorizeCommandHandler struct {
	customerRepository domain.CustomerRepository
	uuidGenerator      *domainServices.UUIDGenerator
	jwtService         *appServices.JwtService
	initDataService    *appServices.InitDataService
}

func NewCustomerAuthorizeCommandHandler(
	customerRepository domain.CustomerRepository,
	uuidGenerator *domainServices.UUIDGenerator,
	jwtService *appServices.JwtService,
	initDataService *appServices.InitDataService,
) *CustomerAuthorizeCommandHandler {
	return &CustomerAuthorizeCommandHandler{
		customerRepository,
		uuidGenerator,
		jwtService,
		initDataService,
	}
}

func (handler *CustomerAuthorizeCommandHandler) Handle(
	createCommand *CustomerAuthorizeCommand) (CustomerAuthorizeCommandHandlerResult, error) {
	if !handler.initDataService.VerifyInitData(createCommand.InitData) {
		return CustomerAuthorizeCommandHandlerResult{}, appErrors.IncorrectInitDataError{
			InitData: createCommand.InitData,
		}
	}

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

	jwtToken, err := handler.jwtService.GenerateJWT(customer.Get().GetUUID())
	if err != nil {
		logging.GetLogger().Error("error while generating JWT token",
			zap.String("user_uuid", customer.Get().GetUUID().String()),
			zap.Error(err))

		return CustomerAuthorizeCommandHandlerResult{}, err
	}

	return CustomerAuthorizeCommandHandlerResult{CustomerUUID: customer.Get().GetUUID(), JwtToken: jwtToken}, err
}
