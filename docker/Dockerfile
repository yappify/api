# Stage 1: Build the Go binary
FROM golang:1.20-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/api

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/api

FROM scratch

WORKDIR /root/

COPY --from=builder /app/bin/api .

ENV ENVIRONMENT="development"
ENV PORT=8000
ENV DB_URL="postgresql://postgres:postgres@localhost:5432/db?sslmode=disable"

EXPOSE 8000

CMD ["./api"]
