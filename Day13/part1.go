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
    positionsForDots := [][]int{}
    maxY := 0
    maxX := 0
	for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            break
        }
        lineSplit := strings.Split(line, ",")
        newLines := []int{}
        x, _ := strconv.ParseInt(lineSplit[0], 10, 32)
        newLines = append(newLines, int(x))
        if int(x) > maxX {
            maxX = int(x)
        }
        y, _ := strconv.ParseInt(lineSplit[1], 10, 32)
        newLines = append(newLines, int(y))
        if int(y) > maxY {
            maxY = int(y)
        }
        positionsForDots = append(positionsForDots, newLines)
	}
    foldInstructions := []string{}
    for scanner.Scan() {
        foldLine := strings.Split(scanner.Text(), " ")[2]
        foldInstructions = append(foldInstructions, foldLine)
    }

	fmt.Println(positionsForDots)
	fmt.Println(maxX)
	fmt.Println(maxY)
    fmt.Println(foldInstructions)
    dotGrid := [][]string{}
    for _, instruction := range foldInstructions {
        instructionRelevant := strings.Split(instruction, "=")
        positionToSplitOn, _ := strconv.ParseInt(instructionRelevant[1], 10, 32)
        if instructionRelevant[0] == "x" {
            maxX = int(positionToSplitOn)
            positionsForDots = foldAlongXAxis(positionsForDots, maxX)
        }

        if instructionRelevant[0] == "y" {
            maxY = int(positionToSplitOn)
            positionsForDots = foldAlongYAxis(positionsForDots, maxY)
        }
        break
    }

    for i := 0; i <= maxY; i++ {
        newLine := []string{}
        for j :=0; j <= maxX; j++ {
            newLine = append(newLine, " ")
        }
        dotGrid = append(dotGrid, newLine)
    }
    for _, coordinate := range positionsForDots {
        dotGrid[coordinate[1]][coordinate[0]] = "#"
    }
    count := 0
    for _, row := range dotGrid {
        for _, val := range row {
            if val == "#" {
                count += 1
            }
        }
    }
    fmt.Println("Part 1 answer : ", count)

    for _, instruction := range foldInstructions[1:] {
        instructionRelevant := strings.Split(instruction, "=")
        positionToSplitOn, _ := strconv.ParseInt(instructionRelevant[1], 10, 32)
        if instructionRelevant[0] == "x" {
            maxX = int(positionToSplitOn)
            positionsForDots = foldAlongXAxis(positionsForDots, maxX)
        }

        if instructionRelevant[0] == "y" {
            maxY = int(positionToSplitOn)
            positionsForDots = foldAlongYAxis(positionsForDots, maxY)
        }
    }
    dotGrid = [][]string{}

    for i := 0; i <= maxY; i++ {
        newLine := []string{}
        for j :=0; j <= maxX; j++ {
            newLine = append(newLine, " ")
        }
        dotGrid = append(dotGrid, newLine)
    }
    for _, coordinate := range positionsForDots {
        dotGrid[coordinate[1]][coordinate[0]] = "#"
    }

    for _, line := range dotGrid {
        fmt.Println(line)
    }

}

func foldAlongXAxis(positionsForDots [][]int, xPosition int) [][]int {
    newCoordinates := [][]int{}
    for _, coordinate := range positionsForDots {
        if coordinate[0] < xPosition {
            newCoordinates = append(newCoordinates, coordinate)
        } else { // Needs moving in X Axis.
            diff := coordinate[0] - xPosition
            newCoordinates = append(newCoordinates, []int{xPosition-diff, coordinate[1]})
        }
    }
    return newCoordinates
}

func foldAlongYAxis(positionsForDots [][]int, yPosition int) [][]int {
    newCoordinates := [][]int{}
    for _, coordinate := range positionsForDots {
        if coordinate[1] < yPosition {
            newCoordinates = append(newCoordinates, coordinate)
        } else { // Needs moving in X Axis.
            diff := coordinate[1] - yPosition
            newCoordinates = append(newCoordinates, []int{coordinate[0], yPosition-diff})
        }
    }
    return newCoordinates
}
