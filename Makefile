# help: list all defined make commands
help:
	@echo "[Yappify] Defined make commands:"
	@echo "1. 'make run': run the api"
	@echo "2. 'make help': list all defined make commands"
	@echo "3. 'make tv': tidy go dependencies and store in /vendor"
	@echo "4. 'make db-up': start postgres service in a docker container"
	@echo "5. 'make db-down': stop postgres service"
	@echo "6. 'make db-psql': connect to database via psql CLI utility"
	@echo "7. 'make db-migrate-up': run goose database up migration"
	@echo "8. 'make db-migrate-down': run goose database down migration"

# run: run the api
run:
	@echo "[Yappify] Running Yappify API..."
	go run ./cmd/api

# tv: clean up unused dependencies and add new dependencies to /vendor
tv:
	@echo "[Yappify] Cleaning up unused dependencies and storing new dependencies in /vendor..."
	go mod tidy && go mod vendor
	@echo "[Yappify] Operation successfully completed!"

# up-db: start postgres service in a docker container
db-up:
	@echo "[Yappify] Starting postgres service in the background..."
	docker-compose -f ./docker/docker-compose.yml up -d
	@echo "[Yappify] Successfully started postgres service!"

# down-db: stop postgres service
db-down:
	@echo "[Yappify] Stopping postgres service..."
	docker-compose -f ./docker/docker-compose.yml down
	@echo "[Yappify] Successfully stopped postgres service!"

# db-psql: access containerised postgres database via psql CLI utility
db-psql:
	@echo "[Yappify] Accessing postgres database via psql utility..."
	psql -h localhost -U postgres -d postgres

# db-migrate-up: run goose database up migration
db-migrate-up:
	@echo "[Yappify] Running goose up migration..."
	chmod +x scripts/run_goose.sh
	@./scripts/run_goose.sh up
	@echo "[Yappify] Successfully ran goose up migration!"

# db-migrate-down: run goose database down migration
db-migrate-down:
	@echo "[Yappify] Running goose down migration..."
	chmod +x scripts/run_goose.sh
	@./scripts/run_goose.sh down
	@echo "[Yappify] Successfully ran goose down migration!"