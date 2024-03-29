all: deps build

deps:
	go run build.go setup
	godep restore
	npm install

build:
	go run build.go build
	npm run build

test:
	godep go test -v ./pkg/...
	npm run test

run:
	./bin/adhdtracker-server