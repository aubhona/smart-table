package app

import (
	"go.uber.org/zap"

	"github.com/google/uuid"
	app "github.com/smart-table/src/domains/admin/app/services"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/domains/admin/domain"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	domainServices "github.com/smart-table/src/domains/admin/domain/services"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
)

type UserSingUpCommandHandlerResult struct {
	UserUUID uuid.UUID
	JwtToken string
}

type UserSingUpCommandHandler struct {
	userRepository domain.UserRepository
	uuidGenerator  *domainServices.UUIDGenerator
	hashService    *app.HashService
	jwtService     *app.JwtService
}

func NewUserSingUpCommandHandler(
	userRepository domain.UserRepository,
	uuidGenerator *domainServices.UUIDGenerator,
	hashService *app.HashService,
	jwtService *app.JwtService,
) *UserSingUpCommandHandler {
	return &UserSingUpCommandHandler{
		userRepository,
		uuidGenerator,
		hashService,
		jwtService,
	}
}

func (handler *UserSingUpCommandHandler) Handle(userSignUpCommand *UserSingUpCommand) (UserSingUpCommandHandlerResult, error) {
	_, err := handler.userRepository.FindUserByLoginOrTgLogin(userSignUpCommand.Login, userSignUpCommand.TgLogin)
	if err == nil {
		logging.GetLogger().Error("login or tg_login already exists",
			zap.String("login", userSignUpCommand.Login),
			zap.String("tg_login", userSignUpCommand.TgLogin))

		return UserSingUpCommandHandlerResult{}, appErrors.LoginOrTgLoginAlreadyExists{
			Login:   userSignUpCommand.Login,
			TgLogin: userSignUpCommand.TgLogin,
		}
	}

	if !utils.IsTheSameErrorType[domainErrors.UserNotFoundByLoginOrTgLogin](err) {
		logging.GetLogger().Error("error while checking login and tg_login existence",
			zap.String("login", userSignUpCommand.Login),
			zap.String("tg_login", userSignUpCommand.TgLogin),
			zap.Error(err))

		return UserSingUpCommandHandlerResult{}, err
	}

	passwordHash, err := handler.hashService.HashPassword(userSignUpCommand.Password)
	if err != nil {
		logging.GetLogger().Error("error while hashing password",
			zap.Error(err))

		return UserSingUpCommandHandlerResult{}, err
	}

	user := domain.NewUser(
		userSignUpCommand.Login,
		"userSignUpCommand.TgID",
		userSignUpCommand.TgLogin,
		"userSignUpCommand.ChatID",
		userSignUpCommand.FirstName,
		userSignUpCommand.LastName,
		passwordHash,
		handler.uuidGenerator,
	)

	tx, err := handler.userRepository.Begin()
	if err != nil {
		logging.GetLogger().Error("error while beginning transaction",
			zap.Error(err))
		return UserSingUpCommandHandlerResult{}, err
	}

	defer utils.Rollback(handler.userRepository, tx)

	err = handler.userRepository.Save(tx, user)
	if err != nil {
		logging.GetLogger().Error("error while saving user",
			zap.String("login", userSignUpCommand.Login),
			zap.Error(err))

		return UserSingUpCommandHandlerResult{}, err
	}

	err = handler.userRepository.Commit(tx)
	if err != nil {
		logging.GetLogger().Error("error while committing transaction",
			zap.Error(err))
		return UserSingUpCommandHandlerResult{}, err
	}

	userUUID := user.Get().GetUUID()

	jwtToken, err := handler.jwtService.GenerateJWT(userUUID)
	if err != nil {
		logging.GetLogger().Error("error while generating JWT token",
			zap.String("user_uuid", userUUID.String()),
			zap.Error(err))

		return UserSingUpCommandHandlerResult{}, err
	}

	return UserSingUpCommandHandlerResult{userUUID, jwtToken}, nil
}
