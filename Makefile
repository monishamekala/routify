# Variables
APP_NAME = routify
API_DIR = ./cmd/api
PROTO_DIR = ./proto
PROTO_OUT = ./proto
DOCKER_API_IMAGE = routify-api
DOCKER_UI_IMAGE = routify-ui

# Default target
.PHONY: all
all: build

# Build Go API
.PHONY: build
build:
	go build -o bin/$(APP_NAME) $(API_DIR)

# Run API locally
.PHONY: run
run:
	go run $(API_DIR)/main.go

# Generate gRPC code
.PHONY: proto
proto:
	protoc --go_out=$(PROTO_OUT) --go-grpc_out=$(PROTO_OUT) $(PROTO_DIR)/service.proto

# Clean build artifacts
.PHONY: clean
clean:
	rm -rf bin/
	go clean

# Run unit tests
.PHONY: test
test:
	go test ./... -v

# Docker build
.PHONY: docker-build
docker-build:
	docker build -f docker/api.Dockerfile -t $(DOCKER_API_IMAGE) .
	docker build -f docker/ui.Dockerfile -t $(DOCKER_UI_IMAGE) .

# Docker Compose up
.PHONY: up
up:
	docker-compose up --build

# Docker Compose down
.PHONY: down
down:
	docker-compose down

# Kubernetes deploy
.PHONY: deploy
deploy:
	kubectl apply -f deployments/

# Kubernetes teardown
.PHONY: undeploy
undeploy:
	kubectl delete -f deployments/
