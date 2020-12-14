package utils

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

func GetPgDbConnection() *sqlx.DB {
	dbUser := os.Getenv(PgUser)
	dbPassword := os.Getenv(PgPassword)
	dbName := os.Getenv(PgDb)
	dbHost := os.Getenv(PgHost)
	dsn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbHost, dbName)

	var err error
	db, err := sqlx.Connect(Postgres, dsn)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Succesfully connected to Postgres DB")
	return db
}
