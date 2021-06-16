package handler

import (
	"encoding/json"
	"gorest/config"
	"gorest/database"
	"gorest/fixture"
	"gorest/model"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func newTestFiber(t *testing.T, withDb bool, fixtures string) *fiber.App {
	c := config.New()
	h := New(nil, c)
	if withDb {
		db, err := database.ConnectDB(c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName+"_test")
		if err != nil {
			t.Fatal(err)
		}
		db = db.Begin()
		db.AutoMigrate(&model.User{}, &model.Book{})
		if fixtures != "" {
			if err := fixture.Load(db, fixtures); err != nil {
				t.Fatal(err)
			}
		}
		t.Cleanup(func() { db.Rollback() })
		h.DB = db
	}
	app := fiber.New()
	users := app.Group("/users")
	users.Post("", h.CreateUserRequest)
	users.Get("", h.ListUsersRequest)
	return app
}

func TestCreateUserRequest(t *testing.T) {
	app := newTestFiber(t, true, "")
	body := `{"first_name": "John","last_name": "Doe","email": "john.doe@test.com"}`
	req := httptest.NewRequest("POST", "/users", strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	res, _ := app.Test(req)
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
	app := newTestFiber(t, false, "")
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

func TestListUsersRequest(t *testing.T) {
	app := newTestFiber(t, true, "test_fixtures")
	req := httptest.NewRequest("GET", "/users", nil)
	req.Header.Add("Content-Type", "application/json")
	res, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	got, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("Failed reading response body")
	}
	want := `{"page":1,"count":3,"total_pages":1,"data":[{"id":1,"first_name":"John1","last_name":"Doe1","email":"john.doe1@test.com","created_at":"2021-01-01T00:00:00Z"},{"id":2,"first_name":"John2","last_name":"Doe2","email":"john.doe2@test.com","created_at":"2022-01-01T00:00:00Z"},{"id":3,"first_name":"John3","last_name":"Doe3","email":"john.doe3@test.com","created_at":"2023-01-01T00:00:00Z"}]}`
	assert.JSONEq(t, want, string(got))
}
