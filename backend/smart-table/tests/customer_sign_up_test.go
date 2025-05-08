package smarttable_test

import (
	"encoding/json"
	"strconv"

	"gopkg.in/telebot.v4"

	"testing"

	"github.com/google/uuid"
	defsInternalCustomerDb "github.com/smart-table/src/codegen/intern/customer_db"
	"github.com/stretchr/testify/assert"
)

func FindCustomerByTgID(tgID string) (defsInternalCustomerDb.PgCustomer, error) {
	customer := defsInternalCustomerDb.PgCustomer{}

	customerJSON, err := GetCustomerQueries().FetchCustomerByTgId(GetCtx(), tgID)
	if err != nil {
		return customer, err
	}

	err = json.Unmarshal(customerJSON, &customer)

	return customer, err
}

func CreateCustomer(tgLogin string, tgID, chatID int64) (uuid.UUID, error) {
	user := &telebot.User{
		ID:       tgID,
		Username: tgLogin,
	}
	chat := &telebot.Chat{
		ID: chatID,
	}

	bot.ProcessUpdate(telebot.Update{
		Message: &telebot.Message{
			Text:   "/start",
			Chat:   chat,
			Sender: user,
		},
	})

	customerPg, err := FindCustomerByTgID(strconv.FormatInt(tgID, 10))
	if err != nil {
		return uuid.Nil, err
	}

	return customerPg.UUID, err
}

func TestCustomerRegisterHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	user := &telebot.User{
		ID:       123,
		Username: "test_login",
	}
	chat := &telebot.Chat{
		ID: 1169524813,
	}

	bot.ProcessUpdate(telebot.Update{
		Message: &telebot.Message{
			Text:   "/start",
			Chat:   chat,
			Sender: user,
		},
	})

	customerPg, err := FindCustomerByTgID("123")
	assert.NoError(t, err)

	assert.Equal(t, "test_login", customerPg.TgLogin)
	assert.Equal(t, "123", customerPg.TgID)
	assert.Equal(t, "1169524813", customerPg.ChatID)
}
