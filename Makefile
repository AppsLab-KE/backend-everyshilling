build-dev:
	docker-compose --verbose -f docker-compose-dev.yml -p backend-everyshilling build

run-dev:
	docker-compose --verbose  -f docker-compose-dev.yml -p backend-everyshilling up --build --remove-orphans

prune-dev:
	docker-compose -f docker-compose-dev.yml -p backend-everyshilling down --remove-orphans --volumes --rmi all

