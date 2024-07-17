FROM golang:1.22.3

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the application binary and place it in the /app directory
RUN go build -o ./message-service ./cmd/main.go

# Set execute permissions for the binary
RUN chmod +x ./message-service

EXPOSE 8080

CMD ["./message-service"]
