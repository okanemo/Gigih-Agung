# Golang Back-end

## How to Run

Initialize project by:

1. `make postgres`
2. `make createdb`
3. `make migrateup`
4. `make server`

`make server` will make the app run on localhost port `8080`

`make test` will provide test coverage using random data.


## Installing Dependencies

`go mod tidy` is your friend.

Dependencies for this project:
- Golang-Migrate CLI (https://github.com/golang-migrate/migrate) (required)
- SQLC (https://docs.sqlc.dev/en/latest/overview/install.html) (optional)


## Common Errors

- Postgres/n
Can't run `make postgres` or `make createdb`? A common error is running Postgres Docker container while having Postgres already installed in your PC. This will cause conflict and error. Solution is to uninstall Postgres.

- SQLC/n
When running `make sqlc` your terminal might say "sqlc.json doesn't exist". You need to specify your project directory. Manually typing down the location works if $(pwd) somehow doesn't work.


## Features:

1. Basic CRUD queries in Go.
2. Test coverage for queries with random data.
3. Endpoints, although only for GET and POST at the moment.


## Endpoints

- **/accounts**

| Method | Header | Params | JSON                                                      |
| ------ | ------ | ------ | --------------------------------------------------------- |
| `POST` | `none` | `none` | username: `string`<br>line_id: `string` <br> email: `string` |
| `GET` | `none` | page_id: min=1 <br> page_size: min=5, max=10 | `none` |

- **/accounts/:id**

| Method | Header | Params | JSON                                    |
| ------ | ------ | ------ | --------------------------------------- |
| `GET` | `none` | `none` | `none` |
