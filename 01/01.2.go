package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var dial int = 50
	var password int = 0

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
		startingDial := dial

		_, err := fmt.Sscanf(line, "%1s%d", &direction, &steps)

		if err != nil {
			panic(err)
		}

		// Check for complete rotations (100 steps or more)
		// each full rotation adds 1 to the password
		// (oh and only keep the remainder, like the non full rotations as the new step)
		if steps >= 100 {
			wholeNumberRotations := steps / 100
			password += wholeNumberRotations
			steps = steps % 100
			fmt.Printf("...%s instructions caused %d whole rotations detected, password increased to %d\n", line, wholeNumberRotations, password)
		}

		// move the dial
		switch direction {
		case "L":
			dial -= steps

		case "R":
			dial += steps
		}

		dialBeforeMod := dial
		dial = (dial + 10000) % 100

		fmt.Printf("Dial is rotated %s to point at %d\n", line, dial)

		zeroWasAlreadyCounted := startingDial == 0
		if zeroWasAlreadyCounted {
			continue
		}

		// If we are at zero, or passed it
		passedZero := dialBeforeMod != dial
		if dial == 0 || passedZero {
			password++
			fmt.Println("...dial landed on zero, or passed zero", dial, "(before mod", dialBeforeMod, ")", "password is now", password)
		}

	}

	fmt.Println("Password is", password)
}
