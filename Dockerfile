FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/app/main.go

RUN go build -o /myapp ./cmd/app

COPY wait-for.sh .

EXPOSE 8080

CMD ["/app/wait-for.sh", "db", "/myapp"]
