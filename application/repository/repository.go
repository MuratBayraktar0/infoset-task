package repository

import (
	"context"
	"infoset-task/domain"
)

// Repository manages database operations.
type Repository struct {
	db             Database
	contextTimeout context.Context
}

// NewRepository creates a new instance of a repository.
func NewRepository(timeout context.Context, database Database) *Repository {
	return &Repository{
		db:             database,
		contextTimeout: timeout,
	}
}

// GetNearbyRestaurants retrieves nearby restaurant branches from the database based on customer's location.
func (u *Repository) GetNearbyRestaurants(ctx context.Context, costumerLocation domain.CostumerLocation, distance float64, limit int) (*[]domain.RestaurantBranch, error) {
	// Retrieve restaurant branches near customer's location from the database
	restaurantBranches, err := u.db.GetNearbyRestaurants(ctx, costumerLocation, distance, limit)
	if err != nil {
		return nil, err
	}

	return restaurantBranches, nil
}
