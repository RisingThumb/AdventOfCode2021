package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type bitCount struct {
	ONE  uint32
	ZERO uint32
}

func calculateBitcounts(binValues []uint32, length int) []bitCount {
	newBitCounts := []bitCount{}
	for i := 0; i < length; i++ {
		newBitCounts = append(newBitCounts, bitCount{0, 0})
	}
	for _, value := range binValues {
		//fmt.Println(i)
		// All binary values are 12 bits long
		for j := 0; j < length; j++ {
			bitOfInterest := value & 0x01
			if bitOfInterest == 0x01 {
				newBitCounts[length-1-j].ONE += 1
			} else {
				newBitCounts[length-1-j].ZERO += 1
			}
			value = value >> 1
		}
	}
	return newBitCounts
}

func getMeGasValue(co2 bool, bitCounts []bitCount, binValuesCpy []uint32) uint32 {
	for i := 0; i < len(bitCounts); i++ {
		// We need to update them as the size will change over time
		bitCounts = calculateBitcounts(binValuesCpy, len(bitCounts))

		relevantBitsToLookAt := bitCounts[i]
		var filter uint32 = 0
		if co2 && relevantBitsToLookAt.ONE < relevantBitsToLookAt.ZERO {
			filter = 1
		} else if !co2 && relevantBitsToLookAt.ONE >= relevantBitsToLookAt.ZERO {
			filter = 1
		}
		indicesToRemove := []int{}
		for j, value := range binValuesCpy {
			bitShiftedValue := value >> uint32(len(bitCounts)-1-i)
			bitShiftedValue &= 1
			// We now have the value. Compare it to see if it needs removal
			if bitShiftedValue != filter {
				indicesToRemove = append(indicesToRemove, j)
			}
		}
		for j, value := range indicesToRemove {
			if value-j == len(binValuesCpy) {
				binValuesCpy = binValuesCpy[:value-j]
			} else {
				binValuesCpy = append(binValuesCpy[:value-j], binValuesCpy[value-j+1:]...)
			}
		}
		if len(binValuesCpy) == 1 {
			break
		}

	}
	return binValuesCpy[0]
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	binValues := []uint32{}
	for scanner.Scan() {
		newInstruction := scanner.Text()
		value, _ := strconv.ParseInt(newInstruction, 2, 32)
		uValue := uint32(value)
		binValues = append(binValues, uValue)
	}
	bitCounts := []bitCount{}
	for i := 0; i < 12; i++ {
		bitCounts = append(bitCounts, bitCount{0, 0})
	}
	binValuesCpy := make([]uint32, len(binValues))
	binValuesCpy2 := make([]uint32, len(binValues))
	copy(binValuesCpy, binValues)
	copy(binValuesCpy2, binValues)

	oxyValue := getMeGasValue(false, bitCounts, binValuesCpy)
	co2Value := getMeGasValue(true, bitCounts, binValuesCpy2)
	fmt.Println(oxyValue)
	fmt.Println(co2Value)
	fmt.Println(oxyValue * co2Value)
}
