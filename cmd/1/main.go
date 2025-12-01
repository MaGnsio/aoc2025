package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne() {
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var password, position int = 0, 50
	for scanner.Scan() {
		var c string = scanner.Text()
		m, err := strconv.Atoi(c[1:])
		if err != nil {
			log.Fatal(err)
		}
		if strings.HasPrefix(c, "R") {
			position = (position + m) % 100
		} else {
			position = ((position-m)%100 + 100) % 100
		}
		if position == 0 {
			password += 1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("password is %d\n", password)
}

func partTwo() {
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var password, position int = 0, 50
	for scanner.Scan() {
		var c string = scanner.Text()
		m, err := strconv.Atoi(c[1:])
		if err != nil {
			log.Fatal(err)
		}
		if strings.HasPrefix(c, "R") {
			var needToZero int = (100 - position) % 100
			if needToZero <= m {
				password += (m-needToZero)/100 + 1
			}
			if needToZero == 0 {
				password -= 1
			}
			position = (position + m) % 100
		} else {
			var needToZero int = position
			if needToZero <= m {
				password += (m-needToZero)/100 + 1
			}
			if needToZero == 0 {
				password -= 1
			}
			position = ((position-m)%100 + 100) % 100
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("password is %d\n", password)
}

func main() {
	partOne()
	partTwo()
}
