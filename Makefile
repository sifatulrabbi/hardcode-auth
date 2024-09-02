run:
	GO_ENV=development go run ./cmd/api/main.go

build:
	rm -rf ./build
	mkdir build
	go build ./cmd/api/main.go -o ./build/hardcodeauth

.PHONY: run, build
