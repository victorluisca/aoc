package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() string {
	filePath := "./3/input.txt"
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return ""
	}
	return string(content)
}

func findHighest(bank []int, start int, end int) (int, int) {
	high := bank[start]
	highIndex := start

	for i := start; i < end; i++ {
		if bank[i] > high {
			high = bank[i]
			highIndex = i
		}
	}

	return high, highIndex
}

func solve(bank []int, digits int) int {
	res := ""
	start := 0

	for i := digits - 1; i >= 0; i-- {
		end := len(bank) - i
		high, highIndex := findHighest(bank, start, end)
		res += strconv.Itoa(high)
		start = highIndex + 1
	}

	result, _ := strconv.Atoi(res)

	return result
}

func main() {
	lines := strings.Split(strings.TrimSpace(getInput()), "\n")

	sum := 0
	for _, line := range lines {
		bank := make([]int, 0, len(line))
		for _, battery := range line {
			num, _ := strconv.Atoi(string(battery))
			bank = append(bank, num)
		}

		joltage := solve(bank, 12)
		sum += joltage
	}

	fmt.Println(sum)
}
