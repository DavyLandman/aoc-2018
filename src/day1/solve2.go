package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var deltas []int

	input, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	reader := bufio.NewScanner(input)
	for reader.Scan() {
		i, err := strconv.Atoi(reader.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		deltas = append(deltas, i)
	}

	var freq int
	var seen = make(map[int]bool)
	seen[0] = true
	for {
		for _, d := range deltas {
			freq += d
			_, alreadySeen := seen[freq]
			if alreadySeen {
				fmt.Printf("Double freq: %d", freq)
				os.Exit(0)
			}
			seen[freq] = true
		}
	}
}
