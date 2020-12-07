package utils

import (
	"context"
	"github.com/jmoiron/sqlx"
	"net/http"
)

const (
	dbContextKey = "db"
)

func GetDbFromContext(ctx context.Context) (*sqlx.DB, error){
	db, ok := ctx.Value("db").(*sqlx.DB)
	if !ok {
		return nil, NoDbInContext
	}
	return db, nil
}

func InjectDB(db *sqlx.DB, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), dbContextKey, db)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}