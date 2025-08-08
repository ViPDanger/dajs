APP_NAME=dajs
BUILD_DIR=/bin
GOAPI_DIR=/go-api
SRC=/cmd/main.go

.PHONY: all build clean run

all: build run

build:
	mkdir -p .$(BUILD_DIR)
	go build -C .$(GOAPI_DIR) -o .$(BUILD_DIR)/$(APP_NAME)  .$(SRC)
	cp .$(GOAPI_DIR)/cmd/config.ini  .$(GOAPI_DIR)$(BUILD_DIR)/config.ini
	docker compose build
run:
	docker compose up -d  > /dev/null
	exec .$(GOAPI_DIR)$(BUILD_DIR)/$(APP_NAME)
clean:	
	rm -rf .$(GOAPI_DIR)/$(BUILD_DIR)
	docker compose down -v
