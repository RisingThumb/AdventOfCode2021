package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
    "math"
)

type snailFish struct {
    parentElement *snailFish
    leftElement *snailFish
    rightElement *snailFish
    isLiteralValue bool
    literalValue int
}

func newSnailFish() *snailFish {
    s := snailFish{nil, nil, nil, false, -1}
    return &s
}

func (snailFish *snailFish) SetLiteralValue(literalValue int) {
    snailFish.leftElement = nil
    snailFish.rightElement = nil
    snailFish.isLiteralValue = true
    snailFish.literalValue = literalValue
}


func (snailFish *snailFish) SetParent(parentElement *snailFish) {
    snailFish.parentElement = parentElement
}

func (snailFish *snailFish) SetLeftElement(leftElement *snailFish) {
    snailFish.leftElement = leftElement
    snailFish.isLiteralValue = false
}

func (snailFish *snailFish) SetRightElement(rightElement *snailFish) {
    snailFish.rightElement = rightElement
    snailFish.isLiteralValue = false
}

func (snailFish *snailFish) PrintS() {
    snailFish.Print()
    fmt.Println()
}

func (snailFish *snailFish) Print() {
    if !snailFish.isLiteralValue {
        fmt.Printf("[")
        //fmt.Print(snailFish.parentElement)
    }
    if snailFish.leftElement != nil {
        snailFish.leftElement.Print()
    }
    if snailFish.isLiteralValue {
        fmt.Printf(" %d ", snailFish.literalValue)
        //fmt.Print(snailFish.parentElement)
    }
    if snailFish.rightElement != nil {
        snailFish.rightElement.Print()
    }
    if !snailFish.isLiteralValue {
        fmt.Printf("]",)
        //fmt.Print(snailFish.parentElement)
    }
}

func addSnailFish(leftSnailFish, rightSnailFish *snailFish) *snailFish {
    snailFish := newSnailFish()
    snailFish.SetLeftElement(leftSnailFish)
    leftSnailFish.SetParent(snailFish)
    rightSnailFish.SetParent(snailFish)
    snailFish.SetRightElement(rightSnailFish)
    return snailFish
}

func (snailFish *snailFish) AddLiteralToLeft(leftLiteral int) {
    currentNode := snailFish.parentElement
    previousNode := snailFish
    for currentNode.leftElement == previousNode && currentNode.parentElement != nil {
        previousNode = currentNode
        currentNode = currentNode.parentElement
    }
    if currentNode.parentElement == nil && currentNode.leftElement == previousNode {
        return
    }
    currentNode = currentNode.leftElement
    for !currentNode.isLiteralValue {
        currentNode = currentNode.rightElement
    }
    currentNode.SetLiteralValue(currentNode.literalValue + leftLiteral)
}

func (snailFish *snailFish) AddLiteralToRight(rightLiteral int) {
    currentNode := snailFish.parentElement
    previousNode := snailFish
    for currentNode.rightElement == previousNode && currentNode.parentElement != nil {
        previousNode = currentNode
        currentNode = currentNode.parentElement
    }
    if currentNode.parentElement == nil && currentNode.rightElement == previousNode {
        return
    }
    currentNode = currentNode.rightElement
    for !currentNode.isLiteralValue {
        currentNode = currentNode.leftElement
    }
    currentNode.SetLiteralValue(currentNode.literalValue + rightLiteral)
}

func (snailFish *snailFish) Explode(depth int) bool {
    if depth > 4 && !snailFish.isLiteralValue {
        // Seek for left most element and put it on
        snailFish.AddLiteralToLeft(snailFish.leftElement.literalValue)
        snailFish.AddLiteralToRight(snailFish.rightElement.literalValue)
        snailFish.SetLiteralValue(0)

        return true
    }
    if !snailFish.isLiteralValue {
        if snailFish.leftElement.Explode(depth+1) {
            return true
        }
        if snailFish.rightElement.Explode(depth+1) {
            return true
        }
    }
    return false
}

