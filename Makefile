APP_NAME=adhd-tracker
APP_SRC=./src
APP_OUT=./bin/$(APP_NAME)
APP_PORT=8080

.PHONY: all build clean run

all: build

build:
	@mkdir -p bin
	go build -o $(APP_OUT) $(APP_SRC)

clean:
	rm -rf bin

run:
	@echo "Starting $(APP_NAME) on port $(APP_PORT)..."
	$(APP_OUT) -port $(APP_PORT)
