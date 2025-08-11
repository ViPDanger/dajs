APP_NAME=dajs
BUILD_DIR=/bin
GOAPI_DIR=/go-api
WEBAPP_DIR=/webapp
TGBOT_DIR=/telegram-bot
SRC=/cmd/main.go

.PHONY: all build docker webapp go-api tg-bot run clean cert

all: build docker run


cert:
	sudo certbot certonly --standalone -d dajs.vipdanger.keenetic.pro

build:
	mkdir -p .$(BUILD_DIR)
	go build -C .$(GOAPI_DIR) -o .$(BUILD_DIR)/$(APP_NAME)  .$(SRC)
	go build -C .$(WEBAPP_DIR) -o $(APP_NAME) ./main.go
	cp .$(GOAPI_DIR)/cmd/config.ini  .$(GOAPI_DIR)$(BUILD_DIR)/config.ini
	docker compose build
run:
	make -j3 webapp go-api tg-bot
docker:
	docker compose up -d
webapp:
	.$(WEBAPP_DIR)/$(APP_NAME)
go-api:
	.$(GOAPI_DIR)$(BUILD_DIR)/$(APP_NAME)
tg-bot:
	. .$(TGBOT_DIR)/venv/bin/activate && python .$(TGBOT_DIR)/main.py
clean:	
	rm -rf .$(GOAPI_DIR)/$(BUILD_DIR)
	rm -rf .$(WEBAPP_DIR)/$(APP_NAME)
	docker compose down -v
