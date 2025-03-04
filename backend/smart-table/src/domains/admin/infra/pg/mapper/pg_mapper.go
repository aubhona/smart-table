package mapper

import (
	"encoding/json"
	defs_internal_admin_user_db "github.com/es-debug/backend-academy-2024-go-template/src/codegen/intern/admin_user_db"
	"github.com/es-debug/backend-academy-2024-go-template/src/domains/admin/domain"
	"github.com/es-debug/backend-academy-2024-go-template/src/utils"
)

func ConvertToPgUser(user utils.SharedRef[domain.User]) ([]byte, error) {
	pgUser := defs_internal_admin_user_db.PgUser{
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
	pgUser := defs_internal_admin_user_db.PgUser{}
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
