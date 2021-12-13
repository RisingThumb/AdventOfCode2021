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
    smallCaveMap := make(map[string]bool)
	for scanner.Scan() {
		thisLine := scanner.Text()
		partsToAddToMap := strings.Split(thisLine, "-")
		_, ok := newCaveMap[partsToAddToMap[0]]
		if ok {
			newCaveMap[partsToAddToMap[0]] = append(newCaveMap[partsToAddToMap[0]], partsToAddToMap[1])
            if !isBigCave(partsToAddToMap[0]) && partsToAddToMap[0] != "start" && partsToAddToMap[0] != "end" {
                smallCaveMap[partsToAddToMap[0]] = true
            }
		} else {
			newCaveMap[partsToAddToMap[0]] = make([]string, 0, 1)
			newCaveMap[partsToAddToMap[0]] = append(newCaveMap[partsToAddToMap[0]], partsToAddToMap[1])
		}
		_, ok = newCaveMap[partsToAddToMap[0]]
		if ok {
			newCaveMap[partsToAddToMap[1]] = append(newCaveMap[partsToAddToMap[1]], partsToAddToMap[0])
            if !isBigCave(partsToAddToMap[1]) && partsToAddToMap[1] != "start" && partsToAddToMap[1] != "end" {
                smallCaveMap[partsToAddToMap[1]] = true
            }
		} else {
			newCaveMap[partsToAddToMap[1]] = make([]string, 0, 1)
			newCaveMap[partsToAddToMap[1]] = append(newCaveMap[partsToAddToMap[1]], partsToAddToMap[0])
		}
	}
	paths := [][]string{}
    for k := range smallCaveMap {
        paths = findAllPaths(newCaveMap, "start", []string{}, paths, k, 0)
    }

	fmt.Println(len(paths))
}

func isBigCave(value string) bool {
	return strings.Compare(value, strings.ToUpper(value)) == 0
}

func findAllPaths(caveMap map[string][]string, currentCave string, visitedCaves []string, paths [][]string, caveToVisitTwice string, visitedCount int) [][]string {
	couldVisit := caveMap[currentCave]
    if currentCave == caveToVisitTwice {
        visitedCount += 1
    }
	visitedCaves = append(visitedCaves, currentCave)
    copyVisitedCaves := make([]string, len(visitedCaves))
    copy(copyVisitedCaves, visitedCaves)
	// filter out any we have visited and are small
	if currentCave == "end" {
        if doesPathsContainPath(paths, copyVisitedCaves) {
            return paths
        } else {
            return append(paths, copyVisitedCaves)
        }
	}
	couldVisitFiltered := []string{}
	for _, cave := range couldVisit {
        if visitedCount < 2 && caveToVisitTwice == cave {
			couldVisitFiltered = append(couldVisitFiltered, cave)
        } else if !contains(cave, copyVisitedCaves) {
			couldVisitFiltered = append(couldVisitFiltered, cave)
		}
	}
	for _, cave := range couldVisitFiltered {
		paths = findAllPaths(caveMap, cave, copyVisitedCaves, paths, caveToVisitTwice, visitedCount)
	}
	return paths
}

func doesPathsContainPath(paths [][]string, newpath []string) bool {
    for _, path := range paths {
        if arrayEquals(path, newpath) {
            return true
        }
    }
    return false
}

func arrayEquals(path1, path2 []string) bool {
    if len(path1) != len(path2) {
        return false
    }
    for i := range path1 {
        if path1[i] != path2[i] {
            return false
        }
    }
    return true
}

func contains(test string, visitedCaves []string) bool {
	for _, value := range visitedCaves {
		if value == test && !isBigCave(test) {
			return true
		}
	}
	return false
}
