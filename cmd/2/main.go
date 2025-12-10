package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func solve(isInvalid func(int) bool) {
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var ranges []string = strings.Split(scanner.Text(), ",")
	var n int = len(ranges)
	var L []int = make([]int, n)
	var R []int = make([]int, n)
	for i := range n {
		parts := strings.Split(ranges[i], "-")
		L[i], err = strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		R[i], err = strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	result := 0
	for i := range n {
		for x := L[i]; x <= R[i]; x++ {
			if isInvalid(x) {
				result += x
			}
		}
	}
	fmt.Printf("Result: %d\n", result)
}

func main() {
	solve(isInvalidPartone)
	solve(isInvalidPartTwo)
}

func isInvalidPartone(x int) bool {
	s := strconv.Itoa(x)
	if len(s)%2 != 0 {
		return false
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)/2+i] {
			return false
		}
	}
	return true
}

func isInvalidPartTwo(x int) bool {
	s := strconv.Itoa(x)
	for n := 1; n < len(s); n++ {
		if len(s)%n != 0 {
			continue
		}
		pattern := s[0:n]
		isInvalid := true
		for i := 0; i < len(s); i += n {
			if s[i:i+n] != pattern {
				isInvalid = false
				break
			}
		}
		if isInvalid {
			return true
		}
	}
	return false
}
