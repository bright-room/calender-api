DOCKER_RUN ?= docker compose --env-file .env --env-file .env.local

up:
	$(DOCKER_RUN) --profile default up -d

down:
	$(DOCKER_RUN) --profile default down -v

restart: down up

rmi:
	$(DOCKER_RUN) --profile all down --rmi all --remove-orphans

db/generate/%:
	$(DOCKER_RUN) --profile migrate run --remove-orphans migrate-generator $(@F)

db/up:
	$(DOCKER_RUN) --profile migrate run --remove-orphans migrate-up up

db/down:
	$(DOCKER_RUN) --profile migrate run --remove-orphans migrate-up down
