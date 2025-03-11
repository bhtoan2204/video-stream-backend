run-gateway:
	@echo "Running Gateway"
	@cd gateway && go run ./cmd/main/main.go

run-gateway-dev:
	@echo "Running Gateway"
	@cd gateway && go run ./cmd/main/main.go -dev

run-user:
	@echo "Running User Service"
	@cd user && go run ./cmd/main/main.go

run-user-dev:
	@echo "Running User Service"
	@cd user && go run ./cmd/main/main.go -dev

run-video:
	@echo "Running Video Service"
	@cd video && go run ./cmd/main/main.go

run-video-dev:
	@echo "Running Video Service"
	@cd video && go run ./cmd/main/main.go -dev

run-worker:
	@echo "Running Worker"
	@cd worker && go run ./cmd/main/main.go

run-worker-dev:
	@echo "Running Worker"
	@cd worker && go run ./cmd/main/main.go -dev

test-gateway:
	@echo "Running Gateway Tests"
	@cd gateway && go test -v ./...

test-user:
	@echo "Running User Service Tests"
	@cd user && go test -v ./...

test-video:
	@echo "Running Video Service Tests"
	@cd video && go test -v ./...

build-gateway:
	@echo "Building Gateway"
	@cd gateway && go build -o gateway-service ./cmd/main/main.go

build-user:
	@echo "Building User Service"
	@cd user && go build -o user-service ./cmd/main/main.go

build-video:
	@echo "Building Video Service"
	@cd video && go build -o video-service ./cmd/main/main.go

.PHONY: run-gateway run-user run-video run-gateway-dev run-user-dev run-video-dev run-worker run-worker-dev test-gateway test-user test-video build-gateway build-user build-video