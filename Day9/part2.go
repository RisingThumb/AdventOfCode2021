package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	floorGrid := [][]int{}
	for scanner.Scan() {
		floorRow := scanner.Text()
		floorRowSlice := []int{}
		for _, char := range floorRow {
			newHeight, _ := strconv.ParseInt(string(char), 10, 32)
			floorRowSlice = append(floorRowSlice, int(newHeight))
		}
		floorGrid = append(floorGrid, floorRowSlice)
	}
	sizes := []int{}
	for y, row := range floorGrid {
		for x, val := range row {
			if val < 9 && val != -1 {
				sizes = append(sizes, floodFill(x, y, floorGrid))
			}
		}
	}
	sort.Ints(sizes)
	// Too lazy to reverse
	fmt.Println(sizes[len(sizes)-3] * sizes[len(sizes)-2] * sizes[len(sizes)-1])
}

func floodFill(x, y int, floorGrid [][]int) int {
	count := 1
	floorGrid[y][x] = -1
	adjacencyCheck := [][]int{}
	if y > 0 {
		adjacencyCheck = append(adjacencyCheck, []int{x, y - 1})
	}
	if x > 0 {
		adjacencyCheck = append(adjacencyCheck, []int{x - 1, y})
	}
	if y < len(floorGrid)-1 {
		adjacencyCheck = append(adjacencyCheck, []int{x, y + 1})
	}
	if x < len(floorGrid[0])-1 {
		adjacencyCheck = append(adjacencyCheck, []int{x + 1, y})
	}
	for _, adjacency := range adjacencyCheck {
		if floorGrid[adjacency[1]][adjacency[0]] < 9 && floorGrid[adjacency[1]][adjacency[0]] != -1 {
			count += floodFill(adjacency[0], adjacency[1], floorGrid)
		}
	}
	return count
}
