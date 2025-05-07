package app

import (
	"go.uber.org/zap"

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

func (handler *UserSingInCommandHandler) Handle(userSignInCommand *UserSingInCommand) (UserSingInCommandHandlerResult, error) {
	user, err := handler.userRepository.FindUserByLogin(userSignInCommand.Login)
	if err != nil {
		logging.GetLogger().Error("error while finding user by login",
			zap.String("login", userSignInCommand.Login),
			zap.Error(err))

		return UserSingInCommandHandlerResult{}, err
	}

	if !handler.hashService.ComparePasswords(user.Get().GetPasswordHash(), userSignInCommand.Password) {
		logging.GetLogger().Info("incorrect password attempt",
			zap.String("login", userSignInCommand.Login))

		return UserSingInCommandHandlerResult{}, appErrors.IncorrectPassword{}
	}

	userUUID := user.Get().GetUUID()

	jwtToken, err := handler.jwtService.GenerateJWT(userUUID)
	if err != nil {
		logging.GetLogger().Error("error while generating JWT token",
			zap.String("user_uuid", userUUID.String()),
			zap.Error(err))

		return UserSingInCommandHandlerResult{}, err
	}

	return UserSingInCommandHandlerResult{userUUID, jwtToken}, nil
}
