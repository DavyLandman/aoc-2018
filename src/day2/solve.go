package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readLines(file string) []string {
	var result []string
	input, err := os.Open(file)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	} else {
		defer input.Close()
	}
	lines := bufio.NewScanner(input)
	for lines.Scan() {
		result = append(result, strings.TrimSpace(lines.Text()))
	}
	return result
}

func count(freqs []int, search int) int {
	for _, f := range freqs {
		if f == search {
			return 1
		}
	}
	return 0
}

func solvePart1(ids []string) {

	var twos, threes int
	for _, id := range ids {

		var chars [256]int
		for _, c := range id {
			chars[c]++
		}
		twos += count(chars[:], 2)
		threes += count(chars[:], 3)
	}

	fmt.Printf("Checksum: %d\n", twos*threes)
}

func offByOne(a string, b string) (bool, string) {
	difference := -1

	for ix, ca := range a {
		cb := []rune(b)[ix]
		if cb != ca {
			if difference == -1 {
				difference = ix
			} else {
				return false, ""
			}
		}
	}

	if difference == -1 {
		fmt.Printf("Unexpected equal strings: %s and %s\n", a, b)
		os.Exit(2)
	}
	return true, a[:difference] + a[difference+1:]
}

func solvePart2(ids []string) {
	for ix, a := range ids {
		for _, b := range ids[ix+1:] {
			match, common := offByOne(a, b)
			if match {
				fmt.Printf("Common letters with id: %s\n", common)
				return
			}
		}
	}
	fmt.Println("Could not find common")
}

func main() {
	ids := readLines(os.Args[1])
	solvePart1(ids)
	solvePart2(ids)
}
