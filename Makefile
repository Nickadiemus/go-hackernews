ifneq (,$(wildcard ./.env))
    include .env
    export
endif

cmd-exists-%:
	@hash $(*) > /dev/null 2>&1 || \
		(echo "ERROR: '$(*)' must be installed and available on your PATH."; exit 1)

postgres: cmd-exists-postgres
	docker run --name $(DOCKER_CONTAINER) -p $(DB_PORT):$(DB_PORT) -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d $(DB):$(DOCKER_DB_IMAGE)

createdb: cmd-exists-createdb
	docker exec -it $(DOCKER_CONTAINER) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME);

dropdb: cmd-exists-dropdb
	docker exec -it $(DOCKER_CONTAINER) dropdb --username=$(DB_USER) $(DB_NAME);

migrateup: cmd-exists-migrate
	migrate -path ./pkg/db/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migratedown: cmd-exists-migrate
	migrate -path ./pkg/db/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown