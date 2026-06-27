package service

import (
	"final_project/internal/auth"
	"net/http"
)

func (s *Service) SignIn(password string) (string, error) {
	if password == "" {
		return "", NewError(http.StatusUnauthorized, NoPassword)
	}

	if s.Password != password {
		return "", NewError(http.StatusUnauthorized, InvalidPassword)
	}

	return auth.GenerateToken(password)
}
