package mapper

import (
	"encoding/json"

	defsInternalAdminRestaurantDb "github.com/smart-table/src/codegen/intern/admin_restaurant_db"
	defsInternalAdminUserDb "github.com/smart-table/src/codegen/intern/admin_user_db"
	"github.com/smart-table/src/domains/admin/domain"
	"github.com/smart-table/src/utils"
)

func ConvertToPgUser(user utils.SharedRef[domain.User]) ([]byte, error) {
	pgUser := defsInternalAdminUserDb.PgUser{
		UUID:         user.Get().GetUUID(),
		Login:        user.Get().GetLogin(),
		TgID:         user.Get().GetTgID(),
		TgLogin:      user.Get().GetTgLogin(),
		ChatID:       user.Get().GetChatID(),
		FirstName:    user.Get().GetFirstName(),
		LastName:     user.Get().GetLastName(),
		PasswordHash: user.Get().GetPasswordHash(),
		CreatedAt:    user.Get().GetCreatedAt(),
		UpdatedAt:    user.Get().GetUpdatedAt(),
	}

	jsonBytes, err := json.Marshal(pgUser)

	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func ConvertPgUserToModel(pgResult []byte) (utils.SharedRef[domain.User], error) {
	pgUser := defsInternalAdminUserDb.PgUser{}
	err := json.Unmarshal(pgResult, &pgUser)

	if err != nil {
		return utils.SharedRef[domain.User]{}, err
	}

	return domain.RestoreUser(
		pgUser.UUID,
		pgUser.Login,
		pgUser.TgID,
		pgUser.TgLogin,
		pgUser.ChatID,
		pgUser.FirstName,
		pgUser.LastName,
		pgUser.PasswordHash,
		pgUser.CreatedAt,
		pgUser.UpdatedAt,
	), nil
}

func ConvertToPgRestaurant(user utils.SharedRef[domain.Restaurant]) ([]byte, error) {
	pgRestaurant := defsInternalAdminRestaurantDb.PgRestaurant{
		UUID:      user.Get().GetUUID(),
		OwnerUUID: user.Get().GetOwnerUUID(),
		Name:      user.Get().GetName(),
		CreatedAt: user.Get().GetCreatedAt(),
		UpdatedAt: user.Get().GetUpdatedAt(),
	}

	jsonBytes, err := json.Marshal(pgRestaurant)

	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func ConvertPgRestaurantToModel(pgResult []byte) (utils.SharedRef[domain.Restaurant], error) {
	pgRestaurant := defsInternalAdminRestaurantDb.PgRestaurant{}
	err := json.Unmarshal(pgResult, &pgRestaurant)

	if err != nil {
		return utils.SharedRef[domain.Restaurant]{}, err
	}

	return domain.RestoreRestaurant(
		pgRestaurant.UUID,
		pgRestaurant.OwnerUUID,
		pgRestaurant.Name,
		pgRestaurant.CreatedAt,
		pgRestaurant.UpdatedAt,
	), nil
}
