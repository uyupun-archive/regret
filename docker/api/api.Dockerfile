FROM golang:1.16.5-alpine3.13 AS builder
RUN mkdir /go/src/regret
WORKDIR /go/src/regret
COPY ./ .
ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64
RUN go build -o bin/regret cmd/main.go cmd/config.go cmd/router.go
RUN go build -o bin/migrate database/migrations/migrate.go
RUN go build -o bin/seeder database/seeds/*.go
RUN go build -ldflags '-X "main.ApiEndpoint=https://api.regret.uyupun.tech/api/v0/category"' -o bin/category tests/category/main.go
RUN go build -ldflags '-X "main.ApiEndpoint=https://api.regret.uyupun.tech/api/v0/inquiry"' -o bin/inquiry tests/inquiry/main.go

FROM alpine:3.13
WORKDIR /app/
COPY --from=builder /go/src/regret/bin/regret .
COPY --from=builder /go/src/regret/bin/migrate .
COPY --from=builder /go/src/regret/bin/seeder .
COPY --from=builder /go/src/regret/bin/category .
COPY --from=builder /go/src/regret/bin/inquiry .
COPY .env ./.env
COPY ./database/migrations/ ./database/migrations
ENTRYPOINT ["./regret"]
