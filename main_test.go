package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sachin2605/cast0r/controller"
	"github.com/sachin2605/cast0r/models"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestListFruitHandler(t *testing.T) {
	r := SetUpRouter()
	fruitCtrl := new(controller.FruitController)
	r.GET("/api/v1/fruits", fruitCtrl.List)

	req, _ := http.NewRequest("GET", "/api/v1/fruits", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostFruitHandler(t *testing.T) {
	r := SetUpRouter()
	fruitCtrl := new(controller.FruitController)
	r.POST("/api/v1/fruits", fruitCtrl.Create)
	company := models.Fruit{
		Fruit: "Orange",
		Color: "Orange",
	}
	jsonValue, _ := json.Marshal(company)
	req, _ := http.NewRequest("POST", "/api/v1/fruits", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetFruitNotFoundHandler(t *testing.T) {
	r := SetUpRouter()
	fruitCtrl := new(controller.FruitController)
	r.POST("/api/v1/fruits/:id", fruitCtrl.Get)

	req, _ := http.NewRequest("get", "/api/v1/fruits/random-id", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
