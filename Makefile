.PHONY: build clean deploy

build:
	go get ./...
	go mod vendor
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello api/hello/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/world api/world/main.go

clean:
	rm -rf ./bin ./vendor

deploy: clean build
	sls deploy --verbose
