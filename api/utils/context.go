package utils

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
)

func GetDbFromContext(ctx context.Context) (*sqlx.DB, error){
	db, ok := ctx.Value("db").(*sqlx.DB)
	if !ok {
		return nil, errors.New("could not get database connection pool from context")
	}
	return db, nil
}