package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() string {
	filePath := "./1/input.txt"
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return ""
	}
	return string(content)
}

func parseLine(line string) int {
	if line == "" {
		return 0
	}

	direction := string(line[0])
	distance, _ := strconv.Atoi(line[1:])

	if direction == "L" {
		return -distance
	}
	return distance
}

func main() {
	lines := strings.Split(strings.TrimSpace(getInput()), "\n")

	dialSize := 100

	state := 50
	count := 0

	for _, line := range lines {
		distance := parseLine(line)

		state += distance
		state = ((state % dialSize) + dialSize) % dialSize

		if state == 0 {
			count++
		}
	}

	fmt.Println(count)
}
