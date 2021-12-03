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

	for _, value := range binValues {
		//fmt.Println(i)
		// All binary values are 12 bits long
		for j := 0; j < len(bitCounts); j++ {
			bitOfInterest := value & 0x01
			if bitOfInterest == 0x01 {
				bitCounts[j].ONE += 1
			} else {
				bitCounts[j].ZERO += 1
			}
			value = value >> 1
		}
	}
	epsilon := 0
	gamma := 0
	for i, _ := range bitCounts {
		value := bitCounts[len(bitCounts)-1-i]
		if value.ONE > value.ZERO {
			epsilon = epsilon | 1
		} else {
			gamma = gamma | 1
		}
		if i < len(bitCounts)-1 {
			epsilon = epsilon << 1
			gamma = gamma << 1
		}
	}
	fmt.Println(epsilon * gamma)
}
