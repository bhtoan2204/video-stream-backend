run-gateway:
	@echo "Running Gateway"
	@cd gateway && go run ./cmd/main/main.go

run-user:
	@echo "Running User Service"
	@cd user && go run ./cmd/main/main.go

run-video:
	@echo "Running Video Service"
	@cd video && go run ./cmd/main/main.go

.PHONY: run-gateway run-user run-video