FROM golang:1.19.0-alpine3.16 as builder
WORKDIR /work

COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY api ./api
COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o microservice ./cmd/microservice/main.go

FROM alpine:3.16.2
WORKDIR /bin
COPY db/migrate /migrations
COPY --from=builder /work/microservice /bin/microservice
ENV GIN_MODE=release
CMD /bin/microservice