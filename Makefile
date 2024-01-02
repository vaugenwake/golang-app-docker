build:
	@echo "Building"
	@go build -o main cmd/api/main.go

run:
	@echo "Starting..."
	@go run cmd/api/main.go

watch:
	@if command -v air > /dev/null; then \
		air; \
		echo "Watching..."; \
	else \
		echo "Air is not installed"; \
		exit 1; \
	fi