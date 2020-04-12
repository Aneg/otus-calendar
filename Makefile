docker-up:
	docker-compose up -d
docker-down:
	docker-compose down
init-db:
	docker-compose exec postgres psql -U postgres -a -q -f  /data/sql/dll.sql