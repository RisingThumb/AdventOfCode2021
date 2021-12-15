package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
    "sort"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
    grid := [][]int{}
	for scanner.Scan() {
        newLine := scanner.Text()
        row := []int{}
        for _, char := range newLine {
            value, _ := strconv.ParseInt(string(char), 10, 32)
            row = append(row, int(value))
        }
        grid = append(grid, row)
	}
    fmt.Println(computeShortestDistance(grid))
    newMegaGrid := [][]int{}
    for i := 0; i < 5; i++ {
        for _, line := range grid {
            newMegaLine := []int{}
            for j := 0; j < 5; j++ {
                for _, value := range line {
                    newMegaLine = append(newMegaLine, (value + i + j -1) % 9 + 1)
                }
            }
            newMegaGrid = append(newMegaGrid, newMegaLine)
        }
    }
    fmt.Println(computeShortestDistance(newMegaGrid))
    /*
    0 1 2 3 4
    1 2 3 4 5
    2 3 4 5 6
    3 4 5 6 7
    4 5 6 7 8
    */
}

func computeShortestDistance(grid [][]int) int{
    visitedMap := [][]bool{}
    for _, line := range grid {
        distanceRow := []int{}
        visitedRow := []bool{}
        for range line {
            visitedRow = append(visitedRow, false)
            distanceRow = append(distanceRow, -1)
        }
        visitedMap = append(visitedMap, visitedRow)
    }
    visitedMap[0][0] = true
    queue := [][]int{{0, 0, 0}, }
    for len(queue) > 0 {
        queuePop := queue[0]
        queue = queue[1:]
        if queuePop[1] == len(grid)-1 && queuePop[0] == len(grid[0])-1 {
            return queuePop[2]
        }

        // moving up
        if (queuePop[1] > 0 && !visitedMap[ queuePop[1]-1][queuePop[0]]) {
            queue = append(queue,
                        []int{queuePop[0],
                            queuePop[1]-1,
                            queuePop[2] + grid[ queuePop[1]-1 ][queuePop[0]],
                        },
                    )
            visitedMap[queuePop[1]-1][queuePop[0]] = true
        }

        // moving down
        if (queuePop[1] < len(grid) - 1  && !visitedMap[ queuePop[1]+1][queuePop[0]]) {
            queue = append(queue,
                        []int{queuePop[0],
                            queuePop[1]+1,
                            queuePop[2] + grid[ queuePop[1]+1 ][queuePop[0]],
                        },
                    )
            visitedMap[queuePop[1]+1][queuePop[0]] = true
        }
        // moving left
        if (queuePop[0] > 0  && !visitedMap[ queuePop[1]][queuePop[0] - 1]) {
            queue = append(queue,
                        []int{queuePop[0] - 1,
                            queuePop[1],
                            queuePop[2] + grid[ queuePop[1]][queuePop[0] - 1],
                        },
                    )
            visitedMap[queuePop[1]][queuePop[0]- 1] = true
        }

        // moving right
        if (queuePop[0] < len(grid[0]) - 1  && !visitedMap[ queuePop[1]][queuePop[0] + 1]) {
            queue = append(queue,
                        []int{queuePop[0] + 1,
                            queuePop[1],
                            queuePop[2] + grid[ queuePop[1]][queuePop[0] + 1],
                        },
                    )
            visitedMap[queuePop[1]][queuePop[0] + 1] = true
        }
        sort.SliceStable(queue, func(i, j int) bool {
            return queue[i][2] < queue[j][2]
        })
    }
    return -1
}

