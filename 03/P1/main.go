package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type operation struct {
	operand  string
	leftVal  int32
	rightVal int32
}

func main() {
	parser := regexp.MustCompile(`(mul)\(([0-9]{1,3}),([0-9]{1,3})\)`)

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Populate the operation stack
	operations := make([]operation, 0, 1000)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := parser.FindAllStringSubmatch(scanner.Text(), -1)

		for _, opMatch := range matches {
			leftVal, paramParseErr := strconv.Atoi(opMatch[2])

			if paramParseErr != nil {
				log.Fatal(paramParseErr)
			}

			rightVal, paramParseErr := strconv.Atoi(opMatch[3])

			if paramParseErr != nil {
				log.Fatal(paramParseErr)
			}
			operations = append(operations, operation{opMatch[1], int32(leftVal), int32(rightVal)})
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("reading standard input: %q", err)
	}

	// compute sum
	var sum int64 = 0
	for _, op := range operations {
		switch op.operand {
		case "mul":
			sum += int64(op.leftVal * op.rightVal)
		default:
			log.Fatalf("Unkown operand in stack of operations")
		}
	}

	println(sum)
}
