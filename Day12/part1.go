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
	newCaveMap := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		thisLine := scanner.Text()
		partsToAddToMap := strings.Split(thisLine, "-")
		_, ok := newCaveMap[partsToAddToMap[0]]
		if ok {
			newCaveMap[partsToAddToMap[0]] = append(newCaveMap[partsToAddToMap[0]], partsToAddToMap[1])
		} else {
			newCaveMap[partsToAddToMap[0]] = make([]string, 0, 1)
			newCaveMap[partsToAddToMap[0]] = append(newCaveMap[partsToAddToMap[0]], partsToAddToMap[1])
		}
		_, ok = newCaveMap[partsToAddToMap[0]]
		if ok {
			newCaveMap[partsToAddToMap[1]] = append(newCaveMap[partsToAddToMap[1]], partsToAddToMap[0])
		} else {
			newCaveMap[partsToAddToMap[1]] = make([]string, 0, 1)
			newCaveMap[partsToAddToMap[1]] = append(newCaveMap[partsToAddToMap[1]], partsToAddToMap[0])
		}
	}
	for k, v := range newCaveMap {
		fmt.Println(k, v)
		fmt.Println(isBigCave(k))
	}
	paths := [][]string{}
	paths = findAllPaths(newCaveMap, "start", []string{}, paths)
	for _, path := range paths {
		fmt.Println(path)
	}
	fmt.Println(len(paths))
	//paths := [][]string{}
	// Brute force keep visiting left-right in array order until running out of paths.
	// Convert the map into a tree?
}

func isBigCave(value string) bool {
	return strings.Compare(value, strings.ToUpper(value)) == 0
}

func findAllPaths(caveMap map[string][]string, currentCave string, visitedCaves []string, paths [][]string) [][]string {
	couldVisit := caveMap[currentCave]
	visitedCaves = append(visitedCaves, currentCave)
	// filter out any we have visited and are small
	if currentCave == "end" {
		return append(paths, visitedCaves)
	}
	couldVisitFiltered := []string{}
	for _, cave := range couldVisit {
		if !contains(cave, visitedCaves) {
			couldVisitFiltered = append(couldVisitFiltered, cave)
		}
	}
	if len(couldVisitFiltered) <= 0 {
		return paths
	}
	for _, cave := range couldVisitFiltered {
		paths = findAllPaths(caveMap, cave, visitedCaves, paths)
	}
	return paths
}

func contains(test string, visitedCaves []string) bool {
	for _, value := range visitedCaves {
		if strings.Compare(value, test) == 0 && !isBigCave(test) {
			return true
		}
	}
	return false
}
