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

func countValid(passports []string) int {
	var res int
	for _, p := range passports {
		passport := make(map[string]string)
		fields := strings.Fields(p)
		for _, f := range fields {
			key, value := f[0:3], f[4:]
			passport[key] = value
		}
		if (len(passport) == 8) || (len(passport) == 7 && passport["cid"] == "") {
			res++
		}
	}
	return res
}

func main() {
	passports, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(countValid(passports))
}
