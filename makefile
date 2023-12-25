up:
	docker-compose build && docker-compose up
down:
	docker-compose stop && docker-compose down
rerun: down up
