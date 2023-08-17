package repository

import (
	"context"
	"infoset-task/domain"
)

type Database interface {
	GetNearbyRestaurants(ctx context.Context, costumerLocation domain.CostumerLocation, distance float64, limit int) (*[]domain.RestaurantBranch, error)
}
