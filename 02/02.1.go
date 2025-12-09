package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	input := string(content)

	var id_ranges []string

	id_ranges = strings.Split(strings.TrimSpace(input), ",")

	var sum_of_palindromes int = 0

	for _, inventory := range id_ranges {
		range_parts := strings.Split(strings.TrimSpace(inventory), "-")

		start, err := strconv.Atoi(range_parts[0])
		if err != nil {
			fmt.Printf("Error converting start: %v\n", err)
			continue
		}
		end, err := strconv.Atoi(range_parts[1])
		if err != nil {
			fmt.Printf("Error converting end: %v\n", err)
			continue
		}

		for i := start; i <= end; i++ {

			digits := strconv.Itoa(i)
			halfway := len(digits) / 2

			if digits[0:halfway] == digits[halfway:] {
				fmt.Println("This is a palindrome:", i)
				sum_of_palindromes += i
			}

		}
	}
	fmt.Println("sum_of_palindromes", sum_of_palindromes)
}
