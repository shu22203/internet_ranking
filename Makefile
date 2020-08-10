fmt:
	docker-compose run --rm api go fmt

test:
	docker-compose run --rm api go test ./...
