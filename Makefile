test:
	docker compose exec api go test ./internal/config -v 
lint:
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v2.1.2 golangci-lint run
start: docker compose up -d
ci: start lint test