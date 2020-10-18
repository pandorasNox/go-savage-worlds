

PWD=$(shell pwd)


.PHONY: test
test:
	${PWD}/bin/exhaustive ./...
	go test ./...
	cd cmd/cli && cat ../../assets/tests/valid-character.yaml | go run .


.PHONY: test-sheet
test-sheet:
	cd cmd/cli && cat ../../assets/tests/valid-character.yaml | go run .


.PHONY: githooks
githooks:
	ln -s ../../githooks/pre-commit .git/hooks/pre-commit

