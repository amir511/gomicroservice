FROM golang:1.20-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8080

# Run the application
CMD ["go", "run", "./cmd/main.go"]
