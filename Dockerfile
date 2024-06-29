FROM golang:1.22.3

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY tests ./tests
COPY cmd ./cmd
COPY internal ./internal
COPY config ./config

EXPOSE 8082

RUN GOOS=linux go build -o app ./cmd/url-shortener/main.go

ENV CONFIG_PATH ./config/prod.yaml

CMD ["./app"]
