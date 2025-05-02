# CRM API - Go & MongoDB Learning Project

This is a test project created to learn how to build RESTful APIs using Go (Golang) and MongoDB. The project implements a basic Customer Relationship Management (CRM) system with features like user authentication and lead management.

## Project Structure

```
crm-go/
├── modules/
│   ├── auth/     # Authentication module
│   └── lead/     # Lead management module
├── router/       # Route definitions
└── main.go       # Application entry point
```

## Features

- User Authentication (Register, Login)
- Lead Management (CRUD operations)
- RESTful API endpoints
- MongoDB integration
- Gin web framework
- JWT-based authentication
- Docker and Docker Compose support

## Prerequisites

- Docker and Docker Compose
- Git

## Setup Instructions

1. Clone the repository:
```bash
git clone <repository-url>
cd crm-go
```

2. Set up environment variables:
Create a `.env` file in the root directory following the format of the .env.example file

3. Run the application using Docker Compose:
```bash
docker-compose up --build
```

This will:
- Start a MongoDB container with authentication enabled
- Build and start the Go application container
- Set up the necessary networking between containers
- Mount persistent volumes for MongoDB data

The application will be available at `http://localhost:8080`

## API Endpoints

### Authentication
- POST `/api/auth/register` - Register a new user
- POST `/api/auth/login` - Login and get JWT token

### Leads
- POST `/api/leads` - Create a new lead
- GET `/api/leads/:id` - Get a specific lead

## Development

If you want to run the application locally without Docker:

1. Install Go 1.16 or higher
2. Install MongoDB
3. Start MongoDB locally
4. Run the application:
```bash
go run main.go
```

## Note

This is a learning project and it's meant to demonstrate basic concepts and serve as a starting point for learning Go and MongoDB. 