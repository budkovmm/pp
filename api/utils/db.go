package utils

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

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
