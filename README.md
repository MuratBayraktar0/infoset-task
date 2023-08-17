# Infoset Task

The project has been developed using the Test-Driven Development (TDD) methodology and structured according to the principles of Clean Architecture.

## How to Run

(To achieve this, you should modify the database settings in the .env file according to your own configuration, and ensure that you have a functional MySQL running locally.)

1. Clone the project folder to your computer.
2. Open a terminal and navigate to the project folder: `cd infoset-task`
3. Install the required dependencies: `go mod tidy`
4. To start the application: `go run main.go`
5. The application will run at [http://localhost:8080](http://localhost:8080).

## How to Run with Docker-Compose

1. Clone the project folder to your computer.
2. Open a terminal and navigate to the project folder: `cd infoset-task`
3. Install the required dependencies: `docker-compose up -d`
4. The application will run at [http://localhost:8080](http://localhost:8080).

## API Endpoints

**`GET /nearby-restaurants`**

Lists nearby restaurant branches.

**Example request:**

```bash
GET http://localhost:8080/nearby-restaurants?distance=10&limit=5
```

**Example request body:**

```json
{
  "latitude": 80,
  "longitude": 90
}
```

## Objective

In this project, our goal is to provide basic HTTP APIs for managing data. The project has been developed in compliance with Clean Architecture principles, ensuring tight control over dependencies within internal layers while facilitating extensibility outward.

## Used Technologies and Architecture

The project has been developed using the following technologies and architectural principles:

- **Programming Language:** Go
- **HTTP Framework:** Fiber (Version 2)
- **Database:** MySQL(MariaDB)
- **Test:** Testify

## Project Architecture

The project follows the Clean Architecture principles and is structured into the following layers:

1. **Application Layer:** This layer handles HTTP requests and manages the business logic. It communicates with the Domain Layer.
2. **Domain Layer:** This layer defines the business logic. Objects, business rules, and fundamental rules are defined here.
3. **Infrastructure Layer:** This layer provides access to the database and other external resources. Dependencies are managed in this layer.
4. **Presentation Layer:** This layer contains routers that handle HTTP requests and interface code.
