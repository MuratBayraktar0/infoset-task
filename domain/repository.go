package domain

import "context"

type Repository interface {
	GetNearbyRestaurants(ctx context.Context, costumerLocation CostumerLocation, distance float64, limit int) (*[]RestaurantBranch, error)
}
