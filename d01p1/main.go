package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const filename = "input.txt"

func readInput() (map[int]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	nums := make(map[int]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		nums[n]++
	}
	return nums, nil
}

func findMultiple(nums map[int]int) int {
	for n, o := range nums {
		c := 2020 - n
		// special case if number exists twice in the input
		if c == n {
			if o > 1 {
				return c * n
			}
			continue
		}

		if nums[c] > 0 {
			return n * c
		}
	}
	return -1
}

func main() {
	nums, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(findMultiple(nums))
}
