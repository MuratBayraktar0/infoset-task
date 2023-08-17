package usecase

import (
	"context"
	"infoset-task/domain"
)

// Usecase manages usage scenarios in the business logic layer.
type Usecase struct {
	repository     domain.Repository
	contextTimeout context.Context
}

// NewUsecase creates a new instance of a usage scenario.
func NewUsecase(timeout context.Context, Repository domain.Repository) *Usecase {
	return &Usecase{
		repository:     Repository,
		contextTimeout: timeout,
	}
}

// GetNearbyRestaurants retrieves nearby restaurant branches based on customer's location.
func (u *Usecase) GetNearbyRestaurants(ctx context.Context, costumerLocation domain.CostumerLocation, distance float64, limit int) (*[]domain.RestaurantBranch, error) {
	// Get nearby restaurant branches through the repository
	restaurantBranches, err := u.repository.GetNearbyRestaurants(ctx, costumerLocation, distance, limit)
	if err != nil {
		return nil, err
	}

	return restaurantBranches, nil
}
