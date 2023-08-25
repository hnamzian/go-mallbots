build:
	docker build . -f docker/Dockerfile -t mallbots
up:
	docker-compose up -d
down:
	docker-compose down -v
run:
	go run ./cmd/
generate:
	@echo running code generation
	@go generate ./...
	@echo done
