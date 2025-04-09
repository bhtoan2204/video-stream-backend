run-gateway:
	@echo "Running Gateway"
	@cd apps/gateway && go run ./cmd/main/main.go

run-gateway-dev:
	@echo "Running Gateway"
	@cd apps/gateway && go run ./cmd/main/main.go -dev

run-user:
	@echo "Running User Service"
	@cd apps/user && go run ./cmd/main/main.go

run-user-dev:
	@echo "Running User Service"
	@cd apps/user && go run ./cmd/main/main.go -dev

run-video:
	@echo "Running Video Service"
	@cd apps/video && go run ./cmd/main/main.go

run-video-dev:
	@echo "Running Video Service"
	@cd apps/video && go run ./cmd/main/main.go -dev

run-worker:
	@echo "Running Worker"
	@cd apps/worker && go run ./cmd/main/main.go

run-comment:
	@echo "Running Comment Service"
	@cd apps/comment && go run ./cmd/main/main.go

run-worker-dev:
	@echo "Running Worker"
	@cd apps/worker && go run ./cmd/main/main.go -dev

test-gateway:
	@echo "Running Gateway Tests"
	@cd apps/gateway && go test -v ./...

test-user:
	@echo "Running User Service Tests"
	@cd apps/user && go test -v ./...

test-video:
	@echo "Running Video Service Tests"
	@cd apps/video && go test -v ./...

test-comment:
	@echo "Running Comment Service Tests"
	@cd apps/comment && go test -v ./...

build-gateway:
	@echo "Building Gateway"
	@cd apps/gateway && go build -o gateway-service ./cmd/main/main.go

build-user:
	@echo "Building User Service"
	@cd apps/user && go build -o user-service ./cmd/main/main.go

build-video:
	@echo "Building Video Service"
	@cd apps/video && go build -o video-service ./cmd/main/main.go

run-k6:
	@echo "Running K6"
	@cd testing && k6 run login-test.js

generate-swagger:
	@echo "Generating Swagger Documentation"
	@cd apps/gateway/cmd/main && swag init --parseDependency -g main.go -d .,../../internal/modules --output docs/swagger --parseDepth 2 --exclude "test,mocks,docs"

.PHONY: \
	run-gateway \
	run-user \
	run-video \
	run-comment \
	run-gateway-dev \
	run-user-dev \
	run-video-dev \
	run-worker \
	run-worker-dev \
	test-gateway \
	test-user \
	test-video \
	test-comment \
	build-gateway \
	build-user \
	build-video \
	run-k6
	generate-swagger