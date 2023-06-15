package isbndb

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetSubject(t *testing.T) {
	expectedSubject := Subject{
		Subject: "TestSubject",
		Parent:  "TestParent",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/subject/testing" {
			t.Error("Recieved URL did not match expected")
		}

		subjectBytes, _ := json.Marshal(expectedSubject)
		w.WriteHeader(http.StatusOK)
		w.Write(subjectBytes)
	}))

	defer server.Close()

	serverURL, _ := url.Parse(server.URL)
	client := New(WithURL(serverURL))

	response, err := client.GetSubject(context.TODO(), "testing")
	if err != nil {
		t.Errorf("Expected to receive no error: %v", err)
	}

	if response.Parent != expectedSubject.Parent {
		t.Errorf("Expected parents to match")
	}

	if response.Subject != expectedSubject.Subject {
		t.Errorf("Expected subjects to match")
	}
}

func TestQuerySubjects(t *testing.T) {
	expectedSubjectQuery := SubjectQuery{
		Total:    3,
		Subjects: []string{"one", "two", "three"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/subjects/testing" {
			t.Error("Recieved URL did not match expected")
		}

		subjectBytes, _ := json.Marshal(expectedSubjectQuery)
		w.WriteHeader(http.StatusOK)
		w.Write(subjectBytes)
	}))

	defer server.Close()

	serverURL, _ := url.Parse(server.URL)
	client := New(WithURL(serverURL))

	response, err := client.QuerySubject(context.TODO(), "testing")
	if err != nil {
		t.Errorf("Expected to receive no error: %v", err)
	}

	if response.Total != expectedSubjectQuery.Total {
		t.Errorf("Expected totals to match")
	}

	for index, subject := range expectedSubjectQuery.Subjects {
		if subject != response.Subjects[index] {
			t.Errorf("Expected subjects to match")
		}
	}
}
