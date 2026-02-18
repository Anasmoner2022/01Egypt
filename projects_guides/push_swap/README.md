# Push-Swap Project Guide

## 📋 Project Overview
Build a sorting program using only stack operations. You have two stacks (a and b) and a limited set of operations to sort integers in ascending order with the minimum number of moves. This project teaches you about sorting algorithms, optimization, and working with constraints. You'll build two programs: **push-swap** (the solver) and **checker** (the verifier).

**The Challenge**: Sort numbers using ONLY the allowed operations, and do it efficiently!

---

## 🎯 Learning Objectives

By completing this project, you will learn:
1. **Stack Data Structures**: Operations, memory management
2. **Sorting Algorithms**: Especially constrained sorting
3. **Algorithm Optimization**: Minimizing operation count
4. **Problem Decomposition**: Breaking complex problems into steps
5. **Algorithm Design**: Creating custom sorting strategies
6. **Complexity Analysis**: Understanding time and space complexity
7. **Input Validation**: Robust error handling
8. **Testing**: Comprehensive test coverage
9. **Performance Metrics**: Measuring algorithm efficiency

---

## 📚 Prerequisites - Topics You Must Know

### 1. **Stack Data Structure**
**What is a Stack?**
- LIFO (Last In, First Out)
- Like a stack of plates: add/remove from top only
- Two main operations: push (add), pop (remove)

```
Stack Example:
         [5]  ← top (can push/pop here)
         [3]
         [7]
         [1]  ← bottom (cannot access directly)
```

**Stack in Go**:
```go
// Using slice as stack
stack := []int{}

// Push
stack = append(stack, 5)

// Pop
top := stack[len(stack)-1]
stack = stack[:len(stack)-1]

// Peek (look at top without removing)
top := stack[len(stack)-1]
```

### 2. **Sorting Algorithms Basics**
You should understand:
- **Bubble Sort**: Compare and swap neighbors
- **Selection Sort**: Find minimum, place it
- **Insertion Sort**: Insert element in sorted position
- **Quick Sort**: Partition and recursion
- **Merge Sort**: Divide and conquer

**But here's the catch**: These won't work directly! You can only use the allowed operations.

### 3. **Algorithm Complexity**
- **Time Complexity**: How operations grow with input size
- **O(n)**: Linear - operations proportional to n
- **O(n²)**: Quadratic - operations grow with n²
- **O(n log n)**: Efficient sorting algorithms

### 4. **Command-Line Arguments in Go**
```go
import "os"

args := os.Args[1:]  // Get arguments (skip program name)
```

### 5. **String to Int Conversion**
```go
import "strconv"

num, err := strconv.Atoi("123")
if err != nil {
    // Handle error
}
```

---

## 🎮 Understanding the Operations

### **Available Operations**

**Push Operations**:
```
pa (push a): Take top of b, put on top of a
pb (push b): Take top of a, put on top of b

Before:     After pa:    After pb:
 2  3        3            5  2
 5  7        2  7         9  3
 9           5  9            7
 = =         = =          = =
 a  b        a  b         a  b
```

**Swap Operations**:
```
sa (swap a): Swap first two elements of a
sb (swap b): Swap first two elements of b
ss: Do both sa and sb

Before:     After sa:
 2           5
 5           2
 9           9
 =           =
 a           a
```

**Rotate Operations** (shift up):
```
ra (rotate a): First element becomes last
rb (rotate b): First element becomes last
rr: Do both ra and rb

Before:     After ra:
 2           5
 5           9
 9           2
 =           =
 a           a
```

**Reverse Rotate Operations** (shift down):
```
rra (reverse rotate a): Last element becomes first
rrb (reverse rotate b): Last element becomes first  
rrr: Do both rra and rrb

Before:     After rra:
 2           9
 5           2
 9           5
 =           =
 a           a
```

---

## 🧠 Understanding the Problem

### **Goal**
Sort stack `a` in ascending order (smallest on top) with minimum operations.

### **Constraints**
- Only use the 11 allowed operations
- Can use stack `b` as helper
- Must minimize operation count

### **Example Walkthrough**

**Input**: `"2 1 3 6 5 8"`

