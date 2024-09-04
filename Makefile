# run tests
test:
	@echo "🟢 Running tests..."

# run node
run:
	@echo "🏁 Running code..."

run-front:
	@echo "🏁 Running front code..."
	cd front-end && go run ./cmd/web

help:
	@echo "📖 Available commands:"
	@echo "  make run"
	@echo "  make test"
	@echo "  make help"
