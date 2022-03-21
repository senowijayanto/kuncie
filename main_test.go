package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestMacPromo(t *testing.T) {
	// Router
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/checkout", CheckoutItems).Methods(http.MethodPost)

	requestBody := strings.NewReader(`{"sku": "43N23P", "qty": 1}, {"sku": "234234", "qty": 1}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/checkout", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]float32
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, float32(0), responseBody["Total"])
}
