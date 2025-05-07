FROM golang:latest as builder

WORKDIR /app

COPY ../../go.mod ../../go.sum ./

RUN go mod download

COPY ../../cmd/ ./cmd/
COPY ../../internal/ ./internal
COPY ../../pkg/ ./pkg/

RUN mkdir -p ./build/package

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/package/server ./cmd/rest-api

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/build/package/server .

ENV ENVIRONMENT=production
ENV GIN_MODE=release

EXPOSE 8080

CMD ["./server"]