```
Initial:
 2
 1
 3
 6
 5
 8
 = =
 a b

Goal:
 1
 2
 3
 5
 6
 8
 = =
 a b
```

**One Solution**:
```
1. pb    (push 2 to b)
2. pb    (push 1 to b)
3. ra    (rotate a: 3→6→5→8→3)
4. sa    (swap 6 and 3)
5. rrr   (reverse rotate both)
6. pa    (push 1 back)
7. pa    (push 2 back)
```

---

## 🛠️ Step-by-Step Implementation Guide

### **Phase 1: Data Structures** 📦

#### Step 1: Define Stack Structure
```go
type Stack struct {
    data []int
}

func NewStack() *Stack {
    return &Stack{
        data: make([]int, 0),
    }
}

func (s *Stack) Push(value int) {
    s.data = append(s.data, value)
}

func (s *Stack) Pop() (int, error) {
    if len(s.data) == 0 {
        return 0, errors.New("stack is empty")
    }
    value := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]
    return value, nil
}

func (s *Stack) Peek() (int, error) {
    if len(s.data) == 0 {
        return 0, errors.New("stack is empty")
    }
    return s.data[len(s.data)-1], nil
}

func (s *Stack) Size() int {
    return len(s.data)
}

func (s *Stack) IsEmpty() bool {
    return len(s.data) == 0
}
```

**Test Your Stack**:
```go
func TestStack() {
    s := NewStack()
    s.Push(1)
    s.Push(2)
    s.Push(3)
    
    top, _ := s.Pop()  // Should be 3
    fmt.Println(top)
}
```

---

#### Step 2: Create Stack Operations
```go
type StackPair struct {
    A *Stack
    B *Stack
}

func NewStackPair() *StackPair {
    return &StackPair{
        A: NewStack(),
        B: NewStack(),
    }
}
```

**Implement Each Operation**:

**Push Operations**:
```go
func (sp *StackPair) Pa() {
    // Push from B to A
    if !sp.B.IsEmpty() {
        value, _ := sp.B.Pop()
        sp.A.Push(value)
    }
}

func (sp *StackPair) Pb() {
    // Push from A to B
    if !sp.A.IsEmpty() {
        value, _ := sp.A.Pop()
        sp.B.Push(value)
    }
}
```

**Swap Operations**:
```go
func (sp *StackPair) Sa() {
    // Swap first two elements of A
    if sp.A.Size() >= 2 {
        first, _ := sp.A.Pop()
        second, _ := sp.A.Pop()
        sp.A.Push(first)
        sp.A.Push(second)
    }
}

func (sp *StackPair) Sb() {
    // Swap first two elements of B
    if sp.B.Size() >= 2 {
        first, _ := sp.B.Pop()
        second, _ := sp.B.Pop()
        sp.B.Push(first)
        sp.B.Push(second)
    }
}

func (sp *StackPair) Ss() {
    sp.Sa()
    sp.Sb()
}
```

**Rotate Operations**:
```go
func (sp *StackPair) Ra() {
    // Rotate A (first becomes last)
    if sp.A.Size() >= 2 {
        // Remove from top
        value, _ := sp.A.Pop()
        
        // Shift all elements up
        temp := make([]int, sp.A.Size())
        copy(temp, sp.A.data)
        
        // Clear stack
        sp.A.data = make([]int, 0)
        
        // Put value at bottom
        sp.A.Push(value)
        
        // Add rest back
        for _, v := range temp {
            sp.A.Push(v)
        }
    }
}
```

**Better Rotate Implementation** (more efficient):
```go
func (sp *StackPair) Ra() {
    if sp.A.Size() >= 2 {
        // Take top element
        top := sp.A.data[len(sp.A.data)-1]
        
        // Shift everything
        sp.A.data = sp.A.data[:len(sp.A.data)-1]
        
        // Insert at bottom
        sp.A.data = append([]int{top}, sp.A.data...)
    }
}

func (sp *StackPair) Rb() {
    if sp.B.Size() >= 2 {
        top := sp.B.data[len(sp.B.data)-1]
        sp.B.data = sp.B.data[:len(sp.B.data)-1]
        sp.B.data = append([]int{top}, sp.B.data...)
    }
}

func (sp *StackPair) Rr() {
    sp.Ra()
    sp.Rb()
}
```

