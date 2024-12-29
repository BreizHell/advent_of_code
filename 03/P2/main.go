package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func If[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

func main() {
	parser := regexp.MustCompile(`do\(\)|don't\(\)|mul\((?P<left>[0-9]{1,3}),(?P<right>[0-9]{1,3})\)`)
	iL, iR := parser.SubexpIndex("left"), parser.SubexpIndex("right")

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sum int64 = 0
	scanner := bufio.NewScanner(file)
	enabled := true
	for scanner.Scan() {
		matches := parser.FindAllStringSubmatch(scanner.Text(), -1)
		for _, opMatch := range matches {
			switch opMatch[0] {
			case "do()":
				enabled = true
			case "don't()":
				enabled = false
			default:
				left, right := opMatch[iL], opMatch[iR]
				if enabled && left != "" && right != "" {
					fmt.Printf("%v,%v\n", left, right)
					leftVal, paramParseErr := strconv.Atoi(left)

					if paramParseErr != nil {
						log.Fatal(paramParseErr)
					}

					rightVal, paramParseErr := strconv.Atoi(right)

					if paramParseErr != nil {
						log.Fatal(paramParseErr)
					}
					sum += int64(leftVal * rightVal)
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("reading standard input: %q", err)
	}

	println(sum)
}
