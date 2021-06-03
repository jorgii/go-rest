package handler

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"gorest/model"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	app     *fiber.App
	sqlMock sqlmock.Sqlmock
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func setApp() {
	var sqlDb *sql.DB
	sqlDb, sqlMock, _ = sqlmock.New()
	db, _ := gorm.Open(postgres.New(postgres.Config{Conn: sqlDb}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	app = fiber.New()
	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Locals("db", db.Debug())
		return ctx.Next()
	})
	users := app.Group("/users")
	users.Post("", CreateUserRequest)
	users.Get("", ListUsersRequest)
}

func TestMain(m *testing.M) {
	setApp()
	os.Exit(m.Run())
}

func TestCreateUserRequest(t *testing.T) {
	body := `{"first_name": "George","last_name": "Goranov","email": "g.p.goranov@gmail.com"}`
	req := httptest.NewRequest("POST", "/users", strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	sqlMock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(1)
	sqlMock.ExpectQuery("^INSERT INTO \"users\".* RETURNING \"id\"$").
		WithArgs("George", "Goranov", "g.p.goranov@gmail.com", AnyTime{}).
		WillReturnRows(rows)
	sqlMock.ExpectCommit()
	res, _ := app.Test(req)
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Create user request returned unexpected status code, %d.", res.StatusCode)
	}
	user := &model.User{}
	d := json.NewDecoder(res.Body)
	if err := d.Decode(user); err != nil {
		t.Error("Failed decoding JSON: ", err.Error())
	}
	if user.FirstName != "George" {
		t.Error("First name does not match")
	}
	if user.LastName != "Goranov" {
		t.Error("Last name does not match")
	}
	if user.Email != "g.p.goranov@gmail.com" {
		t.Error("Email does not match")
	}
	if user.CreatedAt.IsZero() {
		t.Error("Created at is missing.")
	}
	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Error("Create user did not execute the expected query", err.Error())
	}
}
