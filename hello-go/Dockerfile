FROM golang:latest AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o main main.go

# Use scratch to create a minimal container
#FROM scratch
FROM alpine:latest
COPY --from=builder /app/main ./

CMD ["./main"]
