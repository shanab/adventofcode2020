package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const filename = "input.txt"

func ReadInput() ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n\n"), nil
}

func countAllAnswers(answers []string) int {
	var count int
	for _, a := range answers {
		count += countUniqueAnswers(a)
	}
	return count
}

func countUniqueAnswers(a string) int {
	m := make(map[rune]int)
	for _, c := range a {
		if c != '\n' {
			m[c]++
		}
	}
	return len(m)
}

func main() {
	answers, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(countAllAnswers(answers))
}
