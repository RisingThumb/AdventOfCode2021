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
    grid := [1000][1000]int{}

	for scanner.Scan() {
        currentLine := scanner.Text()
        currentLineFormatted := strings.Replace(currentLine, " -> ", ",", 1)
        currentLineSlice := strings.Split(currentLineFormatted, ",")
        currentLineCasted := [4]int{}
        for i, coord := range currentLineSlice {
            parsedCoord, _ := strconv.ParseInt(coord, 10, 32)
            currentLineCasted[i] = int(parsedCoord)
        }
        grid = draw_line(currentLineCasted[0], currentLineCasted[1], currentLineCasted[2], currentLineCasted[3], grid)
	}
    count := 0
    for _, row := range grid {
        for _, val := range row {
            if val >= 2 {
                count+=1
            }
        }
    }
    fmt.Println(count)
}

func draw_line(x1, y1, x2, y2 int, grid [1000][1000]int) [1000][1000]int {
    // Vertical case
    if x1 == x2 {
        // Swap around params
        if y2 < y1 {
            temp := y2
            y2 = y1
            y1 = temp
        }
        for i:= y1; i <= y2; i++ {
            grid[i][x1] += 1
        }
        return grid
    }
    if y1 == y2 { // Horizontal Case
        // Swap around params
        if x2 < x1 {
            temp := x2
            x2 = x1
            x1 = temp
        }
        for i:= x1; i <= x2; i++ {
            grid[y1][i] += 1
        }
        return grid
    }
    // Diagonal case
    dirX := -1
    if x1 < x2 {
        dirX = 1
    }
    dirY := -1
    if y1 < y2 {
        dirY = 1
    }
    delta := x1 - x2
    if delta < 0 {
        delta = -delta
    }
    for i:=0; i <= delta; i++ {
        grid[y1+i*dirY][x1+i*dirX] += 1
    }

    return grid
}
