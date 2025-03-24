package app

import (
	"context"

	app_errors "github.com/es-debug/backend-academy-2024-go-template/src/domains/admin/app/use_cases/errors"
	domainServices "github.com/smart-table/src/domains/customer/domain/services"
	"github.com/es-debug/backend-academy-2024-go-template/src/domains/admin/domain"
	"github.com/google/uuid"
)

type UserSingUpCommandHandlerResult struct {
	UserUUID uuid.UUID
}

type UserSingUpCommandHandler struct {
	userRepository domain.UserRepository
	uuidGenerator  *domainServices.UUIDGenerator
}

func NewOrderCreateCommandHandler(
	userRepository domain.UserRepository,

	uuidGenerator *domainServices.UUIDGenerator,
) *UserSingUpCommandHandler {
	return &UserSingUpCommandHandler{
		userRepository,
		uuidGenerator,
	}
}

func (handler *UserSingUpCommandHandler) Handle(signUpCommand *UserSingUpCommand) (UserSingUpCommandHandlerResult, error) {
	isExist, err := handler.userRepository.CheckLoginOrTgLoginExist(context.Background(), signUpCommand.Login, signUpCommand.TgLogin)
	if err != nil {
		return UserSingUpCommandHandlerResult{}, err
	}

	if isExist {
		return UserSingUpCommandHandlerResult{}, app_errors.LoginOrTgLoginAlreadyExists{Login: signUpCommand.Login, TgLogin: signUpCommand.TgLogin}
	}

	user := domain.NewUser(signUpCommand.Login,
		signUpCommand.TgId,
		signUpCommand.TgLogin,
		signUpCommand.ChatId,
		signUpCommand.FirstName,
		signUpCommand.LastName,
		signUpCommand.PasswordHash,
		handler.uuidGenerator)

	return UserSingUpCommandHandlerResult{user.Get().GetUUID()}, nil
}
