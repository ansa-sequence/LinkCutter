package ansaserver

import (
	"LinkCutter/internal/model"
	"LinkCutter/internal/store/teststore"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New())
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]interface{}{
				"email":    "user@example.org",
				"password": "secret",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]interface{}{
				"email":    "invalid",
				"password": "short",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
func TestServer_HandleUserRemove(t *testing.T) {
	s := newServer(teststore.New())
	testCases := []struct {
		name         string
		caller       func()
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			caller: func() {
				_ = s.store.User().Create(&model.User{Email: "user@example.org", Password: "secret"})
			},
			payload: map[string]interface{}{
				"email":    "user@example.org",
				"password": "secret",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:   "invalid",
			caller: nil,
			payload: map[string]interface{}{
				"email":    "user2125215@example.org",
				"password": "password",
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			if tc.caller != nil {
				tc.caller()
			}
			json.NewEncoder(b).Encode(tc.payload)
			fmt.Print(b)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodDelete, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
