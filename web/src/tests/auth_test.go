package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	userData := map[string]string{
		"username":              "Test User",
		"email":                 "testuser@example.com",
		"password":              "password123",
		"password_confirmation": "password123",
	}
	body, _ := json.Marshal(userData)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/create_user", bytes.NewBuffer(body))
	testRequest(w, *req)
	assert.Equal(t, 201, w.Code)
}
