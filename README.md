# Link shortener

## DevEnv quickstart
1. Make sure [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) is installed
2. Make sure docker is installed and running
3. Set up `.env` file - copy from `.env.example` and fill in missing values
4. `docker compose up -d`
5. Run migrations `source .env && migrate -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable" -path db/migrations up`

## Architecture
- Single service for both generating shortlinks and redirecting via short links
  - Under heavy loads, it'd be better to separate Redirect API and Generate API to scale better
- Single DB & Cache instance
  - Under heavy loads, one should add read replicas. To keep things simple, I suggest using [Pgpool](https://pgpool.net/mediawiki/index.php?title=Main_Page), but other options could also work
- Plays nice with CDNs
  - 302 Redirects with Cache-Control for great edge caching
- Simple shortlink generation - random string
  - Allows duplicate urls to be shortened
  - Might be better to migrate to smth like Snowflake IDs