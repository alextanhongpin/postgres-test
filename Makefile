test-integration:
	@go test -v --tags=integration

test-unit:
	@go test -v --tags=unit

test-all:
	@go test -v --tags="unit integration"
