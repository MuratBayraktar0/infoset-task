package presentation

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"infoset-task/application/usecase"
	"infoset-task/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock api
type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) GetNearbyRestaurants(ctx context.Context, costumerLocation domain.CostumerLocation, distance float64, limit int) (*[]domain.RestaurantBranch, error) {
	args := m.Called(ctx)
	return args.Get(0).(*[]domain.RestaurantBranch), args.Error(1)
}

func TestApi_GetNearbyRestaurants(t *testing.T) {
	mockRepo := &mockRepo{}
	mockRestaurantBranches := []domain.RestaurantBranch{
		{ID: "1", Name: "John Doe", Latitude: 1, Longitude: 1},
		{ID: "2", Name: "Jane Doe 2", Latitude: 2, Longitude: 2},
	}

	mockRepo.On("GetNearbyRestaurants", mock.Anything).Return(&mockRestaurantBranches, nil)

	usecase := usecase.NewUsecase(context.Background(), mockRepo)

	// Create a Fiber app
	app := fiber.New()

	// Create an instance of the Api with mock usecase
	api := NewApi(usecase)

	// Register routes
	api.Router(context.Background(), app)

	reqBody := []byte(`{"latitude": 1, "longitude": 1}`)
	req := httptest.NewRequest("GET", fmt.Sprintf("/nearby-restaurants?distance=10&limit=5"), bytes.NewReader(reqBody))
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var results []domain.RestaurantBranch
	err := json.NewDecoder(resp.Body).Decode(&results)
	assert.NoError(t, err)

	for i, restaurantBranch := range results {
		assert.Equal(t, mockRestaurantBranches[i].ID, restaurantBranch.ID)
		assert.Equal(t, mockRestaurantBranches[i].Name, restaurantBranch.Name)
		assert.Equal(t, mockRestaurantBranches[i].Latitude, restaurantBranch.Latitude)
		assert.Equal(t, mockRestaurantBranches[i].Longitude, restaurantBranch.Longitude)
	}

	mockRepo.AssertExpectations(t)
}
