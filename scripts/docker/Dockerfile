FROM golang:1.19-alpine3.17 as builder
LABEL tag="builder"

ARG REVISION
ENV REVISION=$REVISION

ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build \
    -ldflags "-s -w -X main.Version=$REVISION" \
    -o /bin/app ./cmd/main.go

FROM gcr.io/distroless/base:latest

WORKDIR /bin
COPY --from=builder /bin .

ENTRYPOINT ["/bin/app"]
