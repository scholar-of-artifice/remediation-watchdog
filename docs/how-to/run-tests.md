# How to Run Tests

## Unit Tests
Unit Tests verify the behaviour of individual encapsulated components (functions, types, interfaces etc).

### `absurd-iguana`
Here is how to build and run the unit tests for `absurd-iguana`.

```shell
docker compose -f docker-compose.test.yml up --build --abort-on-container-exit
```

#### What does this do?
Build the test-stage of `absurd-iguana` Dockerfile.
Start an ephemeral Redis and Apache Kafka cluster.
Wait for the services to pass their healthchecks.
Executes the `go test` command.
Tears down all containers once the tests conclude.

### How can I run tests for specific packages?
```shell
docker compose -f docker-compose.test.yml up --build absurd-iguana-test go test ./internal/store/...
```



## Integration Tests
Integration Tests test how 2 independent components (libraries, services, etc) interact. This project is composed of several interconnected services. Here is how you run the tests to verify their behaviour.

### `absurd-iguana` -> `dazzling-remora`
`absurd-iguana` sends data to our Kafka cluster `dazzling-remora`.
