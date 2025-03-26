FROM golang:1.24-alpine3.21 AS builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY internal internal
COPY docs docs
COPY main.go .
RUN go build -o ./build/server


FROM alpine:3.21

# Install time zone package
RUN apk add --update-cache tzdata

# Copy binary
WORKDIR /app
COPY --from=builder /app/build/server /app/server
COPY ./config /app/config
COPY ./docs /app/docs

EXPOSE 8082

# Run binary
ENTRYPOINT ["/app/server", "qworker"]
