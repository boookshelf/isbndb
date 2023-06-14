package isbndb

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSubject(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))

	defer server.Close()

}
