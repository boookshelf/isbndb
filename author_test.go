package isbndb

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAuthor(t *testing.T) {
	author := Author{
		Author: "testing",
		Books: []Book{
			{
				Title: "Test Title",
				ISBN:  "123634535",
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/author/testing" {
			t.Errorf("Expected to request '/author/testing', got: %s", r.URL.Path)
		}
		authorBytes, _ := json.Marshal(author)
		w.WriteHeader(http.StatusOK)
		w.Write(authorBytes)
	}))

	defer server.Close()

	client := New(http.DefaultClient)
	client.SetURL(server.URL)

	response, err := client.GetAuthor(context.TODO(), "testing", nil)
	if err != nil {
		t.Errorf("Expected no error: %v", err)
	}

	if response.Author != author.Author {
		t.Errorf("Expected authors to match")
	}

	if response.Books[0].Title != author.Books[0].Title {
		t.Errorf("Returned book titles should match")
	}
}

func TestQueryAuthors(t *testing.T) {
	authorQuery := AuthorQuery{
		Total:   3,
		Authors: []string{"one", "two", "three"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/authors/testing" {
			t.Errorf("Expected to request '/authors/testing', got: %s", r.URL.Path)
		}
		authorBytes, _ := json.Marshal(authorQuery)
		w.WriteHeader(http.StatusOK)
		w.Write(authorBytes)
	}))

	defer server.Close()

	client := New(http.DefaultClient)
	client.SetURL(server.URL)

	response, err := client.QueryAuthors(context.TODO(), "testing", nil)
	if err != nil {
		t.Errorf("Expected no error: %v", err)
	}

	if response.Total != authorQuery.Total {
		t.Errorf("Expected totals to match")
	}

	for index, val := range authorQuery.Authors {
		if val != response.Authors[index] {
			t.Errorf("Expect all returned authors to match")
		}
	}
}
