package piscine

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FindPairs() {
	if len(os.Args) != 3 {
		return
	}

	arrayStr := os.Args[1]
	targetStr := os.Args[2]

	// Parse and validate input
	numbers, err := parseArray(arrayStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	target, err := parseTarget(targetStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Find pairs
	pairs := findPairs(numbers, target)

	// Output result
	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
	} else {
		fmt.Printf("Pairs with sum %d: %v\n", target, pairs)
	}
}

func parseArray(s string) ([]int, error) {
	// Check format: must start with '[' and end with ']'
	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, "[") || !strings.HasSuffix(s, "]") {
		return nil, fmt.Errorf("Invalid input.")
	}

	// Remove brackets
	s = s[1 : len(s)-1]
	s = strings.TrimSpace(s)

	// Empty array
	if s == "" {
		return []int{}, nil
	}

	// Split by comma
	parts := strings.Split(s, ",")
	numbers := make([]int, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			return nil, fmt.Errorf("Invalid input.")
		}

		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("Invalid number: %s", part)
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}

func parseTarget(s string) (int, error) {
	s = strings.TrimSpace(s)

	// Check if contains space (multiple values)
	if strings.Contains(s, " ") {
		return 0, fmt.Errorf("Invalid target sum.")
	}

	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("Invalid target sum.")
	}

	return num, nil
}

func findPairs(numbers []int, target int) [][]int {
	var pairs [][]int
	used := make(map[int]bool)

	for i := 0; i < len(numbers); i++ {
		if used[i] {
			continue
		}

		for j := i + 1; j < len(numbers); j++ {
			if used[j] {
				continue
			}

			if numbers[i]+numbers[j] == target {
				pairs = append(pairs, []int{i, j})
				used[i] = true
				used[j] = true
				break
			}
		}
	}

	return pairs
}
