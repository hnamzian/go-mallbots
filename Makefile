build:
	docker build . -f docker/Dockerfile -t mallbots
up:
	docker-compose up -d
down:
	docker-compose down -v
run:
	go run ./cmd/