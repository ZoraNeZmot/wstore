package storage

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(id uuid.UUID, email string, password string, name string) *User {
	return &User{
		ID:        id.String(),
		Email:     email,
		Password:  password,
		Name:      name,
		CreatedAt: time.Now(),
	}
}

func NewUserWithEmail(email string) *User {
	return &User{
		Email: email,
	}
}

func (u *User) GetSQLQuery(procName string) (string, []any) {
	switch procName {
	case "AddUser":
		return AddUser, []any{u.ID, u.Email, u.Password, u.Name, u.CreatedAt}
	case "CheckUser":
		return CheckUser, []any{u.Email}
	case "GetUserByEmail":
		return GetUserByEmail, []any{u.Email}
	default:
		return "", nil
	}
}
