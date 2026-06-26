package service

import "final_project/internal/auth"

func (s *Service) SignIn(password string) (string, error) {
	if password == "" {
		return "", NewError(401, "password not specified")
	}

	if s.Password != password {
		return "", NewError(401, "incorrect password")
	}

	return auth.GenerateToken(password)
}
