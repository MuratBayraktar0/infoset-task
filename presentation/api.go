package presentation

import (
	"context"
	"infoset-task/domain"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var app = fiber.New()

// Api handles HTTP requests and communicates with the business logic layer.
type Api struct {
	usecase domain.Usecase
}

// NewApi creates a new instance of an API.
func NewApi(usecase domain.Usecase) Api {
	return Api{
		usecase: usecase,
	}
}

// Router sets up the routes and handlers for the API.
func (u *Api) Router(ctx context.Context, app *fiber.App) *fiber.App {
	app.Get("/nearby-restaurants", func(c *fiber.Ctx) error {
		// Parse and validate the distance parameter
		distance, err := strconv.ParseFloat(c.Query("distance"), 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid distance"})
		}

		// Parse and validate the limit parameter
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid limit"})
		}

		// Parse and validate the customer's location from the request body
		var customerLocation domain.CostumerLocation
		if err := c.BodyParser(&customerLocation); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
		}

		// Get nearby restaurant branches using the use case
		restaurantBranches, err := u.usecase.GetNearbyRestaurants(ctx, customerLocation, distance, limit)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err})
		}

		// Return the nearby restaurant branches as JSON response
		return c.Status(fiber.StatusOK).JSON(restaurantBranches)
	})

	return app
}
