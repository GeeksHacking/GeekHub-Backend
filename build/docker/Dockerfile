FROM golang:latest AS builder

WORKDIR /go/src/geekhub-backend/
COPY . .

RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/main cmd/main.go

FROM alpine:latest as certificates

RUN apk --no-cache add ca-certificates

FROM scratch as runner

COPY --from=builder /go/bin/main /go/bin/main
COPY --from=certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 5000

ENTRYPOINT ["/go/bin/main"]