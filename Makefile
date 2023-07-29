# TODO: Finish this
APP_NAME					:= attendant
COMPOSE						:= docker-compose.yml
COMPOSE_OVERRIDE	:= docker-compose.override.yml

build:
	@go build -o bin/${APP_NAME}

run: build
	@./bin/${APP_NAME}

up:
	@docker compose -f ${COMPOSE} -f ${COMPOSE_OVERRIDE} up --build

down:
	@docker compose -f ${COMPOSE} -f ${COMPOSE_OVERRIDE} down --remove-orphans

test:
	@go test -v ./...