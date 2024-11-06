package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var prefix_hash_map = map[string]int {
    "zero": 0,
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
}



func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func findCalibrationValue(input []string) ([]int, error) {
	var firstDigit, lastDigit int
	var result []int

	for _, value := range input {

		for i := 0; i < len(value); i++ {
			if strings.ContainsAny(string(value[i]), "0123456789") {
				integer, err := strconv.Atoi(string(value[i]))
				if err != nil {
					return nil, err
				}
				firstDigit = integer
				break
			}
		}

		for i := len(value) - 1; i >= 0; i-- {
			if strings.ContainsAny(string(value[i]), "0123456789") {
				integer, err := strconv.Atoi(string(value[i]))
				if err != nil {
					return nil, err
				}
				lastDigit = integer
				break
			}
		}

		var tempResult = firstDigit*10 + lastDigit
		result = append(result, tempResult)
	}

	return result, nil
}

func findCalibrationValuePartTwo(input []string) ([]int, error) {
    var firstDigit, lastDigit int
    var result []int

    for _, value := range input {
        firstDigit = findFirstDigit(value)
        lastDigit = findLastDigit(value)
        
        var tempResult = firstDigit*10 + lastDigit
        result = append(result, tempResult)
    }

    return result, nil
}

func findFirstDigit(input string) (int) {
    for i := 0; i <= len(input); i++ {
        if found, digit := containsDigit(input[:i]); found {
            return digit
        } else if strings.ContainsAny(string(input[i]), "0123456789") {
			integer, _ := strconv.Atoi(string(input[i]))
            return integer
        }
    }
    panic("no digit found in " + input)
}

func findLastDigit(input string) (int) {
    for i := len(input) - 1; i >= 0; i-- {
        if found, digit := containsDigit(input[i:]); found {
            return digit
        } else if strings.ContainsAny(string(input[i]), "0123456789") {
			integer, _ := strconv.Atoi(string(input[i]))
            return integer
        }
    }
    panic("no digit found in " + input)
}



func containsDigit(input string) (bool, int) {
    for key, value := range prefix_hash_map {
        if strings.Contains(input, key) {
            return true,value
        }   
    }
    return false, 0
}

func hasNumberAsSuffix(input string) (bool, int) {
    for i := 0; i <= len(input); i++ {
        for key, value := range prefix_hash_map {
            if strings.HasSuffix(input, key) {
                return true, value
            }   
        }
        input = input[:len(input)-1]
    }
    return false, -1
}

func sumInt(input []int) int {
	var result int
	for _, integer := range input {
		result += integer
	}
	return result
}

func main() {
    // partOne()
    partTwo()
}

func partOne() {
    file, err := readFile("input.txt")
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	result, err := findCalibrationValue(file)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	sum := sumInt(result)
	fmt.Println(sum)

}

func partTwo() {
    file, err := readFile("parttwo.txt")
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	result, err := findCalibrationValuePartTwo(file)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	sum := sumInt(result)
	fmt.Println(sum)
}
