.PHONY: build clean deploy

build:
	go get ./...
	go mod vendor
	env GOOS=linux go build -ldflags="-s -w" -o bin/create_article api/create_article/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/create_author api/create_author/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/get_authors api/get_authors/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/get_articles api/get_articles/main.go

clean:
	rm -rf ./bin ./vendor

deploy: clean build
	sls deploy --verbose
