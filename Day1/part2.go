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
	amount := []int{}
	count := 0
	for scanner.Scan() {
		newHeight, _ := strconv.Atoi(scanner.Text())
		if len(amount) >= 3 {
			amount = amount[1:]
		}
		amount = append(amount, newHeight)
		sum := 0
		for _, value := range amount {
			sum += value
		}
		if sum > highestSoFar && len(amount) == 3 && highestSoFar != -1 {
			count += 1
		}
		if len(amount) == 3 {
			highestSoFar = sum
		}
	}
	fmt.Println(count)
}
