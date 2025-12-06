package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInput() string {
	filePath := "./5/input.txt"
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return ""
	}
	return string(content)
}

func splitRangesAndIds(lines []string) ([]string, []string) {
	rangeStrings := []string{}
	idStrings := []string{}

	for i, line := range lines {
		if line == "" {
			rangeStrings = lines[:i]
			idStrings = lines[i+1:]
		}
	}

	return rangeStrings, idStrings
}

func parseRange(rangeString string) (int, int) {
	parts := strings.Split(rangeString, "-")

	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])

	return start, end
}

func main() {
	lines := strings.Split(strings.TrimSpace(getInput()), "\n")

	rangeStrings, idStrings := splitRangesAndIds(lines)

	sum := 0

	for _, id := range idStrings {
		intId, _ := strconv.Atoi(id)
		for _, rangeLine := range rangeStrings {
			start, end := parseRange(rangeLine)
			if start <= intId && intId <= end {
				sum += 1
				break
			}
		}
	}

	fmt.Printf("Part 1: %v\n", sum)

	rangePairs := [][]int{}
	for _, rangeStr := range rangeStrings {
		start, end := parseRange(rangeStr)
		rangePairs = append(rangePairs, []int{start, end})
	}

	sort.Slice(rangePairs, func(i, j int) bool {
		return rangePairs[i][0] < rangePairs[j][0]
	})

	maxValueSeen := -1
	sum = 0
	for _, rangePair := range rangePairs {
		if rangePair[1] < maxValueSeen {
			continue
		}

		upper := rangePair[1]
		lower := int(math.Max(float64(maxValueSeen)+1, float64(rangePair[0])))
		maxValueSeen = rangePair[1]

		diff := upper - lower + 1
		sum += diff
	}
	fmt.Printf("Part 2: %v\n", sum)
}
