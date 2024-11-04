package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setUpRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

func Test_application_GetAllDogBreedsJSON(t *testing.T) {
	r := setUpRouter()
	r.GET("/api/dog-breeds", testApp.GetAllDogBreedsJSON)
	req, _ := http.NewRequest("GET", "/api/dog-breeds", nil)

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}

func Test_application_GetAllCatBreeds(t *testing.T) {
	r := setUpRouter()
	r.GET("api/cat-breeds", testApp.GetAllCatBreeds)
	req, _ := http.NewRequest("GET", "/api/cat-breeds", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}
