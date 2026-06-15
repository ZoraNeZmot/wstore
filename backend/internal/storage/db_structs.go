package storage

import "time"

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(id string, email string, password string, name string) *User {
	return &User{
		ID:        id,
		Email:     email,
		Password:  password,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *User) GetSQLQuery() string {
	return AddUser
}

func (u *User) GetArgs() []any {
	return []any{u.ID, u.Email, u.Password, u.Name, u.CreatedAt}
}
