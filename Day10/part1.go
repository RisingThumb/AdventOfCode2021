package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "sort"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
    incompleteLines := []string{}
    corruptedLines := []string{}
    score := 0
    scoreValues := map[string]int{
        ")": 3,
        "]": 57,
        "}": 1197,
        ">": 25137,
    }
    closingBracket := map[string]string{
        ")": "(",
        "}": "{",
        ">": "<",
        "]": "[",
    }

	for scanner.Scan() {
        line := scanner.Text()
        charStack := []string{}
        isCorrupted := false
        for _, c := range line {
            char := string(c)
            switch char {
            case "(", "[", "{", "<" :
                charStack = append(charStack, char)
            case "}", ")", "]", ">":
                relevantChar := charStack[len(charStack)-1]
                charStack = charStack[:len(charStack)-1]
                if closingBracket[char] != relevantChar {
                    isCorrupted = true
                    score += scoreValues[char]
                    break
                }
            }
        }
        if isCorrupted {
            corruptedLines = append(corruptedLines, line)
        } else {
            incompleteLines = append(incompleteLines, line)
        }
	}
    fmt.Println("Part 1 answer: ", score)
    scoreValues2 := map[string]int {
        "(": 1,
        "[": 2,
        "{": 3,
        "<": 4,
    }
    scores := []int{}
    for _, incompleteLine := range incompleteLines {
        score2 := 0
        charStack := []string{}
        for _, c := range incompleteLine {
            char := string(c)
            switch char {
            case "(", "[", "{", "<" :
                charStack = append(charStack, char)
            case "}", ")", "]", ">":
                charStack = charStack[:len(charStack)-1]
            }
        }
        for i := range charStack {
            relevantChar := charStack[len(charStack)-1-i]
            score2 *= 5
            score2 += scoreValues2[relevantChar]
        }
        scores = append(scores, score2)
    }
    sort.Ints(scores)
    fmt.Println("Part 2 answer: ", scores[len(scores)/2])
}
