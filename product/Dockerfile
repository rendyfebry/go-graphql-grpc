# Start by building the application.
FROM golang:1.14 as dependency
WORKDIR /usr/src/app

ENV GO111MODULE=on

COPY go.mod go.sum ./

RUN go mod download

# Builder Image
FROM dependency AS builder
WORKDIR /usr/src/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags="-s -w" -o /usr/src/app/bin/main /usr/src/app/cmd/app/main.go

# Now copy the executable into our base image.
FROM gcr.io/distroless/base

# Copy bin file
COPY --from=builder /usr/src/app/bin/main .

ENTRYPOINT ["/main"]
EXPOSE 50051
