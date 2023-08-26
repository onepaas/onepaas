docker.enter:
	docker exec -it onepaas-api /bin/sh

docker.up:
	docker compose up --build --remove-orphans -d

docker.stop:
	docker compose down

start: docker.up docker.enter
restart: docker.stop docker.up docker.enter
stop: docker.stop
