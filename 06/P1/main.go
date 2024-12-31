package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var tilesSteppedOnto int64 = 1
	matrix := make([][]rune, 0, 131)
	var posX = 0
	var posY = 0

	scanner := bufio.NewScanner(file)

	y := 0
	for scanner.Scan() {
		line := []rune(scanner.Text())
		matrix = append(matrix, line)
		for x, tile := range line {
			if tile == '^' || tile == '>' || tile == 'v' || tile == '<' {
				posX, posY = x, y
			}
		}
		y++
	}

	play(matrix, &tilesSteppedOnto, posX, posY)

	log.Printf("Tiles stepped on: %v", tilesSteppedOnto)
}

func play(matrix [][]rune, tilesSteppedOnto *int64, startX int, startY int) {
	var forwardTile rune
	posX, posY := startX, startY
	defer func() {
		matrix[posY][posX] = 'X'
	}()
	for {
		fromX, fromY := posX, posY
		defer func() {
			if matrix[fromY][fromX] != 'X' {
				*tilesSteppedOnto += 1
			}
			matrix[fromY][fromX] = 'X'
		}()
		switch matrix[posY][posX] {
		case '^':
			posY--

			if posY-1 < 0 {
				return
			}

			forwardTile = matrix[posY-1][posX]
			if forwardTile == '#' {
				matrix[posY][posX] = '>'
			} else {
				matrix[posY][posX] = '^'
			}
		case '>':
			posX++

			if len(matrix[posY])-1 < posX+1 {
				return
			}

			forwardTile = matrix[posY][posX+1]
			if forwardTile == '#' {
				matrix[posY][posX] = 'v'
			} else {
				matrix[posY][posX] = '>'
			}
		case 'v':
			posY++

			if len(matrix)-1 < posY+1 {
				return
			}

			forwardTile = matrix[posY+1][posX]
			if forwardTile == '#' {
				matrix[posY][posX] = '<'
			} else {
				matrix[posY][posX] = 'v'
			}
		case '<':
			posX--

			if posX-1 < 0 {
				return
			}

			forwardTile = matrix[posY][posX-1]
			if forwardTile == '#' {
				matrix[posY][posX] = '^'
			} else {
				matrix[posY][posX] = '<'
			}
		default:
			log.Fatalf("Unkown tile type '%v'", string(matrix[posY][posX]))
			return
		}
	}
}
