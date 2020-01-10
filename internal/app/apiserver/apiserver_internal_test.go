package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandlePing(t *testing.T) {
	server := New(NewConfig())
	rec := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/ping", nil)

	server.handlePing().ServeHTTP(rec, request)

	assert.Equal(t, "pong", rec.Body.String())
}
