package utils

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
)

func InjectDB(db *sqlx.DB, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "db", db)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func GetDbConnection() *sqlx.DB {
	dbUser := os.Getenv(PgUser)
	dbPassword := os.Getenv(PgPassword)
	dbName := os.Getenv(PgDb)
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)

	var err error
	db, err := sqlx.Connect(Postgres, dsn)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Succesfully connected to DB")
	return db
}
