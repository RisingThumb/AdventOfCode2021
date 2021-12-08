package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	sum := 0
	for _, line := range inputs {
		lettersInNumber := [10]string{}
		sort.Slice(line[:10], func(i, j int) bool {
			return len(line[i]) < len(line[j])
		})
		// Mapping all the sequences to numbers
		for _, sequence := range line[:10] {
			switch len(sequence) {
			case 2:
				lettersInNumber[1] = sequence
			case 3:
				lettersInNumber[7] = sequence
			case 4:
				lettersInNumber[4] = sequence
			case 5:
				if patternOverlap(lettersInNumber[1], sequence, 2) {
					lettersInNumber[3] = sequence
				} else if patternOverlap(lettersInNumber[4], sequence, 2) {
					lettersInNumber[2] = sequence
				} else {
					lettersInNumber[5] = sequence
				}
			case 6:
				if patternOverlap(lettersInNumber[4], sequence, 4) {
					lettersInNumber[9] = sequence
				} else if patternOverlap(lettersInNumber[1], sequence, 2) {
					lettersInNumber[0] = sequence
				} else {
					lettersInNumber[6] = sequence
				}
			case 7:
				lettersInNumber[8] = sequence
			}
		}
		sum += getOutput(lettersInNumber, line[10:])
	}
	fmt.Println(sum)
}

func getOutput(sequences [10]string, output []string) int {
	digits := 0
	for _, val := range output {
		for i := 0; i < 10; i++ {
			if patternEquals(sequences[i], val) {
				digits *= 10
				digits += i
				break
			}
		}
	}
	return digits
}

func patternEquals(sequence1, sequence2 string) bool {
	if len(sequence1) != len(sequence2) {
		return false
	}
	for _, val1 := range sequence1 {
		isIn := false
		for _, val2 := range sequence2 {
			isIn = isIn || val1 == val2
		}
		if !isIn {
			return false
		}
	}
	return true
}

func patternOverlap(sequenceToCheck, sequence string, overlapCount int) bool {
	count := 0
	for _, val := range sequence {
		for _, val2 := range sequenceToCheck {
			if val == val2 {
				count += 1
				break
			}
		}
	}
	return overlapCount == count
}
