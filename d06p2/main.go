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
	return strings.Split(strings.TrimSpace(string(data)), "\n\n"), nil
}

func countAllAnswers(answers []string) int {
	var count int
	for _, a := range answers {
		count += countValidAnswers(a)
	}
	return count
}

func countValidAnswers(a string) int {
	m := make(map[rune]int)
	for _, c := range a {
		if c != '\n' {
			m[c]++
		}
	}

	numGroups := len(strings.Split(a, "\n"))
	var res int
	for _, count := range m {
		if count == numGroups {
			res++
		}
	}
	return res
}

func main() {
	answers, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(countAllAnswers(answers))
}
