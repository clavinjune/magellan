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
