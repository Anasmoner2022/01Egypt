# Piscine Go Makefile
# Makes running tests super easy!

.PHONY: test test-verbose coverage coverage-html bench clean help

# Default target
all: test

# Run all tests
test:
	@echo "🧪 Running tests..."
	@go test ./piscine -cover

# Run tests with verbose output
test-verbose:
	@echo "🧪 Running tests (verbose)..."
	@go test ./piscine -v -cover

# Run tests with coverage
coverage:
	@echo "📊 Generating coverage report..."
	@go test ./piscine -coverprofile=coverage.out
	@go tool cover -func=coverage.out

# Generate HTML coverage report
coverage-html: coverage
	@echo "🌐 Opening coverage in browser..."
	@go tool cover -html=coverage.out

# Run benchmarks
bench:
	@echo "🏎️  Running benchmarks..."
	@go test ./piscine -bench=. -benchmem

# Run specific test (usage: make test-one TEST=TestGcd)
test-one:
	@echo "🎯 Running $(TEST)..."
	@go test ./piscine -v -run $(TEST)

# Clean generated files
clean:
	@echo "🧹 Cleaning up..."
	@rm -f coverage.out
	@go clean

# Run tests and benchmarks
full: test bench

# Show test statistics
stats: coverage
	@echo ""
	@echo "📈 Test Statistics:"
	@go test ./piscine -v 2>&1 | grep -E "PASS|FAIL" | wc -l | xargs -I {} echo "  Total tests: {}"
	@go test ./piscine -cover 2>&1 | grep coverage | cut -d':' -f2

# Watch mode - run tests on file changes (requires entr)
watch:
	@echo "👀 Watching for changes..."
	@find . -name "*.go" | entr -c make test

# Help menu
help:
	@echo "📚 Available commands:"
	@echo ""
	@echo "  make test          - Run all tests"
	@echo "  make test-verbose  - Run tests with verbose output"
	@echo "  make coverage      - Generate coverage report"
	@echo "  make coverage-html - Open coverage report in browser"
	@echo "  make bench         - Run benchmarks"
	@echo "  make test-one      - Run specific test (TEST=TestName)"
	@echo "  make clean         - Remove generated files"
	@echo "  make full          - Run tests and benchmarks"
	@echo "  make stats         - Show test statistics"
	@echo "  make watch         - Watch files and auto-test (needs entr)"
	@echo "  make help          - Show this help message"
	@echo ""
	@echo "Examples:"
	@echo "  make test-one TEST=TestGcd"
	@echo "  make coverage-html"