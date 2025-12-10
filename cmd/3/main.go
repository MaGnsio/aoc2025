package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func partOne() {
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		var c string = scanner.Text()
		n := len(c)
		var maxSuff []int = make([]int, n)
		maxSuff[n-1] = -1
		for i := n - 2; i >= 0; i-- {
			maxSuff[i] = max(maxSuff[i+1], int(c[i+1]-'0'))
		}
		best := 0
		for i := 0; i < n-1; i++ {
			best = max(best, 10*int(c[i]-'0')+maxSuff[i])
		}
		result += best
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result: %d\n", result)
}

func partTwo() {
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		var c string = scanner.Text()
		n, need, s := len(c), 12, ""
		var f [][]int = make([][]int, 10)
		for i := range n {
			f[int(c[i]-'0')] = append(f[int(c[i]-'0')], i)
		}
		for i := range n {
			f[int(c[i]-'0')] = f[int(c[i]-'0')][1:]
			if need == 0 {
				break
			}
			var canDoBetter bool = false
			for d := int(c[i]-'0') + 1; d <= 9; d++ {
				if len(f[d]) > 0 && n-f[d][0] >= need {
					canDoBetter = true
					break
				}
			}
			if canDoBetter {
				continue
			}
			s += string(c[i])
			need--
		}
		v, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("selected number: %s for %s\n", s, c)
		result += v
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result: %d\n", result)
}

func main() {
	partOne()
	partTwo()
}
