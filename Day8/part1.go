package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inputs := [][]string{}
	for scanner.Scan() {
		newLine := scanner.Text()
		newLine = strings.Replace(newLine, " |", "", 1)
		inputs = append(inputs, strings.Split(newLine, " "))
	}
	count := 0
	uniqueDigits := []int{2, 4, 7, 3}
	for _, line := range inputs {
		for i := 10; i < 14; i++ {

			length := len(line[i])
			for _, uniqueLength := range uniqueDigits {
				if length == uniqueLength {
					count += 1
					break
				}
			}
		}
	}
	fmt.Println(count)
}
