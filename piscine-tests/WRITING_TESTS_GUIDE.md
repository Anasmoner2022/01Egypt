# How to Write Your Own Tests

This guide will help you write custom tests for your Piscine functions.

## 🎯 Test Structure

Every Go test follows this basic structure:

```go
func TestFunctionName(t *testing.T) {
    // Arrange: Set up test data
    input := "some input"
    expected := "expected output"
    
    // Act: Call your function
    result := FunctionName(input)
    
    // Assert: Check if result matches expected
    if result != expected {
        t.Errorf("FunctionName(%q) = %q, want %q", input, result, expected)
    }
}
```

## 📋 Table-Driven Tests (Recommended)

For testing multiple cases, use table-driven tests:

```go
func TestMyFunction(t *testing.T) {
    tests := []struct {
        name     string    // Test case name
        input    string    // Input value
        expected string    // Expected output
    }{
        {"case 1", "input1", "output1"},
        {"case 2", "input2", "output2"},
        {"edge case", "", ""},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := MyFunction(tt.input)
            if result != tt.expected {
                t.Errorf("MyFunction(%q) = %q, want %q", 
                    tt.input, result, tt.expected)
            }
        })
    }
}
```

## 🔍 Testing Print Functions

For functions that use `z01.PrintRune`, use the `captureOutput` helper:

```go
func TestMyPrintFunction(t *testing.T) {
    expected := "hello world\n"
    
    result := captureOutput(func() {
        MyPrintFunction()
    })
    
    if result != expected {
        t.Errorf("MyPrintFunction() = %q, want %q", result, expected)
    }
}
```

## ✅ Test Checklist

When writing tests, make sure to cover:

- [ ] **Happy path** - Normal, expected inputs
- [ ] **Edge cases** - Empty strings, zeros, boundaries
- [ ] **Invalid inputs** - Negative numbers, nil values
- [ ] **Large inputs** - Test with big numbers/strings
- [ ] **Special characters** - Test with unicode, newlines, tabs
- [ ] **Boundary values** - Min/max integers, empty arrays

## 📝 Example: Testing a New Function

Let's say you wrote a new function `CountWords`:

```go
// piscine/countwords.go
package piscine

func CountWords(s string) int {
    words := Fields(s)
    return len(words)
}
```

Here's how to test it:

```go
// piscine_test.go
func TestCountWords(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected int
    }{
        {"no words", "", 0},
        {"one word", "hello", 1},
        {"two words", "hello world", 2},
        {"leading spaces", "  hello world", 2},
        {"trailing spaces", "hello world  ", 2},
        {"multiple spaces", "hello    world", 2},
        {"tabs", "hello\tworld", 2},
        {"newlines", "hello\nworld\nfoo", 3},
        {"mixed whitespace", "  hello\t\nworld  ", 2},
        {"only spaces", "   ", 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := CountWords(tt.input)
            if result != tt.expected {
                t.Errorf("CountWords(%q) = %d, want %d", 
                    tt.input, result, tt.expected)
            }
        })
    }
}
```

## 🎨 Testing Best Practices

### 1. **Descriptive Test Names**
```go
// Good ✅
{"empty_string_returns_zero", "", 0}
{"single_word_returns_one", "hello", 1}

// Bad ❌
{"test1", "", 0}
{"test2", "hello", 1}
```

### 2. **Clear Error Messages**
```go
// Good ✅
t.Errorf("CountWords(%q) = %d, want %d", input, result, expected)

// Bad ❌
t.Errorf("wrong result")
```

### 3. **Test One Thing at a Time**
```go
// Good ✅
func TestAdd(t *testing.T) { /* tests addition */ }
func TestSubtract(t *testing.T) { /* tests subtraction */ }

// Bad ❌
func TestMath(t *testing.T) { /* tests everything */ }
```

### 4. **Use Subtests for Organization**
```go
t.Run("with valid input", func(t *testing.T) {
    // test cases
})

t.Run("with invalid input", func(t *testing.T) {
    // test cases
})
```

## 🐛 Debugging Failed Tests

When a test fails:

1. **Read the error message carefully**
   ```
   CountWords("hello  world") = 3, want 2
   ```
   This tells you exactly what went wrong.

2. **Add debug prints if needed**
   ```go
   t.Logf("Input: %q", input)
   t.Logf("Result: %v", result)
   t.Logf("Expected: %v", expected)
   ```

3. **Run just that one test**
   ```bash
   go test ./piscine -v -run TestCountWords
   ```

## 📊 Benchmarking Functions

To measure performance:

```go
func BenchmarkCountWords(b *testing.B) {
    input := "hello world foo bar baz"
    
    for i := 0; i < b.N; i++ {
        CountWords(input)
    }
}
```

Run with:
```bash
go test ./piscine -bench=BenchmarkCountWords
```

## 🎯 Advanced: Testing with Different Types

### Testing Arrays/Slices
```go
func TestArrayFunction(t *testing.T) {
    input := [5]int{1, 2, 3, 4, 5}
    expected := [5]int{2, 4, 6, 8, 10}
    
    result := DoubleArray(input)
    
    if result != expected {
        t.Errorf("DoubleArray(%v) = %v, want %v", 
            input, result, expected)
    }
}
```

### Testing with Pointers
```go
func TestPointerFunction(t *testing.T) {
    value := 42
    ptr := &value
    
    ModifyPointer(ptr)
    
    if value != 84 {
        t.Errorf("Expected value to be 84, got %d", value)
    }
}
```

### Testing Structs
```go
type Person struct {
    Name string
    Age  int
}

func TestStructFunction(t *testing.T) {
    input := Person{Name: "Alice", Age: 30}
    expected := Person{Name: "Alice", Age: 31}
    
    result := IncrementAge(input)
    
    if result != expected {
        t.Errorf("IncrementAge(%+v) = %+v, want %+v", 
            input, result, expected)
    }
}
```

## 🚀 Quick Reference

| Task | Command |
|------|---------|
| Run all tests | `go test ./piscine -v` |
| Run specific test | `go test ./piscine -v -run TestName` |
| Run with coverage | `go test ./piscine -v -cover` |
| Run benchmarks | `go test ./piscine -bench=.` |
| Generate coverage HTML | `go test ./piscine -coverprofile=coverage.out && go tool cover -html=coverage.out` |

## 💡 Tips

1. **Write tests BEFORE fixing bugs** - This helps you verify the fix works
2. **Keep tests simple** - Each test should be easy to understand
3. **Test edge cases** - That's where most bugs hide
4. **Use meaningful test data** - It makes debugging easier
5. **Run tests frequently** - Catch bugs early

## 🎓 Example Test Template

Copy this template for new tests:

```go
func TestYourFunction(t *testing.T) {
    tests := []struct {
        name     string
        input    interface{} // Change to your input type
        expected interface{} // Change to your output type
    }{
        {"describe test case", inputValue, expectedValue},
        // Add more test cases here
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := YourFunction(tt.input)
            if result != tt.expected {
                t.Errorf("YourFunction(%v) = %v, want %v", 
                    tt.input, result, tt.expected)
            }
        })
    }
}
```

Happy testing! 🎉