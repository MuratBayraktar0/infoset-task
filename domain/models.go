package domain

// Creating our models; if the model we return to the user is different from the model we interact with in the database, we can make a distinction between DTO and ENTITY.
type RestaurantBranch struct {
	ID        string  `json:"id" db:"id"`
	Name      string  `json:"name" db:"name"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
}

type CostumerLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
