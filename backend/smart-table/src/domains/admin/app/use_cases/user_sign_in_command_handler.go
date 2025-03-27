package app

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	app "github.com/smart-table/src/domains/admin/app/services"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/domains/admin/domain"
	"github.com/smart-table/src/logging"
)

type UserSingInCommandHandlerResult struct {
	UserUUID uuid.UUID
	JwtToken string
}

type UserSingInCommandHandler struct {
	userRepository domain.UserRepository
	hashService    *app.HashService
	jwtService     *app.JwtService
}

func NewUserSingInCommandHandler(
	userRepository domain.UserRepository,
	hashService *app.HashService,
	jwtService *app.JwtService,
) *UserSingInCommandHandler {
	return &UserSingInCommandHandler{
		userRepository,
		hashService,
		jwtService,
	}
}

func (handler *UserSingInCommandHandler) Handle(signInCommand *UserSingInCommand) (UserSingInCommandHandlerResult, error) {
	user, err := handler.userRepository.FindUser(context.Background(), signInCommand.Login)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while finding user by login: %v", err))
		return UserSingInCommandHandlerResult{}, err
	}

	if !handler.hashService.ComparePasswords(user.Get().GetPasswordHash(), signInCommand.Password) {
		logging.GetLogger().Info(fmt.Sprintf("Incorrect password: %v", err))
		return UserSingInCommandHandlerResult{}, appErrors.IncorrectPassword{}
	}

	userUUID := user.Get().GetUUID()

	jwtToken, err := handler.jwtService.GenerateJWT(userUUID)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while token generating: %v", err))
		return UserSingInCommandHandlerResult{}, err
	}

	return UserSingInCommandHandlerResult{userUUID, jwtToken}, nil
}
