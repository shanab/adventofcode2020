package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
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
		if valid(passport) {
			res++
		}
	}
	return res
}

func valid(p map[string]string) bool {
	byr, err := strconv.Atoi(p["byr"])
	if err != nil || byr < 1920 || byr > 2002 {
		return false
	}
	iyr, err := strconv.Atoi(p["iyr"])
	if err != nil || iyr < 2010 || iyr > 2020 {
		return false
	}
	eyr, err := strconv.Atoi(p["eyr"])
	if err != nil || eyr < 2020 || eyr > 2030 {
		return false
	}
	if !validHeight(p["hgt"]) {
		return false
	}
	if !validColor(p["hcl"]) {
		return false
	}
	if e := p["ecl"]; !(e == "amb" || e == "blu" || e == "brn" || e == "gry" || e == "grn" || e == "hzl" || e == "oth") {
		return false
	}
	if _, err := strconv.Atoi(p["pid"]); err != nil || len(p["pid"]) != 9 {
		return false
	}
	if err != nil {
		return false
	}
	return true
}

func validHeight(h string) bool {
	if len(h) < 4 {
		return false
	}
	value, unit := h[:len(h)-2], h[len(h)-2:]
	if unit != "cm" && unit != "in" {
		return false
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	if unit == "cm" && (v < 150 || v > 193) {
		return false
	}
	if unit == "in" && (v < 59 || v > 76) {
		return false
	}

	return true
}

func validColor(c string) bool {
	r := regexp.MustCompile(`^#(\d|[a-f]){6}$`)
	return r.MatchString(c)
}

func main() {
	passports, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(countValid(passports))
}
