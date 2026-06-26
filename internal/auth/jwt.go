package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "super-secret-key"

func GenerateToken(password string) (string, error) {
	sum := sha256.Sum256([]byte(password))
	hash := hex.EncodeToString(sum[:])

	claims := jwt.MapClaims{
		"hash": hash,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func ValidateToken(tokenString string, password string) error {
	sum := sha256.Sum256([]byte(password))
	hash := hex.EncodeToString(sum[:])

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("explicitly expected HS256, got: %v", token.Method.Alg())
		}
		return []byte(secret), nil
	},
	)
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("token is invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("unexpected type")
	}

	tokenHash, ok := claims["hash"].(string)
	if !ok {
		return errors.New("hash claim missing")
	}
	if tokenHash != hash {
		return errors.New("token hash does not match password hash")
	}

	return nil
}
