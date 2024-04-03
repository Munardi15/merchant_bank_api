package usecase

import (
	"errors"
)

type AuthUseCase interface {
	Login(username, password string) (string, error)
	Logout(sessionID string) error
}

type authUseCase struct {
}

func NewAuthUseCase() AuthUseCase {
	return &authUseCase{}
}

func (uc *authUseCase) Login(username, password string) (string, error) {
	users := map[string]string{
		"user1": "password1",
		"user2": "password2",
	}

	// Validate username
	storedPassword, ok := users[username]
	if !ok || storedPassword != password {
		return "", errors.New("invalid username or password")
	}

	sessionID := username + "_session"

	return sessionID, nil
}

func (uc *authUseCase) Logout(sessionID string) error {
	return nil
}