**Reverse Rotate Operations**:
```go
func (sp *StackPair) Rra() {
    if sp.A.Size() >= 2 {
        // Take bottom element
        bottom := sp.A.data[0]
        
        // Remove it
        sp.A.data = sp.A.data[1:]
        
        // Add to top
        sp.A.data = append(sp.A.data, bottom)
    }
}

func (sp *StackPair) Rrb() {
    if sp.B.Size() >= 2 {
        bottom := sp.B.data[0]
        sp.B.data = sp.B.data[1:]
        sp.B.data = append(sp.B.data, bottom)
    }
}

func (sp *StackPair) Rrr() {
    sp.Rra()
    sp.Rrb()
}
```

---

### **Phase 2: Input Parsing** 📥

#### Step 3: Parse Command-Line Arguments
```go
func ParseArguments(args []string) ([]int, error) {
    if len(args) == 0 {
        return []int{}, nil
    }
    
    // Join all arguments
    input := strings.Join(args, " ")
    
    // Split by spaces
    parts := strings.Fields(input)
    
    numbers := make([]int, 0)
    seen := make(map[int]bool)
    
    for _, part := range parts {
        // Convert to int
        num, err := strconv.Atoi(part)
        if err != nil {
            return nil, fmt.Errorf("invalid input")
        }
        
        // Check for duplicates
        if seen[num] {
            return nil, fmt.Errorf("duplicate value")
        }
        seen[num] = true
        
        numbers = append(numbers, num)
    }
    
    return numbers, nil
}
```

**Test Cases**:
```go
// Valid
nums, _ := ParseArguments([]string{"2", "1", "3"})
// nums = [2, 1, 3]

// Valid (single string)
nums, _ := ParseArguments([]string{"2 1 3 6 5 8"})
// nums = [2, 1, 3, 6, 5, 8]

// Invalid (not a number)
_, err := ParseArguments([]string{"0", "one", "2"})
// err != nil

// Invalid (duplicate)
_, err := ParseArguments([]string{"1", "2", "2", "3"})
// err != nil
```

---

#### Step 4: Initialize Stacks
```go
func InitializeStacks(numbers []int) *StackPair {
    sp := NewStackPair()
    
    // Push numbers to stack A in reverse order
    // (so first number is on top)
    for i := len(numbers) - 1; i >= 0; i-- {
        sp.A.Push(numbers[i])
    }
    
    return sp
}
```

**Example**:
```go
numbers := []int{2, 1, 3, 6, 5, 8}
sp := InitializeStacks(numbers)

// Stack A will be (top to bottom):
// 2, 1, 3, 6, 5, 8
```

---

### **Phase 3: Checker Program** ✅

#### Step 5: Build the Checker
The checker is simpler - it just executes instructions and checks if sorted.

```go
// checker/main.go
func main() {
    args := os.Args[1:]
    
    // Parse arguments
    numbers, err := ParseArguments(args)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error")
        return
    }
    
    // If no arguments, exit
    if len(numbers) == 0 {
        return
    }
    
    // Initialize stacks
    sp := InitializeStacks(numbers)
    
    // Read and execute instructions
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        instruction := scanner.Text()
        
        err := ExecuteInstruction(sp, instruction)
        if err != nil {
            fmt.Fprintln(os.Stderr, "Error")
            return
        }
    }
    
    // Check if sorted
    if IsSorted(sp) {
        fmt.Println("OK")
    } else {
        fmt.Println("KO")
    }
}
```

---

#### Step 6: Execute Instructions
```go
func ExecuteInstruction(sp *StackPair, instruction string) error {
    switch instruction {
    case "pa":
        sp.Pa()
    case "pb":
        sp.Pb()
    case "sa":
        sp.Sa()
    case "sb":
        sp.Sb()
    case "ss":
        sp.Ss()
    case "ra":
        sp.Ra()
    case "rb":
        sp.Rb()
    case "rr":
        sp.Rr()
    case "rra":
        sp.Rra()
    case "rrb":
        sp.Rrb()
    case "rrr":
        sp.Rrr()
    default:
        return fmt.Errorf("invalid instruction: %s", instruction)
    }
    return nil
}
```

