.PHONY: build clean deploy

build:
	@echo "Building Go binary for ARM64..."
	cd src/go-lambda && GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -ldflags="-s -w" -o ../../bin/bootstrap main.go

clean:
	rm -rf ./bin .aws-sam

deploy: build
	sam deploy --guided