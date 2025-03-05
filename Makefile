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

.PHONY: run-gateway run-user run-video run-gateway-dev run-user-dev run-video-dev run-worker run-worker-dev