package refresh_token

import (
	"fmt"
	"givebox/domain/identity"
	"givebox/domain/shared"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const BcryptCost = 10

type RefreshToken struct {
	ID        identity.ID
	UserID    identity.ID
	Token     string
	ExpiresAt time.Time
	shared.Timestamp
}

func HashToken(token string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(token), BcryptCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash token: %w", err)
	}

	return string(bytes), err
}