---

#### Step 7: Check if Sorted
```go
func IsSorted(sp *StackPair) bool {
    // Stack B must be empty
    if !sp.B.IsEmpty() {
        return false
    }
    
    // Stack A must be sorted in ascending order
    // (Remember: top is smallest)
    data := sp.A.data
    
    for i := len(data) - 1; i > 0; i-- {
        if data[i] > data[i-1] {
            return false
        }
    }
    
    return true
}
```

**Test Checker**:
```bash
echo -e "sa\nrra\npb" | ./checker "3 2 1 0"
# Should output: KO

echo -e "rra\npb\nsa\nrra\npa" | ./checker "3 2 1 0"
# Should output: OK
```

---

### **Phase 4: Sorting Strategies** 🧩

Now the challenging part - creating an efficient sorting algorithm!

#### Step 8: Strategy for Small Numbers (2-3 elements)

**Sort 2 Elements**:
```go
func Sort2(sp *StackPair, ops *[]string) {
    // If first > second, swap
    if sp.A.data[len(sp.A.data)-1] > sp.A.data[len(sp.A.data)-2] {
        sp.Sa()
        *ops = append(*ops, "sa")
    }
}
```

**Sort 3 Elements**:
```go
func Sort3(sp *StackPair, ops *[]string) {
    // Get the three values (top to bottom)
    size := sp.A.Size()
    a := sp.A.data[size-1]  // top
    b := sp.A.data[size-2]  // middle
    c := sp.A.data[size-3]  // bottom
    
    // Find which is smallest/largest
    smallest := min(a, min(b, c))
    largest := max(a, max(b, c))
    
    // Different cases
    if a == smallest {
        // Smallest is on top
        if b > c {
            sp.Sa()
            sp.Ra()
            sp.Sa()
            sp.Rra()
            *ops = append(*ops, "sa", "ra", "sa", "rra")
        }
        // Already sorted
    } else if a == largest {
        // Largest is on top
        sp.Ra()
        *ops = append(*ops, "ra")
        if b > c {
            sp.Sa()
            *ops = append(*ops, "sa")
        }
    } else {
        // Middle value on top
        if b == smallest {
            sp.Sa()
            *ops = append(*ops, "sa")
        } else {
            sp.Rra()
            *ops = append(*ops, "rra")
        }
    }
}
```

---

#### Step 9: Strategy for Medium Numbers (4-5 elements)

**Sort 4-5 Elements**:
```go
func SortSmall(sp *StackPair, ops *[]string) {
    size := sp.A.Size()
    
    // Push smallest to B
    for size > 3 {
        // Find position of smallest
        minPos := FindMinPosition(sp.A)
        
        // Rotate to bring smallest to top
        if minPos <= size/2 {
            // Closer to top, rotate forward
            for i := 0; i < minPos; i++ {
                sp.Ra()
                *ops = append(*ops, "ra")
            }
        } else {
            // Closer to bottom, reverse rotate
            for i := minPos; i < size; i++ {
                sp.Rra()
                *ops = append(*ops, "rra")
            }
        }
        
        // Push to B
        sp.Pb()
        *ops = append(*ops, "pb")
        size--
    }
    
    // Sort remaining 3 in A
    Sort3(sp, ops)
    
    // Push back from B
    for !sp.B.IsEmpty() {
        sp.Pa()
        *ops = append(*ops, "pa")
    }
}

func FindMinPosition(s *Stack) int {
    minVal := s.data[0]
    minPos := 0
    
    for i, val := range s.data {
        if val < minVal {
            minVal = val
            minPos = i
        }
    }
    
    // Return position from top
    return len(s.data) - 1 - minPos
}
```

---

#### Step 10: Strategy for Large Numbers (100+ elements)

For large numbers, we use a **Chunk/Radix-based approach**:

**Concept**:
1. Divide numbers into chunks
2. Push chunks to B in order
3. Push back to A in sorted order

