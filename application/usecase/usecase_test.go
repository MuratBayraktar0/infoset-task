package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"infoset-task/domain"
)

// Mock repository
type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) GetNearbyRestaurants(ctx context.Context, costumerLocation domain.CostumerLocation, distance float64, limit int) (*[]domain.RestaurantBranch, error) {
	args := m.Called(ctx)
	return args.Get(0).(*[]domain.RestaurantBranch), args.Error(1)
}

// ... Implement other repository methods ...
func TestUsecase_GetNearbyRestaurants(t *testing.T) {
	mockRepo := &mockRepo{}

	mockRestaurantBranches := []domain.RestaurantBranch{
		{ID: "1", Name: "John Doe", Latitude: 1, Longitude: 1},
		{ID: "2", Name: "Jane Doe 2", Latitude: 2, Longitude: 2},
	}

	mockRepo.On("GetNearbyRestaurants", mock.Anything).Return(&mockRestaurantBranches, nil)

	ctx := context.Background()
	usecase := NewUsecase(ctx, mockRepo)

	results, err := usecase.GetNearbyRestaurants(ctx, domain.CostumerLocation{Latitude: 1, Longitude: 1}, 10, 5)

	assert.NoError(t, err)
	assert.Len(t, *results, len(mockRestaurantBranches))

	for i, restaurantBranch := range *results {
		assert.Equal(t, mockRestaurantBranches[i].ID, restaurantBranch.ID)
		assert.Equal(t, mockRestaurantBranches[i].Name, restaurantBranch.Name)
		assert.Equal(t, mockRestaurantBranches[i].Latitude, restaurantBranch.Latitude)
		assert.Equal(t, mockRestaurantBranches[i].Longitude, restaurantBranch.Longitude)
	}

	mockRepo.AssertExpectations(t)
}
