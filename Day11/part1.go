package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Octopus struct {
    flashed bool
    energy int
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
    grid := [][]Octopus{}
	for scanner.Scan() {
        newLine := scanner.Text()
        row := []Octopus{}
        for _, char := range newLine {
            newValue, _ := strconv.ParseInt(string(char), 10, 32)
            newOctopus := Octopus{false, int(newValue)}
            row = append(row, newOctopus)
        }
        grid = append(grid, row)
	}
    count := 0
    allFlashed := false
    for i:=0; !allFlashed; i++ {
        increaseEnergyLevels(grid)
        handleFlashing(grid)
        newFlashes, didAllFlash := resetEnergy(grid)
        allFlashed = didAllFlash
        count += newFlashes
        if i == 99 {
            fmt.Println(count)
        }
        if (allFlashed) {
            fmt.Println(allFlashed, i+1)
        }
    }
}

func printGrid(grid [][]Octopus) {
    for y, row := range grid {
        for x := range row {
            fmt.Printf("%d", grid[y][x].energy)
        }
        fmt.Printf("\n")
    }
    fmt.Printf("\n")
}

func increaseEnergyLevels(grid [][]Octopus) {
    for y, row := range grid {
        for x := range row {
            grid[y][x].energy += 1
        }
    }
}

func handleFlashing(grid [][]Octopus) {
    flashHappened := true
    for flashHappened {
        flashHappened = false
        for y, row := range grid {
            for x := range row {
                if !grid[y][x].flashed && grid[y][x].energy > 9 {
                    grid[y][x].flashed = true
                    flashHappened = true
                    for _, i := range []int{1, 0, -1} {
                        for _, j := range []int{1, 0, -1} {
                            if i == 0 && i == j {
                                continue
                            }
                            if !((x + i < 0) || (x + i > len(row) - 1)) {
                                if !((y + j < 0) || (y + j > len(grid) - 1)) {
                                    grid[y+j][x+i].energy += 1
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}


func resetEnergy(grid[][]Octopus) (int, bool){
    flashedThisStep := 0
    allFlashedThisStep := true
    for y, row := range grid {
        for x, octopus := range row {
            allFlashedThisStep = allFlashedThisStep && grid[y][x].flashed
            if octopus.energy > 9 {
                grid[y][x].energy = 0
                grid[y][x].flashed = false
                flashedThisStep += 1
            }
        }
    }
    return flashedThisStep, allFlashedThisStep
}
