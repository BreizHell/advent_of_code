package main

import (
	"bufio"
	"log"
	"os"
)

type coordinate struct {
	x int
	y int
}

const wLen = 4
const gridLength = 140

func main() {
	XMAS := []byte("XMAS")

	// Instanciate the matrix
	file, err := os.Open("../input")
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
	for y := range len(matrix) {
		for x := range len(matrix[y]) {
			xmasOccurences += int64(checkEast(coordinate{x, y}, XMAS, matrix))
			xmasOccurences += int64(checkWest(coordinate{x, y}, XMAS, matrix))
			xmasOccurences += int64(checkSouth(coordinate{x, y}, XMAS, matrix))
			xmasOccurences += int64(checkNorth(coordinate{x, y}, XMAS, matrix))
			xmasOccurences += int64(checkSouthEast(coordinate{x, y}, XMAS, matrix))
			xmasOccurences += int64(checkNorthWest(coordinate{x, y}, XMAS, matrix))
			xmasOccurences += int64(checkNorthEast(coordinate{x, y}, XMAS, matrix))
			xmasOccurences += int64(checkSouthWest(coordinate{x, y}, XMAS, matrix))
		}
	}

	log.Printf("Result: %v", xmasOccurences)
}

func checkEast(coord coordinate, XMAS []byte, matrix [][]byte) int {
	if coord.x+(wLen-1) < len(matrix[coord.y]) {
		for i := 0; i < wLen; i++ {
			if XMAS[i] != matrix[coord.y][coord.x+i] {
				return 0
			}
		}
		return 1
	}
	return 0
}
func checkWest(coord coordinate, XMAS []byte, matrix [][]byte) int {
	if coord.x-(wLen-1) >= 0 {
		for i := 0; i < wLen; i++ {
			if XMAS[i] != matrix[coord.y][coord.x-i] {
				return 0
			}
		}
		return 1
	}
	return 0
}
func checkSouth(coord coordinate, XMAS []byte, matrix [][]byte) int {
	if coord.y+(wLen-1) < len(matrix) {
		for i := 0; i < wLen; i++ {
			if XMAS[i] != matrix[coord.y+i][coord.x] {
				return 0
			}
		}
		return 1
	}
	return 0
}
func checkNorth(coord coordinate, XMAS []byte, matrix [][]byte) int {
	if coord.y-(wLen-1) >= 0 {
		for i := 0; i < wLen; i++ {
			if XMAS[i] != matrix[coord.y-i][coord.x] {
				return 0
			}
		}
		return 1
	}
	return 0
}
func checkSouthEast(coord coordinate, XMAS []byte, matrix [][]byte) int {
	if coord.x+(wLen-1) < len(matrix[coord.y]) && coord.y+(wLen-1) < len(matrix) {
		for i := 0; i < wLen; i++ {
			if XMAS[i] != matrix[coord.y+i][coord.x+i] {
				return 0
			}
		}
		return 1
	}
	return 0
}
func checkNorthWest(coord coordinate, XMAS []byte, matrix [][]byte) int {
	if coord.x-(wLen-1) >= 0 && coord.y-(wLen-1) >= 0 {
		for i := 0; i < wLen; i++ {
			if XMAS[i] != matrix[coord.y-i][coord.x-i] {
				return 0
			}
		}
		return 1
	}
	return 0
}
func checkNorthEast(coord coordinate, XMAS []byte, matrix [][]byte) int {
	if coord.x+(wLen-1) < len(matrix[coord.y]) && coord.y-(wLen-1) >= 0 {
		for i := 0; i < wLen; i++ {
			if XMAS[i] != matrix[coord.y-i][coord.x+i] {
				return 0
			}
		}
		return 1
	}
	return 0
}
func checkSouthWest(coord coordinate, XMAS []byte, matrix [][]byte) int {
	if coord.x-(wLen-1) >= 0 && coord.y+(wLen-1) < len(matrix) {
		for i := 0; i < wLen; i++ {
			if XMAS[i] != matrix[coord.y+i][coord.x-i] {
				return 0
			}
		}
		return 1
	}
	return 0
}
