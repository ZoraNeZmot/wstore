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

func (d *Database) CreateUser(ctx context.Context, user *dbusers.User) (int64, error) {
	_, err := d.ExecQuery(ctx, "CreateUser", user)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (d *Database) CreateOAuthUser(ctx context.Context, user *dbusers.User) (uuid.UUID, error) {
	var uid uuid.UUID
	err := d.QueryRow(ctx, "CreateOAuthUser", user).Scan(&uid)
	return uid, err
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

func (d *Database) GetUserByUid(ctx context.Context, user *dbusers.User) (*dbusers.User, error) {
	var userData dbusers.User
	result := d.QueryRow(ctx, "GetUserByUid", user)
	err := result.Scan(&userData.ID, &userData.Email, &userData.Name)
	return &userData, err
}
