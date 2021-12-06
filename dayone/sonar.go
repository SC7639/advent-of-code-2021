package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"runtime"
	"strconv"
	"strings"
)

func partOne(input []string) int {
	var prevVal int
	var i int
	start := 1

	for j, v := range input {
		if j < start {
			continue
		}

		vint, err := strconv.Atoi(v)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		if vint > prevVal {
			i++
		}
		prevVal = vint
	}

	return i
}

func intWindow(input []string) ([]int, error) {
	var intArr []int

	for _, v := range input {
		vint, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		intArr = append(intArr, vint)
	}

	return intArr, nil
}

func windowSum(input []int) int {
	var sum int
	for _, v := range input {
		sum += v
	}
	return sum
}

func partTwo(input []string) int {
	var prevWindow int
	var i int

	for j := 0; j < len(input); j++ {
		if j+3 > len(input) {
			break
		}

		intArr, err := intWindow(input[j : j+3])
		if err != nil {
			log.Println(err.Error())
			continue
		}

		window := windowSum(intArr)
		if window > prevWindow && prevWindow > 0 {
			i++
		}

		prevWindow = window
	}

	return i
}

func main() {
	var __dir string

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Cant' get stack information")
	}
	__dir = path.Dir(filename)

	input, err := ioutil.ReadFile(path.Join(__dir, "./sonar.txt"))
	if err != nil {
		panic(err)
	}

	inputArr := strings.Split(string(input), "\n")
	one := partOne(inputArr)
	two := partTwo(inputArr)

	fmt.Printf("There are %d measurements that are larger than previous measurement\n", one)
	fmt.Printf("There are %d window measurements that are lareger than previous windows\n", two)
}
