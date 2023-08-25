build:
	docker build . -f docker/Dockerfile -t mallbots
up:
	docker-compose up -d
down:
	docker-compose down -v
run:
	go run ./cmd/
pbgen:
	@protoc \
	--go_out=. --go_opt=paths=source_relative \
 	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./*/*pb/*.proto