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
	for i := 0; i < 80; i++ {
		m = runIteration(m)
	}
	// Part1
	fmt.Println("Part 1: ", sumFish(m))
	for i := 0; i < 256-80; i++ {
		m = runIteration(m)
	}
	fmt.Println("Part 2: ", sumFish(m))
}

func runIteration(m map[int]int) map[int]int {
	newMap := make(map[int]int)
	for i := 0; i < 9; i++ {
		newMap[i] = 0
	}
	for key, element := range m {
		if key <= 0 {
			newMap[6] += element
			newMap[8] += element
		} else {
			newMap[key-1] += element
		}
	}
	return newMap
}

func sumFish(m map[int]int) int {
	sum := 0
	for _, element := range m {
		sum += element
	}
	return sum
}
