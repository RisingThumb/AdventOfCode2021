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
	file, err := os.Open("./test.txt")
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
	paths := [][]string{}
	paths = findAllPaths(newCaveMap, "start", []string{}, paths, "", 0)

	paths = removeDuplicateArrays(paths)
	for i, path := range paths {
		fmt.Println(i, path)
	}
	fmt.Println(len(paths))
	//paths := [][]string{}
	// Brute force keep visiting left-right in array order until running out of paths.
	// Convert the map into a tree?
}

func isBigCave(value string) bool {
	return strings.Compare(value, strings.ToUpper(value)) == 0
}

func removeDuplicateArrays(paths [][]string) [][]string {
	indicesToDelete := make(map[int]bool)
	for i, _ := range paths {
		for j := i + 1; j < len(paths); j++ {
			if checkArraysAreEqual(paths[i], paths[j]) {
				indicesToDelete[j] = true
			}
		}
	}
	newIndicesToDelete := []int{}
	for key, _ := range indicesToDelete {
		newIndicesToDelete = append(newIndicesToDelete, key)
	}
	sort.Ints(newIndicesToDelete)
	fmt.Println(newIndicesToDelete)
	fmt.Println(len(newIndicesToDelete))
	fmt.Println(len(paths))
	for i := len(newIndicesToDelete) - 1; i >= 0; i-- {
		indexToDelete := newIndicesToDelete[i]
		paths = append(paths[:indexToDelete], paths[indexToDelete+1:]...)
	}

	return paths
}

func checkArraysAreEqual(path1, path2 []string) bool {
	if len(path1) != len(path2) {
		return false
	}

	for i := 0; i < len(path1); i++ {
		if path1[i] != path2[i] {
			return false
		}
	}
	return true
}

func findAllPaths(caveMap map[string][]string, currentCave string, visitedCaves []string, paths [][]string, caveThatCanBeRevisited string, visitedCount int) [][]string {
	if currentCave == "end" {
		return append(paths, append(visitedCaves, "end"))
	}
	couldVisit := caveMap[currentCave]
	couldVisitFiltered := []string{}

	for _, cave := range couldVisit {
		if !contains(cave, visitedCaves, caveThatCanBeRevisited, visitedCount) {
			couldVisitFiltered = append(couldVisitFiltered, cave)
		}
	}

	if currentCave == caveThatCanBeRevisited {
		visitedCount += 1
	}
	for _, cave := range couldVisitFiltered {
		paths = findAllPaths(caveMap, cave, append(visitedCaves, currentCave), paths, caveThatCanBeRevisited, visitedCount)
		if caveThatCanBeRevisited == "" && (visitedCount == 0) && currentCave != "end" && currentCave != "start" && !isBigCave(currentCave) {
			paths = findAllPaths(caveMap, cave, append(visitedCaves, currentCave), paths, currentCave, visitedCount+1)
		}
	}
	return paths
}

func contains(test string, visitedCaves []string, caveThatCanBeRevisited string, visitedCount int) bool {
	for _, value := range visitedCaves {
		if strings.Compare(value, test) == 0 && !isBigCave(test) {
			if value == caveThatCanBeRevisited && visitedCount < 2 {
				return false
			}
			return true
		}
	}
	return false
}
