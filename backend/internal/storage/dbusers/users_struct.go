package dbusers

import (
	"time"

	"github.com/google/uuid"
)

type UserOpt func(*User)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	ImagePath string    `json:"image_path,omitempty"`
}

func WithEmail(email string) UserOpt {
	return func(u *User) {
		u.Email = email
	}
}

func WithPassword(password string) UserOpt {
	return func(u *User) {
		u.Password = password
	}
}

func WithName(name string) UserOpt {
	return func(u *User) {
		u.Name = name
	}
}

func WithCreatedAt(createdAt time.Time) UserOpt {
	return func(u *User) {
		u.CreatedAt = createdAt
	}
}

func WithImagePath(imagePath string) UserOpt {
	return func(u *User) {
		u.ImagePath = imagePath
	}
}

func WithID(id uuid.UUID) UserOpt {
	return func(u *User) {
		u.ID = id
	}
}

func NewUser(opts ...UserOpt) *User {
	user := &User{
		CreatedAt: time.Now(),
	}
	for _, opt := range opts {
		opt(user)
	}
	return user
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
	case "GetUserByid":
		return GetUserByid, []any{u.ID}
	case "GetUsers":
		return GetUsers, []any{}
	case "GetUserProfileImage":
		return GetUserProfileImage, []any{u.ID}
	case "CreateUserProfileImage":
		return CreateUserProfileImage, []any{u.ID, u.ImagePath}
	default:
		return "", nil
	}
}
