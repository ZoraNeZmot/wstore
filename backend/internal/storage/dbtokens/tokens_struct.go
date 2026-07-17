package dbtokens

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	Token     string    `json:"token"`
	UserID    uuid.UUID `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NewToken(token string, userID uuid.UUID, expiresAt time.Time) *Token {
	return &Token{
		Token:     token,
		UserID:    userID,
		ExpiresAt: expiresAt,
	}
}

func NewTokenWithToken(token string) *Token {
	return &Token{
		Token: token,
	}
}

func (t *Token) GetSQLQuery(procName string) (string, []any) {
	switch procName {
	case "CreateToken":
		return CreateToken, []any{t.Token, t.UserID, t.ExpiresAt}
	case "IsTokenValid":
		return IsTokenValid, []any{t.Token}
	}
	return "", nil
}
