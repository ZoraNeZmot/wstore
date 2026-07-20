package storage

import (
	"context"
	"wstore/internal/storage/dbusers"

	"github.com/google/uuid"
)

/*
	USER RELATED QUERIES
*/

func (d *Database) VerifyUser(ctx context.Context, user *dbusers.User) bool {
	var checked bool
	result := d.QueryRow(ctx, "VerifyUser", user)
	result.Scan(&checked)
	return checked
}

func (d *Database) GetUserByEmail(ctx context.Context, user *dbusers.User) (*dbusers.User, error) {
	var userData dbusers.User
	result := d.QueryRow(ctx, "GetUserByEmail", user)
	err := result.Scan(&userData.ID, &userData.Email, &userData.Password, &userData.Name)
	return &userData, err
}

func (d *Database) IsUserExists(ctx context.Context, user *dbusers.User) bool {
	var exists bool
	result := d.QueryRow(ctx, "IsUserExists", user)
	result.Scan(&exists)
	return exists
}

func (d *Database) CreateUser(ctx context.Context, user *dbusers.User) (uuid.UUID, error) {
	var id uuid.UUID
	err := d.QueryRow(ctx, "CreateUser", user).Scan(&id)
	return id, err
}

func (d *Database) CreateOAuthUser(ctx context.Context, user *dbusers.User) (uuid.UUID, error) {
	var id uuid.UUID
	err := d.QueryRow(ctx, "CreateOAuthUser", user).Scan(&id)
	return id, err
}

func (d *Database) GetUsers(ctx context.Context) ([]dbusers.User, error) {
	var users []dbusers.User
	result, err := d.Query(ctx, "GetUsers", &dbusers.User{})
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		var user dbusers.User
		err = result.Scan(&user.ID, &user.Email, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (d *Database) GetUserByid(ctx context.Context, user *dbusers.User) (*dbusers.User, error) {
	var userData dbusers.User
	result := d.QueryRow(ctx, "GetUserByid", user)
	err := result.Scan(&userData.ID, &userData.Email, &userData.Name)
	return &userData, err
}

func (d *Database) GetUserProfileImage(ctx context.Context, user *dbusers.User) (string, error) {
	var imagePath string
	result := d.QueryRow(ctx, "GetUserProfileImage", user)
	err := result.Scan(&imagePath)
	return imagePath, err
}

func (d *Database) CreateUserProfileImage(ctx context.Context, user *dbusers.User) error {
	_, err := d.ExecQuery(ctx, "CreateUserProfileImage", user)
	return err
}
