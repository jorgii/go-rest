# Go REST API
A fully featured REST API written in Go for the sake of trying out the language.

# Description
This is an example implementation of a REST API in Go. It uses the following libraries for the following features:
* [Fiber](https://github.com/gofiber/fiber) as the backbone.
* [Gorm](https://github.com/go-gorm/gorm) as the ORM.
* [Migrate](https://github.com/golang-migrate/migrate) for database migrations.
* [Validator](https://github.com/go-playground/validator) for validation.
* [Cobra](https://github.com/spf13/cobra) for the cli.

# Useful commands
* Start the database - `docker-compose up -d`
* Migrate the database -
	* `export POSTGRESQL_URL='postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable'`
	* `migrate -database ${POSTGRESQL_URL} -path migrations up`
* Start the API - `go run main.go serve`
