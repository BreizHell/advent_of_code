package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func every[T any](ts []T, fn func(int, T) bool) bool {
	for i, t := range ts[:len(ts) -1] {
		if !fn(i, t) {
			return false
		}
	}
	return true
}


func checkReportValidity(report *[]int8) bool {
	return every(*report, func (i int, left int8) bool {
		right := (*report)[i+1]
		delta := left - right
		return 1 <= delta && delta <= 3
	}) || every(*report, func (i int, left int8) bool {
		right := (*report)[i+1]
		delta := right - left
		return 1 <= delta && delta <= 3
	})
}

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Populate the report matrix
	reports := make([][]int8, 0, 1000)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reportAsStrings := strings.Split(scanner.Text(), " ")
		reportLine := make([]int8, 0, len(reportAsStrings))

		for _, valStr := range(reportAsStrings) {
			val, _ := strconv.Atoi(valStr)
			reportLine = append(reportLine, int8(val))
		}
		reports = append(reports, reportLine)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("reading standard input: %q", err)
	}

	// count valid reports
	var satisfactionsCount int32 = 0

	for _, report := range(reports) {
		isValid := checkReportValidity(&report)
		
		if !isValid {
			substractedSlice := make([]int8, len(report)-1)
			for i := range(report) {
				copy(substractedSlice[:], report[:i])
				copy(substractedSlice[i:], report[i+1:])
				if checkReportValidity(&substractedSlice) {
					isValid = true
					break
				}
			}
		}

		if isValid {
			satisfactionsCount += 1
		}
	}

	println(satisfactionsCount)
}
