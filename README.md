# Magellan

## Dependencies

- [pre-commit](https://pre-commit.com/)

## Getting Started

```shell
git clone https://github.com/clavinjune/magellan
cd magellan
pre-commit install # if you want to contribute only
go mod download
make generate
go run ./cmd/main.go proxy
# open http://localhost:8000 on your browser
```