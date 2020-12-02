package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const filename = "input.txt"

func ReadInput() ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	var passwords []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		passwords = append(passwords, scanner.Text())
	}
	return passwords, nil
}

func countValid(passwords []string) int {
	var res int
	for _, p := range passwords {
		split1 := strings.Split(p, ": ")
		split2, password := split1[0], split1[1]
		split3 := strings.Split(split2, " ")
		split4, char := split3[0], split3[1]
		limits := strings.Split(split4, "-")
		l1, _ := strconv.Atoi(limits[0])
		l2, _ := strconv.Atoi(limits[1])
		if validPassword(password, char, l1, l2) {
			res++
		}
	}
	return res
}

func validPassword(password, char string, l1, l2 int) bool {
	count := 0
	for _, c := range password {
		if string(c) == char {
			count++
		}
	}
	return (count >= l1 && count <= l2)
}

func main() {
	passwords, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(countValid(passwords))
}
