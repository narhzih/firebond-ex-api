FROM golang:alpine AS builder

# Copy base files and install dependency
COPY go.mod go.sum /firebond-ex-api/
WORKDIR /firebond-ex-api/
RUN go mod download

# Copy project files and build
COPY . /firebond-ex-api

# Create build
RUN CGO_ENABLED=0 GOOS=linux go build -a --installsuffix cgo -o bin/api ./cmd/api

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /firebond-ex-api/ /usr/firebond-ex-api

# start project files
ENV APP_ENV=prod
ENTRYPOINT ["/usr/firebond-ex-api/bin/api"]