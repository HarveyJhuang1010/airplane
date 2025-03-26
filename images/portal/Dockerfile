FROM golang:1.22.1-alpine3.19 AS builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY internal internal
COPY docs docs
COPY main.go .
RUN go build -o ./build/server


FROM alpine:3.19

# Install time zone package
RUN apk add --update-cache tzdata

# Copy binary
workdir /app
COPY --from=builder /app/build/server /app/server
COPY ./config /app/config
COPY ./docs /app/docs

EXPOSE 8080

# Run binary
ENTRYPOINT ["/app/server", "portal"]
