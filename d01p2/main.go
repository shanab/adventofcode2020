package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const filename = "input.txt"

func readInput() (map[int]bool, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	nums := make(map[int]bool)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		nums[n] = true
	}
	return nums, nil
}

func findMultiple(nums map[int]bool) int {
	for n := range nums {
		for m := range nums {
			c := 2020 - (n + m)
			if c > 0 && nums[c] {
				return n * m * c
			}
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
