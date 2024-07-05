package test

import (
	"encoding/json"
	"strconv"
	// "strconv"
	// "fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func TestSignUpSuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "users")
	router := SetupRouter(db)

	requestBody := strings.NewReader(`{
		"username": "test",
		"full_name": "Test",
		"password": "test"
	}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/users/signup", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 201, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "test", responseBody["data"].(map[string]any)["username"])
	assert.Equal(t, "Test", responseBody["data"].(map[string]any)["full_name"])
}

func TestSignUpFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "users")
	router := SetupRouter(db)

	requestBody := strings.NewReader(`{
		"username": "t",
		"full_name": "Test",
		"password": "test"
	}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/users/signup", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
}

func TestSignInSuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "users")
	router := SetupRouter(db)

	CreateUser(db, "test")

	requestBody := strings.NewReader(`{
		"username": "test",
		"password": "test"
	}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/users/signin", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "test", responseBody["data"].(map[string]any)["username"])
	assert.Equal(t, "Test", responseBody["data"].(map[string]any)["full_name"])
}

func TestSignInFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "users")
	router := SetupRouter(db)

	CreateUser(db, "test")

	requestBody := strings.NewReader(`{
		"username": "test",
		"password": "tes"
	}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/users/signin", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)
}

func TestSignOutSuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "users")
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/v1/users/signout", nil)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestSignOutFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "users")
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/v1/users/signout", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)
}

func TestCurrentUserSuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "users")
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/users/currentuser", nil)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "test", responseBody["data"].(map[string]any)["username"])
	assert.Equal(t, "Test", responseBody["data"].(map[string]any)["full_name"])
}

func TestCurrentUserFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "users")
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/users/currentuser", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)
}

func TestUpdateSuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "users")
	router := SetupRouter(db)

	user := CreateUser(db, "test")

	requestBody := strings.NewReader(`{
		"username": "testupdated",
		"full_name": "Test",
		"password": "test"
	}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/v1/users/" + strconv.Itoa(user.ID) , requestBody)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "testupdated", responseBody["data"].(map[string]any)["username"])
}

func TestUpdateFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users")
	router := SetupRouter(db)

	user := CreateUser(db, "test")

	requestBody := strings.NewReader(`{
		"username": "testupdated",
		"full_name": "Test",
		"password": "test"
	}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/v1/users/1" + strconv.Itoa(user.ID) , requestBody)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}


