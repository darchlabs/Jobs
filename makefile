# load .env file
include .env
export $(shell sed 's/=.*//' .env)

BIN_FOLDER_PATH=bin
SERVICE_NAME=jobs
DOCKER_USER=darchlabs

rm:
	@echo "[rm] Removing..."
	@rm -rf bin

build:
	@echo "[building node]"
	@docker build -t darchlabs/jobs -f ./Dockerfile --progress tty .
	@echo "Build darchlabs-jobs docker image done ✔︎"

compile: rm
	@echo "[compile] Compiling..."
	@go build -o $(BIN_FOLDER_PATH)/$(SERVICE_NAME) cmd/$(SERVICE_NAME)/main.go

linux: rm
	@echo "[compile-linux] Compiling..."
	@GOOS=linux GOARCH=amd64 go build -o $(BIN_FOLDER_PATH)/$(SERVICE_NAME)-linux cmd/$(SERVICE_NAME)/main.go

dev:
	@echo "[dev] Running..."
	@go run cmd/$(SERVICE_NAME)/main.go

compose-up:
	@echo "[compose-dev]: Running docker compose dev mode..."
	@docker-compose -f docker-compose.yml up --build

compose-stop:
	@echo "[compose-dev]: Running docker compose dev mode..."
	@docker-compose -f docker-compose.yml down

docker-login:
	@echo "[docker] Login to docker..."
	@docker login -u $(DOCKER_USER) -p $(DOCKER_PASS)

docker: docker-login
	@echo "[docker] pushing $(DOCKER_USER)/$(SERVICE_NAME):$(VERSION)"
	@docker buildx create --use 
	@docker buildx build --platform linux/amd64,linux/arm64 --push -t $(DOCKER_USER)/$(SERVICE_NAME):$(VERSION) .