.PHONY: start test test-cover

start:
	goreman -f ./air/Procfile start

test:
	@mkdir -p ./tmp/
	@set -a && . ./env/.env.test && \
		go test -v -count=1 -coverprofile=./tmp/test.out ./...

test-cover:
	@go tool cover -html=./tmp/test.out
