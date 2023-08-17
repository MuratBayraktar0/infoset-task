package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"infoset-task/domain"

	_ "github.com/go-sql-driver/mysql"
)

// MySQL manages mysql operations.
type MySQL struct {
	db *sql.DB
}

// NewMySQL creates a new instance of a MySQL.
func NewMySQL(dbHost, dbPort, dbUsername, dbPassword, dbDBName string) (*MySQL, error) {
	db, err := initialDb(dbHost, dbPort, dbUsername, dbPassword, dbDBName)
	if err != nil {
		return nil, err
	}

	return &MySQL{
		db: db,
	}, nil
}

// initialDb initializes and sets up the database connection.
func initialDb(dbHost, dbPort, dbUsername, dbPassword, dbDBName string) (*sql.DB, error) {
	// Open a connection to create the initial database (if it doesn't exist)
	db, err := sql.Open("mysql", dsn(dbHost, dbPort, dbUsername, dbPassword, ""))
	if err != nil {
		return nil, err
	}

	// Create the "restaurant" database (if it doesn't exist)
	_, err = db.Exec(`CREATE DATABASE IF NOT EXISTS restaurant`)
	if err != nil {
		return nil, err
	}

	// Open a connection to the specified database
	db, err = sql.Open("mysql", dsn(dbHost, dbPort, dbUsername, dbPassword, dbDBName))
	if err != nil {
		return nil, err
	}

	// Create the "restaurant_branches" table (if it doesn't exist)
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS restaurant_branches (
			id VARCHAR(100) PRIMARY KEY,
			name VARCHAR(100),
			latitude DOUBLE,
			longitude DOUBLE
		);
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// dsn returns the data source name for the database connection.
func dsn(dbHost, dbPort, dbUsername, dbPassword, dbDBName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbDBName)
}

// Close closes the database connection.
func (m *MySQL) Close() {
	if m.db != nil {
		m.db.Close()
	}
}

// GetNearbyRestaurants retrieves nearby restaurants within a specified distance.
func (m *MySQL) GetNearbyRestaurants(ctx context.Context, costumerLocation domain.CostumerLocation, distance float64, limit int) (*[]domain.RestaurantBranch, error) {
	// SQL query using Haversine formula to select nearby restaurant branches within the given distance
	query := `
		SELECT id, name, latitude, longitude
		FROM restaurant_branches
		WHERE SQRT(POW(69.1 * (latitude - ?), 2) + POW(69.1 * (? - longitude) * COS(latitude / 57.3), 2)) <= ?
		ORDER BY SQRT(POW(69.1 * (latitude - ?), 2) + POW(69.1 * (? - longitude) * COS(latitude / 57.3), 2))
		LIMIT ?
	`

	// Execute the SQL query with the provided parameters
	rows, err := m.db.QueryContext(ctx, query, costumerLocation.Latitude, costumerLocation.Longitude, distance, costumerLocation.Latitude, costumerLocation.Longitude, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var restaurantBranches []domain.RestaurantBranch
	for rows.Next() {
		var restaurantBranch domain.RestaurantBranch
		err := rows.Scan(&restaurantBranch.ID, &restaurantBranch.Name, &restaurantBranch.Latitude, &restaurantBranch.Longitude)
		if err != nil {
			return nil, err
		}
		restaurantBranches = append(restaurantBranches, restaurantBranch)
	}

	return &restaurantBranches, nil
}
