package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const filename = "input.txt"

func ReadInput() ([][]bool, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var area [][]bool
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var row []bool
		line := scanner.Text()
		for _, c := range line {
			row = append(row, (c == '.'))
		}
		area = append(area, row)
	}
	return area, nil
}

func countTrees(area [][]bool, horizontal, vertical int) int {
	var h, v, count int
	for v < len(area)-1 {
		h = (h + horizontal) % len(area[0])
		v += vertical
		if !area[v][h] {
			count++
		}
	}
	return count
}

func main() {
	area, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	const (
		horizontal = 3
		vertical   = 1
	)
	fmt.Println(countTrees(area, horizontal, vertical))
}
