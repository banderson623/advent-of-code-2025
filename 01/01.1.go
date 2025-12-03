package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var dial int = 50
	var timesAtZero int = 0

	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	input := string(content)

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		var direction string
		var steps int

		_, err := fmt.Sscanf(line, "%1s%d", &direction, &steps)

		if err != nil {
			panic(err)
		}

		switch direction {
		case "L":
			dial -= steps
		case "R":
			dial += steps
		}

		dial = (100 + dial) % 100

		if dial == 0 {
			timesAtZero++
		}

		fmt.Printf("Dial is rotated %s to point at %d\n", line, dial)
	}

	fmt.Println("Password is", timesAtZero)
}
