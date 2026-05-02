# How to Run Tests

## Unit Tests
Unit Tests verify the behaviour of individual encapsulated components (functions, types, interfaces etc).

### `absurd-iguana`
Here is how to build and run the unit tests for `absurd-iguana`.

```shell
docker build absurd-iguana-unit-tests
docker run absurd-iguana-unit-tests
```


## Integration Tests
Integration Tests test how 2 independent components (libraries, services, etc) interact. This project is composed of several interconnected services. Here is how you run the tests to verify their behaviour.

### `absurd-iguana` -> `dazzling-remora`
`absurd-iguana` sends data to our Kafka cluster `dazzling-remora`.
