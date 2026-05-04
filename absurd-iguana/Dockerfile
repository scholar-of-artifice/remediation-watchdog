# ----------------
# STAGE: base & dependency caching
# ----------------
# start with a large image that include the Go compiler and tools
FROM golang:1.26.2-trixie AS base
# install build depenedencies for confluent-kafka-go
# download the system tools.
#   - update refreshes the list of available software
#   - install downloads the following:
#       - build-esentials: includes gcc needed for CGO option
#       -  pkg-config: a helper tool that tells Go where find C libraries
#       - librdkafka-dev: the development files for Kafka, we need these to compile code with the Go app.
RUN apt-get update && apt-get install -y \
    build-essential \
    pkg-config \
    librdkafka-dev

# set the working directory in side the container to this
WORKDIR /app

# copy the go.mod and go.sum from here
# copy the dependency files first
COPY go.mod go.sum ./
# if the files are already there then this wont need to run
RUN go mod download

# copy the rest of the application code
COPY . .

# build the application with CGO enabled
# target the specific main.go location in the project
#   CGO_ENABLED: tells Go to allow C-interop. need this because confluent-kafka-go is a wrapper around a C library
#   GOOS: ensure building for the correct OS
# compile the source code into a single executable
RUN CGO_ENABLED=1 GOOS=linux go build -o absurd-iguana ./cmd/server/main.go

# ----------------
# STAGE: testing
# ----------------
FROM base AS test
# run all tests in the module
CMD ["go", "test", "-v", "-count=1", "./..."]

# ----------------
# STAGE: lab (not deployed just running it. probably could use a better name)
# ----------------
# switch to a slim-ish version of Debian.
# does not have go or the compiler
FROM debian:bookworm-slim as lab

# install a few things:
#   - librdkafka1: allow the compiled binary to talk to kafka when running
#   - ca-certificates: allow the app to make HTTPS connections
# then delete temporary files
RUN apt-get update && apt-get install -y \
    librdkafka1 \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*
# switch to this directory
WORKDIR /root/

# copy the binary from the builder stage
COPY --from=base /app/absurd-iguana .

# expose the port defined in the main.go file
EXPOSE 8080
# create a non-root user
RUN groupadd -r watchdog && useradd -r -g watchdog watchdog
# run the following commands as a non-root user
USER watchdog
# run the app once the container starts
CMD [ "./absurd-iguana" ]
