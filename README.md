# postgres-test
Demonstrates testing Postgres with Docker.


## Approach

1. Unit and integration tests are isolated with the build tags `unit` and `integration` respectively:

```go
// For unit test.
// +build unit

// For integration test.
// +build integration
```

2. Postgres's Docker container is only initialized in integration tests.
3. To run the tests:
```bash
# Running integration tests only.
$ go test -v --tags=integration

# Running unit tests only.
$ go test -v --tags=unit

# Running both unit and integration tests.
$ go test -v --tags="unit integration"
```
