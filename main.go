package main

import (
	"context"
	"infoset-task/application/repository"
	"infoset-task/application/usecase"
	"infoset-task/infrastructure"
	"infoset-task/presentation"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDBName := os.Getenv("DB_DBNAME")
	ctx := context.Background()
	mysql, err := infrastructure.NewMySQL(dbHost, dbPort, dbUsername, dbPassword, dbDBName)
	if err != nil {
		log.Fatal(err)
	}

	defer mysql.Close()
	repository := repository.NewRepository(ctx, mysql)
	usecase := usecase.NewUsecase(ctx, repository)
	api := presentation.NewApi(usecase)
	app := api.Router(ctx, fiber.New())
	log.Fatal(app.Listen(":8080"))
}
