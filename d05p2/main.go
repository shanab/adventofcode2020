package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

const filename = "input.txt"

func ReadInput() ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var passes []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		passes = append(passes, scanner.Text())
	}
	return passes, nil
}

func seatID(pass string) int {
	row := seatPosition([]byte(pass[:7]), 0, 127)
	column := seatPosition([]byte(pass[7:]), 0, 7)
	return row*8 + column
}

func seatPosition(pass []byte, bottom, top int) int {
	if top == bottom {
		return top
	}
	if pass[0] == 'F' || pass[0] == 'L' {
		return seatPosition(pass[1:], bottom, (top-bottom)/2+bottom)
	}
	if pass[0] == 'B' || pass[0] == 'R' {
		return seatPosition(pass[1:], (top+1-bottom)/2+bottom, top)
	}
	return -1
}

func findMissingSeat(passes []string) int {
	var seats []int
	for _, p := range passes {
		seats = append(seats, seatID(p))
	}
	sort.Ints(seats)
	for i := 0; i < len(seats)-1; i++ {
		if seats[i+1] != seats[i]+1 {
			return seats[i] + 1
		}
	}
	return -1
}

func main() {
	passes, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(findMissingSeat(passes))
}