func (snailFish *snailFish) Split() bool {
    if snailFish.isLiteralValue && snailFish.literalValue > 9 {
        newLeftSnailFish := newSnailFish()
        newRightSnailFish := newSnailFish()
        valueToRound := float64(snailFish.literalValue)/2

        newLeftSnailFish.SetLiteralValue(int(math.Floor(valueToRound)))
        newLeftSnailFish.SetParent(snailFish)
        newRightSnailFish.SetLiteralValue(int(math.Ceil(valueToRound)))
        newRightSnailFish.SetParent(snailFish)
        snailFish.SetLeftElement(newLeftSnailFish)
        snailFish.SetRightElement(newRightSnailFish)
        return true
    }
    if snailFish.leftElement != nil && snailFish.leftElement.Split() {
        return true
    }
    if snailFish.rightElement != nil && snailFish.rightElement.Split() {
        return true
    }
    return false
}

func (snailFish *snailFish) Reduce() {
    for true {
        
        if snailFish.Explode(1) {
            continue
        }
        if snailFish.Split() {
            continue
        }
        break
    }
}

func (snailFish *snailFish) DeepCopy(parent *snailFish) *snailFish {
    newSnailFish := newSnailFish()
    newSnailFish.literalValue = snailFish.literalValue
    newSnailFish.SetParent(parent)
    if snailFish.leftElement != nil {
        newSnailFish.leftElement = snailFish.leftElement.DeepCopy(newSnailFish)
    }

    if snailFish.rightElement != nil {
        newSnailFish.rightElement = snailFish.rightElement.DeepCopy(newSnailFish)
    }
    newSnailFish.isLiteralValue = snailFish.isLiteralValue
    return newSnailFish
}

func (snailFish *snailFish) Magnitude() int {
    if snailFish.isLiteralValue {
        return snailFish.literalValue
    }
    return snailFish.leftElement.Magnitude()*3 + snailFish.rightElement.Magnitude()*2
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
    snailFishArray := []*snailFish{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        newLine := scanner.Text()
        masterSnailFish := true
        snailFishStack := []*snailFish{}
        for _, char := range newLine {
            switch string(char) {
                case "[":
                    snailFish := newSnailFish()
                    if masterSnailFish {
                        snailFishArray = append(snailFishArray, snailFish)
                        masterSnailFish = false
                    } else {
                        snailFish.SetParent(snailFishStack[len(snailFishStack)-1])
                    }
                    snailFishStack = append(snailFishStack, snailFish)
                case "]":
                    poppedFish := snailFishStack[len(snailFishStack)-1]
                    snailFishStack = snailFishStack[:len(snailFishStack)-1]
                    snailFishStack[len(snailFishStack)-1].SetRightElement(poppedFish)
                case ",":
                    poppedFish := snailFishStack[len(snailFishStack)-1]
                    snailFishStack = snailFishStack[:len(snailFishStack)-1]
                    snailFishStack[len(snailFishStack)-1].SetLeftElement(poppedFish)
                default: // This is the fallthrough case. Anything
                    snailFish := newSnailFish()
                    snailFish.SetParent(snailFishStack[len(snailFishStack)-1])
                    snailFishStack = append(snailFishStack, snailFish)
                    snailFishValue, _ := strconv.ParseInt(string(char), 10, 64)
                    snailFish.SetLiteralValue(int(snailFishValue))
            }
        }
	}
    snailFish := snailFishArray[0].DeepCopy(nil)
    for i := 1; i< len(snailFishArray); i++ {
        snailFish = addSnailFish(snailFish, snailFishArray[i].DeepCopy(nil))
        snailFish.Reduce()
    }
    fmt.Println("Part 1 solution : ", snailFish.Magnitude())
    largestMagnitude := 0
    for i := 0; i < len(snailFishArray); i++ {
        for j := 0; j < len(snailFishArray); j++ {
            ds1 := snailFishArray[i].DeepCopy(nil)
            ds2 := snailFishArray[j].DeepCopy(nil)
            ds3 := snailFishArray[i].DeepCopy(nil)
            ds4 := snailFishArray[j].DeepCopy(nil)
            addFish1 := addSnailFish(ds1, ds2)
            addFish2 := addSnailFish(ds4, ds3)
            addFish1.Reduce()
            addFish2.Reduce()
            largestMagnitude = max(largestMagnitude, addFish1.Magnitude())
            largestMagnitude = max(largestMagnitude, addFish2.Magnitude())
        }
    }
    fmt.Println("Part 2 solution : ", largestMagnitude)

}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
