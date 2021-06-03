# Go REST API
A fully featured REST API written in Go for the sake of trying out the language.

# Description
This is an example implementation of a REST API in Go. It uses the following libraries for the following features:
* [Fiber](https://github.com/gofiber/fiber) as the backbone.
* [Gorm](https://github.com/go-gorm/gorm) as the ORM.
* [Migrate](https://github.com/golang-migrate/migrate) for database migrations.
* [Validator](https://github.com/go-playground/validator) for validation.
* [Cobra](https://github.com/spf13/cobra) for the cli.
* [Env](https://github.com/caarlos0/env) for configuration.
* [JWT-Go](https://github.com/dgrijalva/jwt-go) for JWT authentication.
* [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) for mocking the db driver during unit tests.
* [Testify](https://github.com/stretchr/testify) for assertions in unit tests.
* [Postman collection](postman_collection.json)

# Useful commands
* Start the database: `docker-compose up -d`
* Migrate the database:
	* `export POSTGRESQL_URL='postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable'`
	* `migrate -database ${POSTGRESQL_URL} -path migrations up`
* Start the API: `go run main.go serve`

# Using the API

## Create a user
```bash
curl --location --request POST 'http://localhost:8080/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@test.com"
}'
```
## JWT

### Use this pregenerated JWT (expires in year 2100)
`eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IkttQzh1TjBNVzFwYWF3T0I3emZ6WllqS0djc19VckpHSHRVc3ljNmRRSlEifQ.eyJleHAiOjQxMTU3MDQzMzAsImlhdCI6MTYyMjcyODM2MCwianRpIjoiN2RlMGY2ZTgtMzk1OC00Njc5LWJkN2ItNGFiMTAyZGNkNmViIiwiaXNzIjoiaHR0cDovL3Nzby50ZXN0LmNvbS9hdXRoL3JlYWxtcy90ZXN0IiwiYXVkIjpbImNvbm5lY3RfYXBpIiwiYWNjb3VudCJdLCJzdWIiOiI5NDg1MWYyYi03ZDBiLTRmMzctYmY0MC1kNTA0ZTU3YzIzMjEiLCJ0eXAiOiJCZWFyZXIiLCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiZW1haWwiOiJqb2huLmRvZUB0ZXN0LmNvbSJ9.D9KCU6t6ZY5NjfMBD2EZ1y9aSQcSJy6Bb7ABSmWvUN-4Ud2QZIcdDRRHRpPIIdba-mYFCpNPr5OVLCsB6cqzWpyAfHdxIfW9aqL9sUs-iL0Vj-ddtmmKGeyrw2z5_Jb0lcm2b9LuzbO4nnDdX1fFbh6VfNPaJDu80wR9Goh0IwE`

### Alternatively generate a JWT
 * Open [jwt.io](https://jwt.io/)
 * Choose algorithm RS256
 * Enter payload (the expiration is year 2100):
```json
{
  "exp": 4115704330,
  "iat": 1622728360,
  "jti": "7de0f6e8-3958-4679-bd7b-4ab102dcd6eb",
  "iss": "http://sso.test.com/auth/realms/test",
  "aud": [
    "connect_api",
    "account"
  ],
  "sub": "94851f2b-7d0b-4f37-bf40-d504e57c2321",
  "typ": "Bearer",
  "scope": "email profile",
  "email_verified": true,
  "email": "john.doe@test.com"
}
```
 * Enter public key (this key is the default in [config/config.go](config/config.go)):
```
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCMPspipOA0/sYR8udpBFT+U8IT
3ynuNMJGuXGaqiawcQqVKFUPMxGrhbS/kp2WCbXNx7ykBB4VFoyAjKS1/rZ2Eaip
cm1vIFa3arDztutrRVjO2yxuDfupWwrZDqYEBf4gqKVwFCO0zjywR6x7/Tf56jcu
5B7PVew7botJgUbZiwIDAQAB
-----END PUBLIC KEY-----
```
 * Enter private key:
```
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCMPspipOA0/sYR8udpBFT+U8IT3ynuNMJGuXGaqiawcQqVKFUP
MxGrhbS/kp2WCbXNx7ykBB4VFoyAjKS1/rZ2Eaipcm1vIFa3arDztutrRVjO2yxu
DfupWwrZDqYEBf4gqKVwFCO0zjywR6x7/Tf56jcu5B7PVew7botJgUbZiwIDAQAB
AoGBAIwyvQlNv2DbDFCnHdTq4rh4LLzGy9j4XvpqqfmufQzHhIfFkPqn19M6z4zv
WZ/Cxz8WnCruftAgAYcEkifpoKX8PmDjknEUK9WhB5Pta1MuPLvlyVsWWpjPsRXS
am7nAcqMeT7w5a8ksZBaum56vRak/SpLwBuXf6/ghTyi6aMBAkEA0z9rK5BmNKFX
3N25xKBpIy0eVr/2uN+RduX0hdHjIme2raiwgQhJ+UsaLFLXp20o5oE+AWZiwOVN
tiwqL33qgQJBAKn0s5WJKGOtTRdyTwoGs+4xgtwoaOA20esdkWuCd99IXgW91f/N
tWqvKZwobBcCGSqNUUS2Dwu/Y3nS4Kh+xgsCQAyQhxVOP2X9+rXeUkBJsjcvZdCP
FmOkmIhT4RlchH/1Xz4w/F6QWaYEO4hLXRxON9KN+Vwn9NO4T8j4E4JDoQECQBYm
MuVRG7dp1XsxxYU7/GUVhMVmasyVuGPY1aPaO/8YclBzPhl5WKRPsa+NQRD6kAcQ
pb+7rlhMmA0/Y1HyO6MCQEECbpG+SlsKhE9dGiWxnWsp+ZbO3UrY/+4iM9VZvj1D
gzVjwMxOQjpBzx8PEgZv/kLImhFFI1d89jD8YuitT6c=
-----END RSA PRIVATE KEY-----
```
 * Copy the generated JWT
## Create a book
```bash
curl --location --request POST 'http://localhost:8080/books' \
--header 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IkttQzh1TjBNVzFwYWF3T0I3emZ6WllqS0djc19VckpHSHRVc3ljNmRRSlEifQ.eyJleHAiOjQxMTU3MDQzMzAsImlhdCI6MTYyMjcyODM2MCwianRpIjoiN2RlMGY2ZTgtMzk1OC00Njc5LWJkN2ItNGFiMTAyZGNkNmViIiwiaXNzIjoiaHR0cDovL3Nzby50ZXN0LmNvbS9hdXRoL3JlYWxtcy90ZXN0IiwiYXVkIjpbImNvbm5lY3RfYXBpIiwiYWNjb3VudCJdLCJzdWIiOiI5NDg1MWYyYi03ZDBiLTRmMzctYmY0MC1kNTA0ZTU3YzIzMjEiLCJ0eXAiOiJCZWFyZXIiLCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiZW1haWwiOiJqb2huLmRvZUB0ZXN0LmNvbSJ9.D9KCU6t6ZY5NjfMBD2EZ1y9aSQcSJy6Bb7ABSmWvUN-4Ud2QZIcdDRRHRpPIIdba-mYFCpNPr5OVLCsB6cqzWpyAfHdxIfW9aqL9sUs-iL0Vj-ddtmmKGeyrw2z5_Jb0lcm2b9LuzbO4nnDdX1fFbh6VfNPaJDu80wR9Goh0IwE' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "National Tactics Manager",
    "author": "Ms. Amy Weimann"
}'
```
