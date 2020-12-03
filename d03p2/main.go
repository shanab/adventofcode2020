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

func findMultiple(area [][]bool, slopes [][]int) int {
	res := 1
	for _, s := range slopes {
		res *= countTrees(area, s[0], s[1])
	}
	return res
}

func main() {
	area, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	fmt.Println(findMultiple(area, slopes))
}
