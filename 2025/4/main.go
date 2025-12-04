package main

import (
	"fmt"
	"os"
	"strings"
)

func getInput() string {
	filePath := "./4/input.txt"
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
	return string(content)
}

func countAdjacent(grid [][]rune, r, c int) int {
	var directions = [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}

	rows := len(grid)
	cols := len(grid[0])

	count := 0
	for _, d := range directions {
		nr := r + d[0]
		nc := c + d[1]

		if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
			if grid[nr][nc] == '@' {
				count++
			}
		}
	}

	return count
}

func main() {
	lines := strings.Split(strings.TrimSpace(getInput()), "\n")
	grid := make([][]rune, len(lines))
	for i := range lines {
		grid[i] = []rune(lines[i])
	}

	count := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == '@' {
				if countAdjacent(grid, r, c) < 4 {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
