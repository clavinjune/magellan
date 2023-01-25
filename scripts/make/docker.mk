.PHONY: docker/compose/up
docker/compose/up:
	@docker compose -f scripts/docker/docker-compose.yaml up

.PHONY: docker/compose/down
docker/compose/down:
	@docker compose -f scripts/docker/docker-compose.yaml down

.PHONY: docker/connect/postgresql
docker/connect/postgresql:
	@docker exec -it postgresql psql -U magellan -d magellan
