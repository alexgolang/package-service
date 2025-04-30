# Basic parameters
APP_NAME=package-service
DOCKER_IMAGE=package-service

# Main go file
MAIN_PATH=./cmd/main.go

.PHONY: build test docker docker-run docker-stop

# Build the binary
build:
	echo "Building..."
	go build -o bin/${APP_NAME} ${MAIN_PATH}

# Run tests
test:
	echo "Running tests..."
	go test ./...

# Build docker image
docker:
	echo "Building docker image..."
	docker build -t ${DOCKER_IMAGE} .

# Run in docker
docker-run: docker
	echo "Running docker container..."
	docker run -d --name ${APP_NAME} -p 8080:8080 ${DOCKER_IMAGE}

# Stop docker container
docker-stop:
	echo "Stopping docker container..."
	docker stop ${APP_NAME} || true
	docker rm ${APP_NAME} || true

# Default target
.DEFAULT_GOAL := test