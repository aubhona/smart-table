FROM golang:1.23.2 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -extldflags '-static'" -o smart-table build/services/smart-table/src/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/smart-table .
COPY /configs /app/configs
CMD ["./smart-table"]
