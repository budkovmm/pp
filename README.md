[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=budkovmm_ppCleanArchitecture&metric=alert_status)](https://sonarcloud.io/dashboard?id=budkovmm_ppCleanArchitecture)

# Pet project in Go

## Issues
[Trello](https://trello.com/b/pPMbbRWT/budkovgopetproject)

## Env
Storing env files in the repository is not a good practice, don't use it for production

`source ./confgis/db/pg.env`

## Run DB
`docker-compose up`

## Migrations
You need to install [go-migrate](https://github.com/golang-migrate/migrate)

`export POSTGRESQL_URL='postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_DB}?sslmode=disable'`

for example:

`export POSTGRESQL_URL='postgres://test:password@0.0.0.0:5432/go_sample?sslmode=disable'`

Create new migration:

`migrate create -ext sql -dir db/migrations -seq create_users_table`

Run migrations:

`migrate -database ${POSTGRESQL_URL} -path db/migrations up`

Rollback migrations:

`migrate -database ${POSTGRESQL_URL} -path db/migrations down`


## Run application
`go run src/server.go`