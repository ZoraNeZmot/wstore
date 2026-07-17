package dbusers

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func NewUser(email string, password string, name string) *User {
	return &User{
		Email:    email,
		Password: password,
		Name:     name,
	}
}

func NewUserWithEmail(email string) *User {
	return &User{
		Email: email,
	}
}

func NewUserWithUid(uid string) *User {
	uidUUID, err := uuid.Parse(uid)
	if err != nil {
		return nil
	}
	return &User{
		ID:       uidUUID,
		Email:    "",
		Password: "",
		Name:     "",
	}
}

func NewOAuthUser(email string, name string) *User {
	return &User{
		Email: email,
		Name:  name,
	}
}

func (u *User) GetSQLQuery(procName string) (string, []any) {
	switch procName {
	case "CreateUser":
		return CreateUser, []any{u.Email, u.Password, u.Name}
	case "CreateOAuthUser":
		return CreateOAuthUser, []any{u.Email, u.Name}
	case "IsUserExists":
		return IsUserExists, []any{u.Email}
	case "GetUserByEmail":
		return GetUserByEmail, []any{u.Email}
	case "GetUserByUid":
		return GetUserByUid, []any{u.ID}
	case "GetUsers":
		return GetUsers, []any{}
	default:
		return "", nil
	}
}
