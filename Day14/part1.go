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
	scanner.Scan()
	elementsLine := scanner.Text()
	scanner.Scan()
	elementsMapping := make(map[string]string)

	for scanner.Scan() {
		newElement := scanner.Text()
		newElementSplit := strings.Split(newElement, " -> ")
		elementsMapping[newElementSplit[0]] = newElementSplit[1]
	}
	/*
		fmt.Println(elementsLine)
		for k, v := range elementsMapping {
			fmt.Println(k, v)
		}
	*/
	elementPairMap, charMap := convertElementsIntoPairMap(elementsLine)

	for i := 0; i < 10; i++ {
		elementPairMap, charMap = createNewElement(elementPairMap, elementsMapping, charMap)
	}
	highest, lowest := countChars(charMap)
	fmt.Println(highest - lowest)
	for i := 0; i < 30; i++ {
		elementPairMap, charMap = createNewElement(elementPairMap, elementsMapping, charMap)
	}
	highest, lowest = countChars(charMap)
	fmt.Println(highest - lowest)
}

func convertElementsIntoPairMap(element string) (map[string]int, map[string]int) {
	charMap := make(map[string]int)
	elementPairMap := make(map[string]int)
	charMap[string(element[0])] += 1

	for i := 0; i < len(element)-1; i++ {
		relevantChars := element[i : i+2]
		charMap[string(element[i+1])] += 1
		elementPairMap[relevantChars] += 1
	}
	return elementPairMap, charMap
}

func createNewElement(elementPairMap map[string]int, elementMapping map[string]string, charMap map[string]int) (map[string]int, map[string]int) {
	copiedElementPairMap := make(map[string]int)
	for k, v := range elementPairMap {
		copiedElementPairMap[k] = v
	}
	for k, v := range elementPairMap {
		newComponent := elementMapping[k]
		newPair := string(k[0]) + newComponent
		newPair2 := newComponent + string(k[1])
		charMap[newComponent] += v
		copiedElementPairMap[newPair] += v
		copiedElementPairMap[newPair2] += v
		copiedElementPairMap[k] -= v
	}
	return copiedElementPairMap, charMap
}

func countChars(charMap map[string]int) (int, int) {
	highestKey := ""
	lowestKey := ""
	for k, value := range charMap {
		if lowestKey == "" || highestKey == "" {
			lowestKey = k
			highestKey = k
			continue
		}
		if charMap[lowestKey] > value {
			lowestKey = k
		}
		if charMap[highestKey] < value {
			highestKey = k
		}
	}
	return charMap[highestKey], charMap[lowestKey]
}
