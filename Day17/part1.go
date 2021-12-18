package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "strings"
    "strconv"

)

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
    scanner.Scan()
    line := scanner.Text()
    line = strings.Replace(line, "target area: ", "", -1)
    parts := strings.Split(line, ", ")
    xPart :=  strings.Split( strings.Replace(parts[0], "x=", "", -1) , "..")
    yPart :=  strings.Split( strings.Replace(parts[1], "y=", "", -1) , "..")
    yPartInt := []int{}
    for _, yValue := range yPart {
        yInt, _ := strconv.ParseInt(yValue, 10, 64)
        yPartInt = append(yPartInt, int(yInt))
    }
    xPartInt := []int{}
    for _, xValue := range xPart {
        xInt, _ := strconv.ParseInt(xValue, 10, 64)
        xPartInt = append(xPartInt, int(xInt))
    }
	fmt.Println(xPart, yPart)
    initialVelocity := 5000
    highestYPosition := 0
    yMin := min(yPartInt[0], yPartInt[1])
    yMax := max(yPartInt[0], yPartInt[1])
    for i := 0; i < initialVelocity; i++ {
        success, highestY := isInRangeForVelocity(i, yMin, yMax)
        if success {
            highestYPosition = max(highestYPosition, highestY)
        }
    }
    fmt.Println("Part 1 answer : ", highestYPosition)

    xMin := min(xPartInt[0], xPartInt[1])
    xMax := max(xPartInt[0], xPartInt[1])
    fmt.Println(xMin, xMax)
    allXInitialVelocities := []int{}
    for i := -initialVelocity; i < initialVelocity; i++ {
        success := isInRangeForXVelocity(i, xMin, xMax)
        if success {
            allXInitialVelocities = append(allXInitialVelocities, i)
        }
    }

    allYInitialVelocities := []int{}
    for i := -initialVelocity; i < initialVelocity; i++ {
        success := isInRangeForYVelocity(i, yMin, yMax)
        if success {
            allYInitialVelocities = append(allYInitialVelocities, i)
        }
    }

    // Generate all Possible velocity pairs then test each one
    velocityPairs := [][]int{}
    for _, xVel := range allXInitialVelocities {
        for _, yVel := range allYInitialVelocities {
            if testVelocityPair(xVel, yVel, yMin, yMax, xMin, xMax) {
                velocityPairs = append(velocityPairs, []int{xVel, yVel})
            }
        }
    }
    fmt.Println("Part 2 answer : ", len(velocityPairs))
}

func testVelocityPair(xVel, yVel, yMin, yMax, xMin, xMax int) (bool) {
    x := 0
    y := 0
    currentXVel := xVel
    currentYVel := yVel
    wentInsideRange := false
    for y >= yMin*2 {
        y += currentYVel
        x += currentXVel
        if currentXVel > 0 {
            currentXVel -= 1
        }
        currentYVel -= 1
        if y >= yMin && y <= yMax && x >= xMin && x <= xMax {
            wentInsideRange = true
        }
    }

    return wentInsideRange
}

// We'll just brute force it. (part 1 func since it also wants highestY)
func isInRangeForVelocity(initialVelocity, yMin, yMax int) (bool, int) {
    currentY := 0
    currentVelocity := initialVelocity
    highestYPosition := 0
    wentInsideRange := false
    for currentY >= yMin {
        currentY += currentVelocity
        currentVelocity -= 1
        highestYPosition = max(highestYPosition, currentY)
        if currentY >= yMin && currentY <= yMax {
            wentInsideRange = true
        }
    }
    return wentInsideRange, highestYPosition
}



// We'll just brute force it.
func isInRangeForYVelocity(initialVelocity, yMin, yMax int) (bool) {
    currentY := 0
    currentVelocity := initialVelocity
    wentInsideRange := false
    for currentY >= yMin {
        currentY += currentVelocity
        currentVelocity -= 1
        if currentY >= yMin && currentY <= yMax {
            wentInsideRange = true
        }
    }
    return wentInsideRange
}


// We'll just brute force it.
func isInRangeForXVelocity(initialVelocity, xMin, xMax int) (bool) {
    currentX := 0
    currentVelocity := initialVelocity
    wentInsideRange := false
    for currentX <= xMax && currentVelocity > 0 {
        currentX += currentVelocity
        currentVelocity -= 1
        if currentX >= xMin && currentX <= xMax {
            wentInsideRange = true
        }
    }
    return wentInsideRange
}
