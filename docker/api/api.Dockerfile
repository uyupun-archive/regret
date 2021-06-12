FROM golang:1.16.5-alpine3.13 AS builder
RUN mkdir /go/src/regret
WORKDIR /go/src/regret
COPY ./ .
ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64
RUN go build -o bin/regret cmd/main.go cmd/config.go cmd/router.go
RUN go build -o bin/migrate database/migrations/migrate.go

FROM alpine:3.13
WORKDIR /app/
COPY --from=builder /go/src/regret/bin/regret .
COPY --from=builder /go/src/regret/bin/migrate .
COPY .env ./.env
COPY ./database/migrations/ ./database/migrations
ENTRYPOINT ["./regret"]
