package app

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/google/uuid"
	app "github.com/smart-table/src/domains/admin/app/services"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/domains/admin/domain"
	domainServices "github.com/smart-table/src/domains/admin/domain/services"
	"github.com/smart-table/src/logging"
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
	ctx := context.Background()
	isExist, err := handler.userRepository.CheckLoginOrTgLoginExist(context.Background(), userSignUpCommand.Login, userSignUpCommand.TgLogin)

	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while checking login and tg_login existence: %v", err))
		return UserSingUpCommandHandlerResult{}, err
	}

	if isExist {
		return UserSingUpCommandHandlerResult{}, appErrors.LoginOrTgLoginAlreadyExists{
			Login:   userSignUpCommand.Login,
			TgLogin: userSignUpCommand.TgLogin,
		}
	}

	passwordHash, err := handler.hashService.HashPassword(userSignUpCommand.Password)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while password hashing: %v", err))
		return UserSingUpCommandHandlerResult{}, err
	}

	user := domain.NewUser(
		userSignUpCommand.Login,
		//nolint: godox, gocritic
		// TODO добавить поход в тг за валидацией логина и получения TgID и ChatID
		"userSignUpCommand.TgID",
		userSignUpCommand.TgLogin,
		"userSignUpCommand.ChatID",
		userSignUpCommand.FirstName,
		userSignUpCommand.LastName,
		passwordHash,
		handler.uuidGenerator,
	)

	tx, err := handler.userRepository.Begin(ctx)
	if err != nil {
		return UserSingUpCommandHandlerResult{}, err
	}

	defer func(userRepository domain.UserRepository, ctx context.Context, tx pgx.Tx) {
		_ = userRepository.Commit(ctx, tx)
	}(handler.userRepository, ctx, tx)

	err = handler.userRepository.Save(ctx, tx, user)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while user saving: %v", err))
		return UserSingUpCommandHandlerResult{}, err
	}

	userUUID := user.Get().GetUUID()

	jwtToken, err := handler.jwtService.GenerateJWT(userUUID)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while token generating: %v", err))
		return UserSingUpCommandHandlerResult{}, err
	}

	return UserSingUpCommandHandlerResult{userUUID, jwtToken}, nil
}
