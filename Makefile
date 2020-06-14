

.PHONY: test
test:
	go test ./...
	cat assets/tests/valid-character.yaml | go run .

