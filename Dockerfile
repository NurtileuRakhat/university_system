FROM golang:1.24rc3-alpine

WORKDIR /app

# dependencies
COPY ["go.mod","go.sum", "./"]
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go mod download


#build
COPY . .
RUN swag init -g cmd/university/main.go --parseDependency --parseInternal

EXPOSE 8080

RUN go build -o university_system cmd/university/main.go
CMD ["./university_system"]


