package main

import (
	"bufio"
	"bytes"
	"log"
	"math"
	"os"
	"slices"
)

const gridLength = 140
var MAS = []byte("MAS")
const A = byte('A')
const wLen = 3
var epicenter = int(math.Floor(float64(wLen) / 2.0))
var segmentBuffer = make([]byte, 3)

func main() {
	if wLen % 2 == 0 {
		log.Fatal("Can't make a simetrical cross with a word that has an even number of letters")
	}

	// Instanciate the matrix
	file, err := os.Open("../easyinput")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	matrix := make([][]byte, 0, gridLength)
	for scanner.Scan() {
		matrix = append(matrix, []byte(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("reading standard input: %q", err)
	}

	// count occurences
	var xmasOccurences int64 = 0
	for y, row := range(matrix[1:len(matrix)-1]) {
		for x, val := range row[1:len(row)-1] {
			if val == A && checkCross(x+1, y+1, matrix) {
				xmasOccurences += 1
			}
		}
	}

	log.Printf("Result: %v", xmasOccurences)
}

func checkCross(x int, y int, matrix [][]byte) bool {
	// NW <=> SE
	for i := range(wLen) {
		segmentBuffer[i] = matrix[y-epicenter+i][x-epicenter+i]
	}
	if !bytes.Equal(segmentBuffer, MAS) {
		slices.Reverse(segmentBuffer)
		if !bytes.Equal(segmentBuffer, MAS) {
			return false
		}
	}

	// NE <=> SW
	for i := range(wLen) {
		segmentBuffer[i] = matrix[y+epicenter-i][x-epicenter+i]
	}
	if !bytes.Equal(segmentBuffer, MAS) {
		slices.Reverse(segmentBuffer)
		if !bytes.Equal(segmentBuffer, MAS) {
			return false
		}
	}
	return true
}
