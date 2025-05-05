package smarttable_test

import (
	"context"
	"encoding/json"
	"strconv"

	views "github.com/smart-table/src/views/bot"
	"github.com/smart-table/tests/mocks"
	"github.com/stretchr/testify/mock"
	"gopkg.in/telebot.v4"

	"testing"

	"github.com/google/uuid"
	defsInternalCustomerDb "github.com/smart-table/src/codegen/intern/customer_db"
	"github.com/stretchr/testify/assert"
)

func FindCustomerByTgID(tgID string) (defsInternalCustomerDb.PgCustomer, error) {
	customer := defsInternalCustomerDb.PgCustomer{}

	customerJSON, err := GetCustomerQueries().FetchCustomerByTgId(context.Background(), tgID)
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

	handler := views.BotUpdatesHandler{
		WebAppURL: "some_url",
	}
	mockContext := new(mocks.Context)

	mockContext.On("Text", mock.Anything, mock.Anything).Return("/start")
	mockContext.On("Sender", mock.Anything, mock.Anything).Return(user)
	mockContext.On("Chat", mock.Anything, mock.Anything).Return(chat)
	mockContext.On("Send", mock.Anything, mock.Anything).Return(nil)
	mockContext.On("Get", mock.Anything, mock.Anything).Return(GetContainer())

	err := handler.HandleOnTextUpdates(mockContext)

	if err != nil {
		return uuid.Nil, err
	}

	customerPg, err := FindCustomerByTgID(strconv.FormatInt(tgID, 10))

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
		ID: 123,
	}

	handler := views.BotUpdatesHandler{
		WebAppURL: "some_url",
	}

	mockContext := new(mocks.Context)

	mockContext.On("Text", mock.Anything).Return("/start")
	mockContext.On("Sender", mock.Anything).Return(user)
	mockContext.On("Chat", mock.Anything).Return(chat)
	mockContext.On("Get", mock.Anything).Return(GetContainer())
	mockContext.On("Send", mock.Anything, mock.Anything).Return(nil)

	err := handler.HandleOnTextUpdates(mockContext)
	assert.NoError(t, err)

	customerPg, err := FindCustomerByTgID("123")
	assert.NoError(t, err)

	assert.Equal(t, "test_login", customerPg.TgLogin)
	assert.Equal(t, "123", customerPg.TgID)
	assert.Equal(t, "123", customerPg.ChatID)
}
