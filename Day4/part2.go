package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "strings"
    "regexp"
	"strconv"
)

func main() {
    // Parsing
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
    scanner.Scan()
    line1 := strings.Split(scanner.Text(), ",")
    callouts := []int{}
    for _, value := range line1 {
        newCall, _ := strconv.ParseInt(value, 10, 32)
        callouts = append(callouts, int(newCall))
    }

    space := regexp.MustCompile(`\s+`)
    boards := [][][]int{}

    for scanner.Scan() {
        newText := strings.TrimSpace(space.ReplaceAllString(scanner.Text(), " "))
        if newText == "" {
            boards = append(boards, [][]int{})
            continue
        }
        newBoardLine := strings.Split(newText, " ")
        newLine := []int{}
        for _, value := range newBoardLine {
            newVal, _ := strconv.ParseInt(value, 10, 32)
            newLine = append(newLine, int(newVal))
        }
        boards[len(boards)-1] = append(boards[len(boards)-1], newLine)
    }
    index := -1
    mostRecentlyCalled := 0
    doneBoards := []int{}
    out:
    for _, call := range callouts {
        breakOut := 0
        mostRecentlyAdded := -1
        for i, board := range boards {
            markBoard(board, call)
            completed := checkBoard(board)
            if completed {
                breakOut += 1
                if !contains(doneBoards, i) {
                    doneBoards = append(doneBoards, i)
                    mostRecentlyAdded = i
                }
                if breakOut == len(boards) {
                    index = mostRecentlyAdded
                    mostRecentlyCalled = call
                    break out
                }
            }
        }
    }
    fmt.Println(index)
    fmt.Println(sumUnamrkedBoards(boards[index]) * mostRecentlyCalled)
    // We now have all the boards.
    // The way to do this, is to have a function called checkBoard, to check a board for success. a line of -1 is bingo
    // We will mark new values with -1 in a function called markBoard
}

func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func sumUnamrkedBoards(board[][]int) int {
    sum := 0
    for _, row := range board {
        for _, value := range row {
            if value == -1 {
                continue
            }
            sum += value
        }
    }
    return sum
}

func markBoard(board [][]int, mark int) {
    for i, row := range board {
        for j, value := range row {
            if value == mark {
                board[i][j] = -1
                return
            }
        }
    }
}

func checkBoard(board [][]int) bool {
    // Marked positions are -1
    for i:= 0; i< 5; i++ {
        // Row check
        if board[i][0] + board[i][1] + board[i][2] + board[i][3] + board[i][4] == -5 {
            return true
        }
        // Column check
        if board[0][i] + board[1][i] + board[2][i] + board[3][i] + board[4][i] == -5 {
            return true
        }
    }
    return false
}
