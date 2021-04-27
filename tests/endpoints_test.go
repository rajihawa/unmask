package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rajihawa/unmask/app"
	"github.com/rajihawa/unmask/app/database"
	"github.com/rajihawa/unmask/app/handlers"
)

// TestHealthEndpoint - tests health endpoint of server
func TestHealthEndpoint(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/api/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HealthHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "{\"healthy\":true}\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestProjectsEndpoint(t *testing.T) {
	app.StartApp(app.AppConfig{
		Port: "4000",
		RethinkDB: database.RethinkConfig{
			DatabaseURL:      "localhost",
			DatabaseName:     "unmask_test",
			DatabaseUsername: "unmask_test",
			DatabasePassword: "",
		},
	})
}
