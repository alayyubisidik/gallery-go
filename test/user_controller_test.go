package test

import (
	"encoding/json"
	"gallery_go/database"
	"gallery_go/helper"
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"

)

func TestSignUpSuccess(t *testing.T) {
	envPath, err := filepath.Abs("../.env")
	helper.PanicIfError(err)
	err = godotenv.Load(envPath)
	helper.PanicIfError(err)

	database.ConnectDatabase()
	app := gin.Default()
	DeleteTestUsernames(database.DB)
	router := InitRouteTest(app)

	requestBody := strings.NewReader(`{
		"username": "test",
		"full_name": "Test",
		"email": "test@gmail.com",
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
	assert.Equal(t, "test@gmail.com", responseBody["data"].(map[string]any)["email"])
	assert.Equal(t, "author", responseBody["data"].(map[string]any)["role"])
}

func TestSignUpFailed(t *testing.T) {
	envPath, err := filepath.Abs("../.env")
	helper.PanicIfError(err)
	err = godotenv.Load(envPath)
	helper.PanicIfError(err)

	database.ConnectDatabase()
	app := gin.Default()
	DeleteTestUsernames(database.DB)
	router := InitRouteTest(app)

	requestBody := strings.NewReader(`{
		"usernam": "te",
		"full_name": "Test",
		"email": "test@gmail.com",
		"password": "test"
	}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/users/signup", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
}