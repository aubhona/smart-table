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

const (
	defaultCustomerLogin  = "testLogin"
	defaultCustomerTgID   = 123
	defaultCustomerChatID = 1169524813
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

func CreateDefaultCustomer() (uuid.UUID, error) {
	return CreateCustomer(
		defaultCustomerLogin,
		defaultCustomerTgID,
		defaultCustomerChatID,
	)
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
		ID:       defaultCustomerTgID,
		Username: defaultCustomerLogin,
	}
	chat := &telebot.Chat{
		ID: defaultCustomerChatID,
	}

	bot.ProcessUpdate(telebot.Update{
		Message: &telebot.Message{
			Text:   "/start",
			Chat:   chat,
			Sender: user,
		},
	})

	customerPg, err := FindCustomerByTgID(strconv.Itoa(defaultCustomerTgID))
	assert.NoError(t, err)

	assert.Equal(t, defaultCustomerLogin, customerPg.TgLogin)
	assert.Equal(t, strconv.Itoa(defaultCustomerTgID), customerPg.TgID)
	assert.Equal(t, strconv.Itoa(defaultCustomerChatID), customerPg.ChatID)
}
