run-gateway:
	@echo "Running Gateway"
	@cd gateway && go run ./cmd/main/main.go

.PHONY: run-gateway