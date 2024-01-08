DB_URL=postgres://postgres:postgres@localhost:5432/blog_realworld_db?sslmode=disable

migration_up:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migration_down:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc_gen:
	sqlc generate

proto-go-gen:
	echo "++++ buf generate ++++"
	rm -rf ./gen/go/* && buf generate && go mod tidy
	echo "++++ complete ++++"
.PHONY: proto-go-gen

clean:
	go clean

wire:
	cd internal/user/app/http_server && wire && cd -
.PHONY: wire

up_core_env:
	docker compose -f docker-compose-core.yaml up -d
.PHONY: up_core_env

down_core_env:
	docker compose -f docker-compose-core.yaml down
.PHONY: down_core_env