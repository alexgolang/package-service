# Package Size Calculator Service

A Go-based microservice that calculates optimal package sizes for given target quantities. The service implements an efficient algorithm to determine the most cost-effective combination of available package sizes to fulfill any order quantity.

## Features

- RESTful API for package size calculations
- Efficient dynamic programming algorithm for optimal package size determination
- Configurable package sizes via environment variables
- Docker support for easy deployment
- Structured logging
- Comprehensive test coverage
- Makefile for easy development and deployment

## Prerequisites

- Go 1.24 or higher
- Docker (optional, for containerized deployment)
- Make (for using Makefile commands)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/alexgolang/package-task.git
cd package-task
```

2. Install dependencies:
```bash
go mod download
```

## Configuration

The service can be configured using environment variables:

| Variable | Description | Default Value |
|----------|-------------|---------------|
| HTTP_SERVER_PORT | Port for the HTTP server | 8080 |
| PACKAGE_SIZES | Comma-separated list of available package sizes | 250,500,1000,2000,5000 |

### Custom Package Sizes

You can modify package sizes without changing the code by setting the `PACKAGE_SIZES` environment variable:

```bash
# Example with custom sizes
export PACKAGE_SIZES=200,400,800,1600,3200
```

## Running the Service

### Using Make Commands

```bash
# Run tests
make test

# Build and run in Docker
make docker-run

# Stop Docker container
make docker-stop

# Build binary locally
make build
```

### Local Development

```bash
go run cmd/main.go
```

### Using Docker Manually

```bash
# Build the image
docker build -t package-service .

# Run with default package sizes
docker run -p 8080:8080 package-service

# Run with custom package sizes
docker run -p 8080:8080 -e PACKAGE_SIZES=200,400,800,1600,3200 package-service
```

## API Documentation

### Get Package Sizes

Calculates the optimal combination of package sizes for a given target quantity.

**Endpoint:** `GET /package?size={target_size}`

**Parameters:**
- `size` (required): The target quantity to be packaged

**Example Request:**
```bash
curl "http://localhost:8080/package?size=501"
```

**Example Response:**
```json
{
    "500": 1,
    "250": 1
}
```

**Response Format:**
- Keys: The package size used
- Values: The number of packages of that size needed

**Error Responses:**

- `400 Bad Request`: Invalid size parameter
- `405 Method Not Allowed`: Wrong HTTP method used

## Algorithm

The service uses a dynamic programming approach to find the optimal combination of packages:

1. Finds the minimum number of packages needed
2. Ensures the sum of packages is greater than or equal to the target size
3. Optimizes for the smallest possible total

Example scenarios:
- Target: 1 → Result: {250: 1}
- Target: 251 → Result: {500: 1}
- Target: 501 → Result: {500: 1, 250: 1}
- Target: 12001 → Result: {5000: 2, 2000: 1, 250: 1}

