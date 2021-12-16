package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
    Sum     int = 0
    Product     = 1
    Minimum     = 2
    Maximum     = 3
    Literal     = 4
    GreaterThan = 5
    LessThan    = 6
    EqualTo     = 7

)

type Packet struct {
    version int
    typeID int
    literalValue int
    subPackets []Packet
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
    scanner.Scan();
    hexString := scanner.Text()
    binaryString := ""
    for _, char := range hexString {
        newIntToAdd, _ := strconv.ParseInt(string(char), 16, 64);
        binaryString += fmt.Sprintf("%04s",strconv.FormatInt(newIntToAdd, 2))
    }

    fp, _ := parsePacket(binaryString)
    fmt.Println(fp.VersionSum(), fp.ComputePacket(), fp.literalValue)

}

func (packet *Packet) SetLiteralValue() {
    packet.literalValue = -1
}

func (packet Packet) PrintPacket() {
    fmt.Println(packet.version, packet.typeID, packet.literalValue, packet.subPackets)
}

func (packet *Packet) ComputePacket() int {
    returnVal := -1
    switch(packet.typeID) {
    case Sum:
        sumValue := 0
        for _, subPacket := range packet.subPackets {
            sumValue += subPacket.ComputePacket()
        }
        returnVal = sumValue
    case Product:
        multiplyValue := 1
        for _, subPacket := range packet.subPackets {
            multiplyValue *= subPacket.ComputePacket()
        }
        returnVal = multiplyValue
    case Minimum:
        startingLowest := packet.subPackets[0].ComputePacket()
        for _, subPacket := range packet.subPackets {
            next := subPacket.ComputePacket()
            if next < startingLowest {
                startingLowest = next
            }
        }
        returnVal = startingLowest
    case Maximum:
        startingHighest := packet.subPackets[0].ComputePacket()
        for _, subPacket := range packet.subPackets {
            next := subPacket.ComputePacket()
            if next > startingHighest {
                startingHighest = next
            }
        }
        returnVal = startingHighest
    case Literal:
        return packet.literalValue
    case GreaterThan:
        value := 0
        if packet.subPackets[0].ComputePacket() > packet.subPackets[1].ComputePacket(){
            value = 1
        }
        returnVal = value
    case LessThan:
        value := 0
        if packet.subPackets[0].ComputePacket() < packet.subPackets[1].ComputePacket() {
            value = 1
        }
        returnVal = value
    case EqualTo:
        value := 0
        if packet.subPackets[0].ComputePacket() == packet.subPackets[1].ComputePacket() {
            value = 1
        }
        returnVal = value
    }
    fmt.Println(packet.typeID, returnVal)
    packet.literalValue = returnVal
    return returnVal
}

func (packet Packet) VersionSum() int {
    versionSum := packet.version
    for _, packet := range packet.subPackets {
        versionSum += packet.VersionSum()
    }
    return versionSum
}

func parsePacket(packetString string) (Packet, int) {
    _, typeIDPart := parseHeader(packetString)

    switch {
        case int(typeIDPart) == Literal:
            return parseLiteralPacket(packetString)
        default:
            return parseOperatorPacket(packetString)
    }
}

func parseHeader(packetString string) (int64, int64) {
    versionPart, _ := strconv.ParseInt(packetString[:3], 2, 64)
    typeIDPart, _ := strconv.ParseInt(packetString[3:6], 2, 64)
    return versionPart, typeIDPart
}

func parseLiteralPacket(packetString string) (Packet, int) {
    versionPart, typeIDPart := parseHeader(packetString)
    // We expect the packet String to *only* start at the version/typeID part
    i := 6
    literalPartBitString := ""
    for packetString[i] != '0' {
        literalPartBitString += packetString[i+1:i+5]
        i += 5
    }
    literalPartBitString += packetString[i+1:i+5]
    literalPart, _ := strconv.ParseInt(literalPartBitString, 2, 64)
    return Packet{int(versionPart), int(typeIDPart), int(literalPart), []Packet{}}, i+5
}

func parseOperatorPacket(packetString string) (Packet, int) { // Also need to count number of bits used in it.
    versionPart, typeIDPart := parseHeader(packetString)
    lengthTypeId := packetString[6:7]
    subPacketArr := []Packet{}
    if lengthTypeId == "0" {
        // Next 15 bits represent total length in bits of the sub-packets contained by this packet
        lengthInBits, _ := strconv.ParseInt(packetString[7: 22], 2, 64)
        i := 0
        for i < int(lengthInBits) - 1{
            newPacket, newAddition := parsePacket(packetString[22+i:])
            subPacketArr = append(subPacketArr, newPacket)
            i += newAddition
        }
        return Packet{int(versionPart), int(typeIDPart), 0, subPacketArr}, int(lengthInBits)+22

    } else {
        // Next 11 bits are a number representing number of subpackets immediately contained by this packet
        numberOfSubPackets, _ := strconv.ParseInt(packetString[7: 18], 2, 64)
        i := 0
        bitsTravelled := 18
        for ;i < int(numberOfSubPackets); i++ {
            newPacket, bitAdjust := parsePacket(packetString[bitsTravelled:])
            bitsTravelled += bitAdjust
            subPacketArr = append(subPacketArr, newPacket)
        }
        return Packet{int(versionPart), int(typeIDPart), 0, subPacketArr}, bitsTravelled
    }
}














