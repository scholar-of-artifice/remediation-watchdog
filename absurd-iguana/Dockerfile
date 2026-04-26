# STAGE: base
FROM golang:1.26.2-bookworm AS base
# install build depenedencies for confluent-kafka-go
RUN apt-get update && apt-get install -y \
    build-essential \
    pkg-config \
    librdkafka-dev

WORKDIR /app

# copy the go.mod and go.sum from here
COPY go.mod go.sum ./
RUN go mod download

# copy the rest of the application code
COPY . .

# build the application with CGO enabled
# target the specific main.go location in the project
RUN CGO_ENABLED=1 GOOS=linux go build -o absurd-iguana ./cmd/server/main.go

# STAGE: runtime
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    librdkafka1 \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /root/

# copy the binary from the builder stage
COPY --from=base /app/absurd-iguana .

# expose the port defined in the main.go file
EXPOSE 8080

# run the app
CMD [ "./absurd-iguana" ]
