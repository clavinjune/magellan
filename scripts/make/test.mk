.PHONY: test
test:
	@go test -v -covermode=count -shuffle=on ./...

.PHONY: test/report
test/report:
	@go test -covermode=count -shuffle=on -coverprofile test-coverage.out -json ./... > test-report.json

.PHONY: test/cover
test/cover: test/report
	@go tool cover -html=test-coverage.out