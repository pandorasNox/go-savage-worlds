

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


.PHONY: wasm
wasm:
	cd cmd/wasm && GOOS=js GOARCH=wasm go build \
	-o ../../web/static/generated/main.wasm
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./web/static/generated


.PHONY: serve
serve:
	$(MAKE) wasm
	cd cmd/webserver && go run . -dir ../../web

