# Link shortener

## Migrations
You first need to install [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

### Creating migration
`migrate create -ext sql -dir db/migrations -seq <MIGRATION_NAME>`

### Running migrations

#### Dev
`source .env && migrate -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable" -path db/migrations up`