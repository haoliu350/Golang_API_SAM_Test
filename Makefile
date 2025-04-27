.PHONY: build clean deploy local

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o bootstrap ./cmd/api
	zip -j function.zip bootstrap

clean:
	rm -rf bootstrap function.zip

deploy: build
	sam deploy --guided

local: build
	sam local start-api

test-get:
	curl -X GET n?name=Sam

test-post:
	curl -X POST http://localhost:3000/submit -H "Content-Type: application/json" -d '{"name":"Sam"}'