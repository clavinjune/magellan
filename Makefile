include tools.mk
REVISION:=$(shell git rev-parse --short HEAD)
APP_NAME:=biko

.PHONY: buf/generate
buf/generate: buf/generate/proto buf/generate/swagger

.PHONY: buf/generate/proto
buf/generate/proto:
	@rm -rf api/proto && go run $(BUF) generate

.PHONY: buf/generate/swagger
buf/generate/swagger:
	@SWAGGER_UI_VERSION=$(SWAGGER_UI_VERSION) \
	./scripts/buf/generate-swagger.sh \
	&& ./scripts/buf/generate-swagger-handler.sh

.PHONY: build
build: wire
	@go build \
	-ldflags "-s -w -X main.Version=${REVISION}" \
	-o ${APP_NAME} cmd/main.go

.PHONY: docker/build
docker/build: wire
	@docker build \
	-f scripts/docker/Dockerfile \
	--build-arg REVISION=${REVISION} \
	-t ${APP_NAME}:${REVISION} .
	@docker tag ${APP_NAME}:${REVISION} ${APP_NAME}:latest

.PHONY: lint
lint:
	@go run $(BUF) lint
	@go vet ./...
	@go run $(GOIMPORTS) -w .
	@gofmt -w -s .
	@#go mod tidy
	@go run $(GOVULNCHECK) ./...

.PHONY: test
test:
	@go test -v -covermode=count -shuffle=on ./...

.PHONY: test/report
test/report:
	@go test -covermode=count -shuffle=on -coverprofile test-coverage.out -json ./... > test-report.json

.PHONY: test/cover
test/cover: test/report
	@go tool cover -html=test-coverage.out

.PHONY: wire
wire:
	@find . -name "*wire_gen.go" -delete
	@go run $(WIRE) ./...