go-all: go-update go-generate go-all-tests
go-all-tests: go-lint go-unit-tests

go-dependencies:
	itbasis-builder dependencies

go-update:
	itbasis-builder update

go-generate:
	itbasis-builder generate

go-lint:
	itbasis-builder lint

go-unit-tests:
	itbasis-builder unit-test
