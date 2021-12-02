package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "strings"
    "strconv"
)

type position struct {
    X int
    Y int
    Aim int
}

type instruction struct {
    Direction string
    Amount int
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
    currentPosition := position{0, 0, 0}
	for scanner.Scan() {
		newInstruction := scanner.Text()
        words := strings.Fields(newInstruction)
        amount, _ := strconv.Atoi(words[1])
        //newDirection := instruction{words[0], amount,}
        switch(words[0]) {
            case "forward":
                currentPosition.X += amount
                currentPosition.Y += amount * currentPosition.Aim
            case "down":
                currentPosition.Aim += amount
            case "up":
                currentPosition.Aim -= amount
        }
	}
    fmt.Println(currentPosition.X*currentPosition.Y);
}
