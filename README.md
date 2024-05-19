# Dating Application Backend Service

## Introduction

Welcome to the Dating Application Backend Service! This service provides the backend functionality for a dating application, including user management, profile creation, swiping, matching, and chatting features.

## Structure of the Service
Implemented Clean Architecture design patterns.
The service is structured using the following components:

- **Models**: Define the data structures used in the application, such as User, Profile, Swipe, MatchRoom, and Message.
  
- **Repositories**: Provide the interface and implementations for interacting with the database. There are repositories for users, profiles, swipes, matches, and messages.
  
- **Use Cases**: Implement the business logic of the application. Use cases orchestrate interactions between repositories to perform operations like user registration, profile viewing, swiping, matching, and messaging.
  
- **Handlers and Routes**: Define the HTTP handlers and routes for handling incoming requests. Handlers use the appropriate use cases to fulfill requests and return responses.

- **Utilities**: Contains utility functions used across the application, such as hashing passwords and handling JWT tokens.

**Project Structure**:
```
dating-app/
├── cmd/
│   └── server/
│       └── main.go
├── config/
├── app/
│   └── models/
│   └── repository/
│   └── usecase/
│   └── handler/
│   └── routes/
│   └── middleware/
│   └── scheduler/
│   └── utils/
├── tests/
├── go.mod
├── go.sum
├── .env
```

## Instructions to Run the Service

To run the Dating Application Backend Service, follow these steps:

1. **Clone the Repository**: Clone the repository to your local machine using the following command:
   ```
   git clone <repository-url>
   ```
2. **Install Dependencies**: Navigate to the project directory and install the dependencies using your preferred package manager. For example, with Go modules:
   ```
   cd existing_repo
   go get -u ./... or go mod tidy
   ```
3. **Set Up the Environment Variables**: Configure your env variables for database connection details in the `.env` file. If there is no .env file, copy from `.env.example`.

4. **Run the Service**: Start the service by running the main application file by running:
   ```golang
   go run main.go
   ```
5. **Access the API**: Once the service is running, you can access the API endpoints using your preferred API client (e.g., Postman). The API endpoints are defined in the routes file and include features such as user registration, profile creation, swiping, matching, and messaging.

6. **Testing**: You can run unit tests for the service to ensure its correctness by running:
   ```golang
   go test -v ./tests
   or
   go test ./...
   ```

## Additional Notes

- **Tech Stack**: The service is built using Go programming language with the Gin framework for routing and GORM for ORM. PostgreSQL is used as the database, and Pusher Channels are utilized for real-time messaging.

- **Authentication**: Authentication is handled using JWT tokens, and authorization checks are implemented where necessary to ensure that only authenticated users can access certain endpoints.
