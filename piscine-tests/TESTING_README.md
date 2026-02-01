# Piscine Go Testing Suite

This comprehensive testing suite helps you validate all your Piscine Go solutions, similar to LeetCode's testing system.

## 🚀 Quick Start

### Run All Tests
```bash
go test ./piscine -v
```

### Run Tests with Coverage
```bash
go test ./piscine -v -cover
```

### Run Tests with Detailed Coverage Report
```bash
go test ./piscine -v -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Run Specific Test
```bash
go test ./piscine -v -run TestFirstWord
```

### Run Benchmarks
```bash
go test ./piscine -bench=. -benchmem
```

## 📊 Test Coverage

The test suite covers the following functions:

### String Manipulation
- ✅ `FirstWord` - Extract first word from string
- ✅ `LastWord` - Extract last word from string
- ✅ `Fields` - Split string into words
- ✅ `IsCapitalized` - Check if words are capitalized
- ✅ `HashCode` - Hash string to printable characters

### Math Functions
- ✅ `Gcd` - Greatest common divisor

### Print Functions
- ✅ `PrintAlpha` - Print alphabet
- ✅ `PrintReverseAlpha` - Print reverse alphabet
- ✅ `PrintDigit` - Print digits 0-9
- ✅ `PrintComp` - Print combinations
- ✅ `PrintComb2` - Print two-number combinations
- ✅ `PrintNbr` - Print numbers
- ✅ `PrintMemory` - Print memory in hex format
- ✅ `IsNegative` - Check if number is negative

## 🎯 Test Cases Explained

### Example: FirstWord Tests
```go
TestFirstWord/simple_word          - Basic input: "hello world"
TestFirstWord/leading_spaces       - Input with spaces: "   hello world"
TestFirstWord/single_word          - Single word input
TestFirstWord/empty_string         - Empty input handling
TestFirstWord/only_spaces          - Whitespace-only input
TestFirstWord/tabs_and_spaces      - Mixed whitespace
TestFirstWord/multiple_words       - Multiple word handling
```

## 📝 Understanding Test Output

### Successful Test
```
=== RUN   TestFirstWord/simple_word
--- PASS: TestFirstWord/simple_word (0.00s)
```

### Failed Test
```
=== RUN   TestFirstWord/simple_word
    piscine_test.go:25: FirstWord("hello world") = "hello", want "hello\n"
--- FAIL: TestFirstWord/simple_word (0.00s)
```

## 🔧 Debugging Failed Tests

If a test fails:

1. **Read the error message** - It shows expected vs actual output
2. **Check edge cases** - Empty strings, spaces, special characters
3. **Review the test case** - Understand what scenario failed
4. **Fix your implementation** - Update your function
5. **Re-run the test** - Verify the fix

## 📈 Coverage Goals

Aim for:
- **100% test coverage** for all functions
- **All edge cases** covered (empty, null, extreme values)
- **Performance benchmarks** passing

## 🎓 Best Practices

1. **Run tests frequently** - After each change
2. **Check coverage** - Ensure all code paths tested
3. **Read test failures carefully** - They guide you to the issue
4. **Add more tests** - If you find edge cases not covered
5. **Use benchmarks** - To optimize performance

## 🚨 Common Issues

### Issue: Import error for z01
**Solution**: Make sure you have the z01 package installed:
```bash
go mod download
go mod tidy
```

### Issue: Tests timeout
**Solution**: Check for infinite loops in your implementations

### Issue: Coverage report doesn't open
**Solution**: Make sure you have a browser configured or manually open `coverage.out`

## 📚 Additional Resources

- [Go Testing Documentation](https://golang.org/pkg/testing/)
- [Table-Driven Tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)
- [Go Test Coverage](https://blog.golang.org/cover)

## 🎉 Tips for Success

- **Green tests = Confidence** - All tests passing means your code works!
- **Red tests = Learning opportunity** - Failed tests show what to fix
- **Yellow coverage = Room for improvement** - Add more test cases
- **Benchmarks = Performance insights** - Know how fast your code runs

---

Happy Testing! 🧪✨