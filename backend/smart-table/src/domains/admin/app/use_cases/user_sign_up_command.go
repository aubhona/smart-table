package app

type UserSingUpCommand struct {
	Login        string
	TgID         string
	TgLogin      string
	ChatID       string
	FirstName    string
	LastName     string
	PasswordHash string
}
