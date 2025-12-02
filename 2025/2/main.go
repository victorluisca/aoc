package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() string {
	filePath := "./2/input.txt"
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return ""
	}
	return string(content)
}

func parseLine(line string) (int, int) {
	parts := strings.Split(line, "-")

	min, _ := strconv.Atoi(parts[0])
	max, _ := strconv.Atoi(parts[1])

	return min, max
}

func hasPattern(id int) bool {
	strId := strconv.Itoa(id)
	middle := len(strId) / 2
	if strId[:middle] == strId[middle:] {
		return true
	}
	return false
}

func main() {
	lines := strings.Split(strings.TrimSpace(getInput()), ",")

	sum := 0
	for _, line := range lines {
		min, max := parseLine(line)
		for i := min; i < max; i++ {
			if hasPattern(i) {
				sum += i
			}
		}
	}

	fmt.Println(sum)
}
