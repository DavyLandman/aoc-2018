package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

type square struct {
	id                       int
	left, right, top, bottom int
}

func (s square) place(canvas [][][]int) {
	for x := s.left; x <= s.right; x++ {
		for y := s.top; y <= s.bottom; y++ {
			canvas[x][y] = append(canvas[x][y], s.id)
		}
	}
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}

func parseSquares(lines []string) []square {
	var result []square
	matcher := regexp.MustCompile("[0-9]+")
	for _, l := range lines {
		numbers := matcher.FindAllString(l, -1)
		result = append(result, square{
			id:     toInt(numbers[0]),
			left:   toInt(numbers[1]),
			top:    toInt(numbers[2]),
			right:  toInt(numbers[1]) + toInt(numbers[3]) - 1,
			bottom: toInt(numbers[2]) + toInt(numbers[4]) - 1})
	}
	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func buildCanvas(squares []square) [][][]int {
	maxWidth := 0
	maxHeight := 0
	for _, sq := range squares {
		maxWidth = max(maxWidth, sq.right)
		maxHeight = max(maxHeight, sq.bottom)
	}

	result := make([][][]int, maxWidth+1)
	for ix := range result {
		result[ix] = make([][]int, maxHeight+1)
	}

	for _, sq := range squares {
		sq.place(result)
	}

	return result
}

func solve1(squares []square) {
	canvas := buildCanvas(squares)

	doubleClaim := 0
	for _, row := range canvas {
		for _, cell := range row {
			if len(cell) > 1 {
				doubleClaim++
			}
		}
	}

	fmt.Printf("Two or more claims: %d\n", doubleClaim)
}

func solve2(squares []square) {
	canvas := buildCanvas(squares)

	singleClaim := make(map[int]bool)
	for _, s := range squares {
		singleClaim[s.id] = true
	}

	for _, row := range canvas {
		for _, cell := range row {
			if len(cell) > 1 {
				for _, id := range cell {
					delete(singleClaim, id)
				}
			}
		}
	}

	for c := range singleClaim {
		fmt.Printf("Single claimed piece: %d\n", c)
	}

}

func main() {
	squares := parseSquares(readLines(os.Args[1]))
	solve1(squares)
	solve2(squares)
}
