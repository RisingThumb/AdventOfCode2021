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
	highestSoFar := -1
	count := 0
	for scanner.Scan() {
		newHeight, _ := strconv.Atoi(scanner.Text())
		if newHeight > highestSoFar {
			if highestSoFar >= 0 {
				count += 1
			}

		}
		highestSoFar = newHeight
	}
	fmt.Println(count)
}
