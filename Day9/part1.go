package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	risk := 0
	for y, row := range floorGrid {
		for x, val := range row {
			adjacencyCheck := [][]int{
				/*
					[]int{x + 1, y}, // right
					[]int{x - 1, y}, // left
					[]int{x, y + 1}, // down
					[]int{x, y - 1}, // up
				*/
			}
			if y > 0 {
				adjacencyCheck = append(adjacencyCheck, []int{x, y - 1})
			}
			if x > 0 {
				adjacencyCheck = append(adjacencyCheck, []int{x - 1, y})
			}
			if y < len(floorGrid)-1 {
				adjacencyCheck = append(adjacencyCheck, []int{x, y + 1})
			}
			if x < len(row)-1 {
				adjacencyCheck = append(adjacencyCheck, []int{x + 1, y})
			}
			adjacentValues := []int{}
			for _, direction := range adjacencyCheck {
				adjacentValues = append(adjacentValues, floorGrid[direction[1]][direction[0]])
			}
			valLowest := true
			for _, adjacentValue := range adjacentValues {
				valLowest = valLowest && adjacentValue > val
			}
			if valLowest {
				risk += val + 1
			}

		}
	}
	fmt.Println(risk)
}
