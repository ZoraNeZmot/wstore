package storage

import (
	"context"
	"wstore/internal/storage/dbtokens"
)

func (d *Database) CreateToken(ctx context.Context, token *dbtokens.Token) error {
	query, args := token.GetSQLQuery("CreateToken")
	_, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) IsTokenValid(ctx context.Context, token *dbtokens.Token) (bool, error) {
	query, args := token.GetSQLQuery("IsTokenValid")
	row := d.db.QueryRowContext(ctx, query, args...)
	var valid bool
	err := row.Scan(&valid)
	if err != nil {
		return false, err
	}
	return valid, nil
}

func (d *Database) RevokeToken(ctx context.Context, token *dbtokens.Token) error {
	query, args := token.GetSQLQuery("RevokeToken")
	_, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}