**Implementation**:
```go
func SortLarge(sp *StackPair, ops *[]string) {
    size := sp.A.Size()
    
    // Calculate chunk size
    chunkSize := calculateChunkSize(size)
    
    // Normalize values (give them ranks)
    normalized := normalizeValues(sp.A.data)
    
    // Push to B in chunks
    pushed := 0
    chunk := 0
    
    for !sp.A.IsEmpty() {
        topValue := sp.A.data[len(sp.A.data)-1]
        normalizedValue := normalized[topValue]
        
        // Check if in current chunk range
        if normalizedValue <= (chunk+1)*chunkSize {
            // Push to B
            sp.Pb()
            *ops = append(*ops, "pb")
            pushed++
            
            // Rotate B to keep larger values near top
            if sp.B.Size() > 1 && 
               normalized[sp.B.data[len(sp.B.data)-1]] < size/2 {
                sp.Rb()
                *ops = append(*ops, "rb")
            }
            
            // Move to next chunk when current is full
            if pushed >= (chunk+1)*chunkSize {
                chunk++
            }
        } else {
            // Not in range, rotate A
            sp.Ra()
            *ops = append(*ops, "ra")
        }
    }
    
    // Push back to A in sorted order
    pushBackSorted(sp, ops, normalized)
}

func calculateChunkSize(size int) int {
    if size <= 100 {
        return size / 5
    } else {
        return size / 11
    }
}
```

---

#### Step 11: Normalize Values (Ranking)
```go
func normalizeValues(data []int) map[int]int {
    // Create sorted copy
    sorted := make([]int, len(data))
    copy(sorted, data)
    sort.Ints(sorted)
    
    // Map each value to its rank
    normalized := make(map[int]int)
    for i, val := range sorted {
        normalized[val] = i
    }
    
    return normalized
}
```

**Example**:
```
Input:  [42, 13, 7, 99, 5]
Sorted: [5, 7, 13, 42, 99]
Ranks:  5→0, 7→1, 13→2, 42→3, 99→4
```

---

#### Step 12: Push Back Sorted
```go
func pushBackSorted(sp *StackPair, ops *[]string, normalized map[int]int) {
    size := sp.B.Size()
    
    for i := size - 1; i >= 0; i-- {
        // Find largest value in B
        maxPos := findMaxPosition(sp.B)
        
        // Rotate to bring it to top
        if maxPos <= sp.B.Size()/2 {
            for j := 0; j < maxPos; j++ {
                sp.Rb()
                *ops = append(*ops, "rb")
            }
        } else {
            for j := maxPos; j < sp.B.Size(); j++ {
                sp.Rrb()
                *ops = append(*ops, "rrb")
            }
        }
        
        // Push to A
        sp.Pa()
        *ops = append(*ops, "pa")
    }
}

func findMaxPosition(s *Stack) int {
    maxVal := s.data[0]
    maxPos := 0
    
    for i, val := range s.data {
        if val > maxVal {
            maxVal = val
            maxPos = i
        }
    }
    
    return len(s.data) - 1 - maxPos
}
```

---

### **Phase 5: Push-Swap Program** 🎯

#### Step 13: Main Push-Swap Logic
```go
// push-swap/main.go
func main() {
    args := os.Args[1:]
    
    // Parse arguments
    numbers, err := ParseArguments(args)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error")
        return
    }
    
    // If no arguments or already sorted, exit
    if len(numbers) == 0 {
        return
    }
    
    if isAlreadySorted(numbers) {
        return
    }
    
    // Initialize stacks
    sp := InitializeStacks(numbers)
    
    // Solve
    operations := SolveOptimized(sp)
    
    // Print operations
    for _, op := range operations {
        fmt.Println(op)
    }
}
```

---

#### Step 14: Optimized Solver
```go
func SolveOptimized(sp *StackPair) []string {
    ops := make([]string, 0)
    size := sp.A.Size()
    
    switch {
    case size == 2:
        Sort2(sp, &ops)
    case size == 3:
        Sort3(sp, &ops)
    case size <= 5:
        SortSmall(sp, &ops)
    default:
        SortLarge(sp, &ops)
    }
    
    return ops
}

func isAlreadySorted(numbers []int) bool {
    for i := 0; i < len(numbers)-1; i++ {
        if numbers[i] > numbers[i+1] {
            return false
        }
    }
    return true
}
```

---

### **Phase 6: Optimization** ⚡

