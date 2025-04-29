FROM golang:1.24rc3-alpine AS builder

WORKDIR /app

# dependencies
COPY ["go.mod","go.sum", "./"]
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go mod download

# build
COPY . .
RUN swag init -g cmd/university/main.go --parseDependency --parseInternal
RUN go build -o university_system cmd/university/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/university_system .

RUN apk --no-cache add ca-certificates

EXPOSE 8080

CMD ["./university_system"]
