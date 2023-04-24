build-dev:
	docker-compose -f docker-compose-dev.yml -p backend-everyshilling build

run-dev:
	docker-compose -f docker-compose-dev.yml -p backend-everyshilling up

run-dev:
        docker-compose -f docker-compose.dev.yml up --build
