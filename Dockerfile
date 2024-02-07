FROM golang:1.21.1 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o alertmanager_telegram .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/alertmanager_telegram .

EXPOSE 8080

CMD ["./alertmanager_telegram"]
