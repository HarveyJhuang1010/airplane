gen:
	@echo "generate files"
	@sh ./scripts/generate.sh

migrate-local:
	@echo "migrate database"
	@sh ./migrate/rdb/scripts/migrate.sh local

migrate-dev:
	@echo "migrate database"
	@sh ./migrate/rdb/scripts/migrate.sh dev

migrate-prod:
	@echo "migrate database"
	@sh ./migrate/rdb/scripts/migrate.sh prod
