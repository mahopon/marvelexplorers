package services_test

import (
	"net/http"
	"net/http/httptest"
	"tcy/marvelexplorers/handler" // Ensure this import remains for the handler package
	model "tcy/marvelexplorers/model/db"
	"tcy/marvelexplorers/repository/mock"
	"tcy/marvelexplorers/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCharacters(t *testing.T) {
	mockRepo := &mock.MockCharacterRepo{
		Characters: []model.Character_db{
			{Name: "Iron Man"},
			{Name: "Spider-Man"},
		},
	}
	svc := services.CharacterService{
		Repo: mockRepo,
	}
	h := handler.CharacterHandler{
		Service: &svc, // Ensure svc implements the interface
	}

	// Use httptest to simulate requests to your handler
	req := httptest.NewRequest("GET", "/api/characters", nil)
	rr := httptest.NewRecorder()
	// Call the handler with the mock request and recorder
	h.GetCharacters(rr, req)

	// Assert that the status code is OK
	assert.Equal(t, http.StatusOK, rr.Code)

	// You can add further assertions to check if the response body is as expected.
	// For example, you can check if the correct characters are returned.
	expectedBody := `[{"Name":"Iron Man"},{"Name":"Spider-Man"}]`
	assert.JSONEq(t, expectedBody, rr.Body.String())
}
