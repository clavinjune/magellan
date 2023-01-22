.PHONY: build/bin
build/bin: generate
	@go build \
	-ldflags "-s -w -X main.version=${REVISION}" \
	-o ${APP_NAME} cmd/main.go

.PHONY: build/image
build/image: generate
	@docker build \
	-f scripts/docker/Dockerfile \
	--build-arg REVISION=${REVISION} \
	-t ${APP_NAME}:${REVISION} .
	@docker tag ${APP_NAME}:${REVISION} ${APP_NAME}:latest