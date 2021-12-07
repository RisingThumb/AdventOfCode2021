package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	linesSplit := strings.Split(line, ",")
	var linesInt []int
	// Read it as integers
	m := make(map[int]int)
	for i := 0; i < 9; i++ {
		m[i] = 0
	}
	for _, val := range linesSplit {
		newInt, _ := strconv.ParseInt(val, 10, 32)
		linesInt = append(linesInt, int(newInt))
		m[int(newInt)] += 1
	}

	minMaxFuel := minAndMaxInMapKeys(m)
	lowestFuel := -1
	for i := minMaxFuel[0]; i < minMaxFuel[1]; i++ {
		if lowestFuel == -1 {
			lowestFuel = fuelCostToMoveToPosition(m, i)
			continue
		}
		newFuelCost := fuelCostToMoveToPosition(m, i)
		if newFuelCost < lowestFuel {
			lowestFuel = newFuelCost
			continue
		}
	}
	fmt.Println(lowestFuel)
}

func minAndMaxInMapKeys(m map[int]int) [2]int {
	minMax := [2]int{}
	for key, _ := range m {
		if key < minMax[0] {
			minMax[0] = key
		}
		if key > minMax[1] {
			minMax[1] = key
		}
	}

	return minMax
}

func fuelCostToMoveToPosition(m map[int]int, pos int) int {
	cost := 0
	for key, value := range m {
		if key < pos {
			cost += (pos - key) * value
		}
		if key > pos {
			cost += (key - pos) * value
		}
	}
	return cost
}
