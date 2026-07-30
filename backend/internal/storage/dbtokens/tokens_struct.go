package dbtokens

import (
	"time"

	"github.com/google/uuid"
)

type TokenOpt func(*Token)

type Token struct {
	Token     string    `json:"token"`
	UserID    uuid.UUID `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
}

func WithToken(token string) TokenOpt {
	return func(t *Token) {
		t.Token = token
	}
}

func WithUserID(userID uuid.UUID) TokenOpt {
	return func(t *Token) {
		t.UserID = userID
	}
}

func WithExpiresAt(expiresAt time.Time) TokenOpt {
	return func(t *Token) {
		t.ExpiresAt = expiresAt
	}
}
func NewToken(opts ...TokenOpt) *Token {
	token := &Token{}
	for _, opt := range opts {
		opt(token)
	}
	return token
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
