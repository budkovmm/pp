package utils

import (
	"github.com/joho/godotenv"
	"log"
)

const (
	Postgres   = "postgres"
	PgUser     = "POSTGRES_USER"
	PgPassword = "POSTGRES_PASSWORD"
	PgDb       = "POSTGRES_DB"
	PgHost     = "POSTGRES_HOST"
	PgPort     = "POSTGRES_PORT"
	PgURL      = "POSTGRES_URL"
)

func LoadEnvs() {
	err := godotenv.Load("./configs/db/pg.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
