build-dev:
	docker-compose -f docker-compose-dev.yml -p backend-everyshilling build

run-dev:
	docker-compose -f docker-compose-dev.yml -p backend-everyshilling up --build

prune-dev:
	docker-compose -f docker-compose-dev.yml -p backend-everyshilling down --remove-orphans

