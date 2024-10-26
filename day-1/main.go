package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
                break;
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

        var tempResult = firstDigit * 10 + lastDigit
        result = append(result, tempResult)
    }

    return result, nil
}

func sumInt(input []int) (int) {
    var result int
    for _, integer := range input {
        result += integer
    }
    return result
}

func main() {
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
