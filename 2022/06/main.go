package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getComms(filename string) string {
	readFile, err := os.Open(filename)
	defer readFile.Close()
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	if !fileScanner.Scan() {
		panic("no input!")
	}

	return fileScanner.Text()
}

func isStartOfPacket(comms string) bool {
	for idx, char := range comms {
		temp := comms[:idx] + comms[idx+1:]
		if strings.ContainsRune(temp, char) {
			return false
		}
	}

	return true
}

func getStartOfPacketWithSizeN(comms string, n int) int {
	nPackets := len(comms) - n + 1

	for i := 0; i < nPackets; i++ {
		packet := comms[i : i+n]
		if isStartOfPacket(packet) {
			return i + n
		}
	}

	return -1
}

func main() {
	input := os.Args[1]

	comms := getComms(input)
	start := getStartOfPacketWithSizeN(comms, 4)
	fmt.Println(start)

	start = getStartOfPacketWithSizeN(comms, 14)
	fmt.Println(start)
}
