

.PHONY: test
test:
	go test ./...
	bash githooks/pre-commit
	cat assets/tests/valid-character.yaml | go run .


.PHONY: githooks
githooks:
	ln -s githooks/pre-commit .git/hooks/pre-commit

