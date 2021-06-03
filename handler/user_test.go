package handler

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"gorest/config"
	"gorest/model"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func newTestFiber() (*fiber.App, sqlmock.Sqlmock) {
	sqlDb, sqlMock, _ := sqlmock.New()
	db, _ := gorm.Open(postgres.New(postgres.Config{Conn: sqlDb}), &gorm.Config{})
	h := New(db, config.New())

	app := fiber.New()
	users := app.Group("/users")
	users.Post("", h.CreateUserRequest)
	users.Get("", h.ListUsersRequest)
	return app, sqlMock
}

func TestCreateUserRequest(t *testing.T) {
	app, sqlMock := newTestFiber()
	body := `{"first_name": "John","last_name": "Doe","email": "john.doe@test.com"}`
	req := httptest.NewRequest("POST", "/users", strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	sqlMock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id"}).
		AddRow("1")
	sqlMock.ExpectQuery("INSERT INTO \"users\"").
		WithArgs("John", "Doe", "john.doe@test.com", AnyTime{}).
		WillReturnRows(rows)
	sqlMock.ExpectCommit()
	res, _ := app.Test(req)
	require.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	user := &model.User{}
	d := json.NewDecoder(res.Body)
	if err := d.Decode(user); err != nil {
		t.Fatal("Failed decoding JSON: ", err.Error())
	}

	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "John", user.FirstName)
	assert.Equal(t, "Doe", user.LastName)
	assert.Equal(t, "john.doe@test.com", user.Email)
	assert.False(t, user.CreatedAt.IsZero())
}

func TestCreateUserRequestBadRequest(t *testing.T) {
	app, _ := newTestFiber()
	body := `{"first_name": "John","last_name": "Doe"}`
	req := httptest.NewRequest("POST", "/users", strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	res, _ := app.Test(req)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	got, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("Failed reading response body")
	}
	want := `{"errors":[{"message":"email is required", "property":"email"}]}`
	assert.JSONEq(t, want, string(got))
}

func TestCreateUserRequestDBFail(t *testing.T) {
	app, sqlMock := newTestFiber()
	body := `{"first_name": "John","last_name": "Doe","email": "john.doe@test.com"}`
	req := httptest.NewRequest("POST", "/users", strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	sqlMock.ExpectBegin()
	sqlMock.ExpectQuery("INSERT INTO \"users\"").
		WithArgs("John", "Doe", "john.doe@test.com", AnyTime{}).
		WillReturnError(errors.New("Test"))
	sqlMock.ExpectRollback()
	res, _ := app.Test(req)
	require.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	got, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("Failed reading response body")
	}
	want := `{"errors":[{"code":"internal_server_error", "message":"Internal server error."}]}`
	assert.JSONEq(t, want, string(got))
}

func TestListUsersRequest(t *testing.T) {
	app, sqlMock := newTestFiber()
	req := httptest.NewRequest("GET", "/users", nil)
	req.Header.Add("Content-Type", "application/json")
	sqlMock.ExpectQuery("SELECT count\\(1\\) FROM \"users\"").
		WithArgs().
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("3"))
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "created_at"}).
		AddRow("1", "John1", "Doe1", "john.doe1@test.com", time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)).
		AddRow("2", "John2", "Doe2", "john.doe2@test.com", time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)).
		AddRow("3", "John3", "Doe3", "john.doe3@test.com", time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC))
	sqlMock.ExpectQuery("SELECT \\* FROM \"users\" LIMIT 10").
		WillReturnRows(rows)
	res, _ := app.Test(req)
	require.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Equal(t, http.StatusOK, res.StatusCode)
	got, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("Failed reading response body")
	}
	want := `{"page":1,"count":3,"total_pages":1,"data":[{"id":1,"first_name":"John1","last_name":"Doe1","email":"john.doe1@test.com","created_at":"2021-01-01T00:00:00Z"},{"id":2,"first_name":"John2","last_name":"Doe2","email":"john.doe2@test.com","created_at":"2022-01-01T00:00:00Z"},{"id":3,"first_name":"John3","last_name":"Doe3","email":"john.doe3@test.com","created_at":"2023-01-01T00:00:00Z"}]}`
	assert.JSONEq(t, want, string(got))
}