#### Step 15: Operation Compression
```go
func CompressOperations(ops []string) []string {
    compressed := make([]string, 0)
    i := 0
    
    for i < len(ops) {
        // Look for consecutive operations that can be combined
        if i < len(ops)-1 {
            curr := ops[i]
            next := ops[i+1]
            
            // Combine sa + sb → ss
            if curr == "sa" && next == "sb" {
                compressed = append(compressed, "ss")
                i += 2
                continue
            }
            
            // Combine ra + rb → rr
            if curr == "ra" && next == "rb" {
                compressed = append(compressed, "rr")
                i += 2
                continue
            }
            
            // Combine rra + rrb → rrr
            if curr == "rra" && next == "rrb" {
                compressed = append(compressed, "rrr")
                i += 2
                continue
            }
            
            // Check for cancellations
            // ra followed by rra cancels out
            if (curr == "ra" && next == "rra") ||
               (curr == "rra" && next == "ra") {
                i += 2
                continue
            }
        }
        
        compressed = append(compressed, ops[i])
        i++
    }
    
    return compressed
}
```

---

### **Phase 7: Testing** 🧪

#### Step 16: Create Test Suite
```go
func TestPushSwap() {
    tests := []struct {
        input    []int
        maxOps   int
        name     string
    }{
        {[]int{2, 1}, 1, "2 elements"},
        {[]int{2, 1, 3}, 3, "3 elements"},
        {[]int{5, 4, 3, 2, 1}, 12, "5 elements reverse"},
        {generateRandom(100), 700, "100 random"},
    }
    
    for _, tt := range tests {
        sp := InitializeStacks(tt.input)
        ops := SolveOptimized(sp)
        
        // Check operation count
        if len(ops) > tt.maxOps {
            fmt.Printf("FAIL %s: %d operations (max %d)\n", 
                tt.name, len(ops), tt.maxOps)
        }
        
        // Verify solution
        sp2 := InitializeStacks(tt.input)
        for _, op := range ops {
            ExecuteInstruction(sp2, op)
        }
        
        if !IsSorted(sp2) {
            fmt.Printf("FAIL %s: not sorted!\n", tt.name)
        }
    }
}
```

---

## 📊 Performance Targets

### **Operation Count Limits**

| Size | Complexity | Max Operations | Strategy |
|------|------------|----------------|----------|
| 2 | O(1) | 1 | Direct swap |
| 3 | O(1) | 3 | Case analysis |
| 5 | O(n) | 12 | Push min, sort 3 |
| 100 | O(n log n) | 700 | Chunk-based |
| 500 | O(n log n) | 5500 | Optimized chunks |

**Grade Targets**:
- 100 numbers in < 700 operations: Good
- 100 numbers in < 550 operations: Excellent
- 500 numbers in < 5500 operations: Bonus

---

## 🐛 Common Issues and Solutions

### Issue 1: Too Many Operations
**Problem**: Solution uses 1000+ operations for 100 numbers
**Solution**: Implement chunk-based algorithm, not bubble sort

### Issue 2: Stack Underflow
**Problem**: Trying to pop from empty stack
**Solution**: Always check `!stack.IsEmpty()` before operations

### Issue 3: Wrong Sort Direction
**Problem**: Stack sorted in descending order
**Solution**: Remember smallest should be on TOP

### Issue 4: Duplicate Detection
**Problem**: Not catching duplicates
**Solution**: Use map to track seen values during parsing

### Issue 5: Memory Issues
**Problem**: Operations list grows too large
**Solution**: Use efficient data structures, consider operation compression

---

## 📋 Testing Checklist

**Basic Functionality**:
- [ ] No arguments → no output
- [ ] Already sorted → no output
- [ ] Invalid input → "Error"
- [ ] Duplicates → "Error"
- [ ] 2 elements sorted correctly
- [ ] 3 elements sorted correctly
- [ ] 5 elements < 12 operations

**Checker**:
- [ ] Accepts valid instructions
- [ ] Rejects invalid instructions
- [ ] Outputs "OK" for correct sort
- [ ] Outputs "KO" for wrong sort
- [ ] Handles errors properly

**Performance**:
- [ ] 100 random numbers < 700 operations
- [ ] Solution verified by checker
- [ ] Different random sets work
- [ ] Large numbers handle correctly

