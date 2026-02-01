#!/bin/bash

# Piscine Go Test Runner
# This script runs all tests with nice formatting and reports

echo "🧪 Piscine Go Testing Suite"
echo "============================"
echo ""

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_color() {
    color=$1
    shift
    echo -e "${color}$@${NC}"
}

# Check if Go is installed
if ! command -v go &> /dev/null; then
    print_color $RED "❌ Go is not installed. Please install Go first."
    exit 1
fi

print_color $BLUE "📦 Checking dependencies..."
go mod download
go mod tidy

echo ""
print_color $BLUE "🏃 Running all tests..."
echo ""

# Run tests with verbose output
if go test ./piscine -v -cover; then
    echo ""
    print_color $GREEN "✅ All tests passed!"
else
    echo ""
    print_color $RED "❌ Some tests failed. Please review the output above."
    exit 1
fi

echo ""
print_color $BLUE "📊 Generating coverage report..."

# Generate coverage report
go test ./piscine -coverprofile=coverage.out > /dev/null 2>&1

if [ -f coverage.out ]; then
    coverage=$(go tool cover -func=coverage.out | grep total | awk '{print $3}')
    print_color $YELLOW "Coverage: $coverage"
    
    # Parse coverage percentage
    coverage_num=${coverage%\%}
    
    if (( $(echo "$coverage_num >= 90" | bc -l) )); then
        print_color $GREEN "🎉 Excellent coverage! ($coverage)"
    elif (( $(echo "$coverage_num >= 70" | bc -l) )); then
        print_color $YELLOW "⚠️  Good coverage, but can be improved ($coverage)"
    else
        print_color $RED "⚠️  Coverage is low ($coverage). Consider adding more tests."
    fi
    
    echo ""
    print_color $BLUE "To view detailed coverage report, run:"
    echo "  go tool cover -html=coverage.out"
fi

echo ""
print_color $BLUE "🏎️  Running benchmarks..."
echo ""

go test ./piscine -bench=. -benchmem -run=^$

echo ""
print_color $GREEN "✨ Testing complete!"
echo ""
print_color $YELLOW "Quick commands:"
echo "  • Run specific test:  go test ./piscine -v -run TestGcd"
echo "  • See coverage HTML:  go tool cover -html=coverage.out"
echo "  • Run benchmarks:     go test ./piscine -bench=."
echo ""