---

## ✅ Submission Checklist

**Code Quality**:
- [ ] Two executables: push-swap, checker
- [ ] Clean, readable code
- [ ] Proper error handling
- [ ] No memory leaks
- [ ] Only standard packages

**Functionality**:
- [ ] All operations implemented correctly
- [ ] Parsing handles all cases
- [ ] Efficient sorting algorithm
- [ ] Operations minimized
- [ ] Checker validates correctly

**Testing**:
- [ ] Unit tests for operations
- [ ] Test suite for sorting
- [ ] Edge cases covered
- [ ] Performance verified

---

## 📖 Algorithm Strategies Explained

### **Why Chunks?**
Instead of sorting all at once:
1. Divide into manageable groups
2. Move groups to B strategically
3. Rebuild in sorted order

**Analogy**: Sorting a deck of cards
- Don't try to sort perfectly in one pass
- Group cards into ranges
- Merge groups in order

### **Turk Algorithm** (Advanced)
Alternative efficient approach:
1. Calculate "cost" to move each element
2. Choose cheapest move
3. Execute and repeat

**Cost Calculation**:
- Rotations needed in A
- Rotations needed in B
- Combined operations (rr, rrr)

---

## 🚀 Pro Tips

1. **Start Simple**: Get 3-element sort working perfectly
2. **Test Incrementally**: Test each size separately
3. **Visualize**: Draw stacks on paper to understand
4. **Use Checker**: Always verify with checker
5. **Benchmark**: Track operation counts
6. **Optimize Later**: Get working solution first
7. **Random Testing**: Test with many random inputs
8. **Edge Cases**: Test empty, single element, sorted
9. **Debug Output**: Add verbose mode during development
10. **Study Algorithms**: Research sorting under constraints

---

## 💡 Extension Ideas

After completing requirements:

1. **Visualizer**: Create visual display of sorting
2. **Statistics**: Track operation types used
3. **Comparison**: Compare different algorithms
4. **Optimizer**: Post-process to reduce operations
5. **Parallel**: Try different strategies, pick best
6. **Benchmarking**: Automated performance testing
7. **Alternative Algos**: Implement Turk algorithm
8. **Analysis Tool**: Analyze why certain inputs are hard

---

## 📚 Additional Resources

**Sorting Algorithms**:
- [Sorting Algorithm Visualizations](https://visualgo.net/en/sorting)
- [Big-O Cheat Sheet](https://www.bigocheatsheet.com/)
- [Push-Swap Visualization](https://github.com/o-reo/push_swap_visualizer)

**Stack Operations**:
- [Stack Data Structure](https://www.geeksforgeeks.org/stack-data-structure/)
- [Go Slices as Stacks](https://gobyexample.com/slices)

**Algorithm Design**:
- [Introduction to Algorithms](https://mitpress.mit.edu/books/introduction-algorithms)
- [Algorithms Course](https://www.coursera.org/learn/algorithms-part1)

---

## 🎓 Learning Path

**Week 1**: Stack operations + Checker + Small sorts (2-5)
**Week 2**: Algorithm design + Medium implementation
**Week 3**: Large number optimization + Testing
**Week 4**: Performance tuning + Edge cases

---

## 🔍 Debugging Strategies

**Visualize Operations**:
```go
func printStacks(sp *StackPair, op string) {
    fmt.Println("Operation:", op)
    fmt.Println("Stack A:", sp.A.data)
    fmt.Println("Stack B:", sp.B.data)
    fmt.Println("---")
}
```

**Track Operation Count**:
```go
func analyzeOperations(ops []string) {
    counts := make(map[string]int)
    for _, op := range ops {
        counts[op]++
    }
    fmt.Println("Operation counts:", counts)
}
```

**Test Specific Cases**:
```bash
# Test worst case for size 5
./push-swap "5 4 3 2 1"

# Test already sorted
./push-swap "1 2 3 4 5"

# Test random
./push-swap "3 1 4 2 5"
```

---

**Remember**: Push-swap teaches you to think creatively within constraints. Real-world programming often involves working with limitations - this project builds that skill! The goal isn't just to sort, but to sort *efficiently* with limited tools. 🎯📚