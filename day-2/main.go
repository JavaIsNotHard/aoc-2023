package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id   int
	sets []Set
}

type Set struct {
	red   int
	blue  int
	green int
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

func createGames(lines []string) ([]Game, error) {
	games := make([]Game, 0)

	for _, line := range lines {
		game, err := convertLineToGame(line)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}

	return games, nil
}

func determineColorVal(sets []string) (int, int, int, error) {
	red, blue, green := 0, 0, 0
	for _, part := range sets {
		splittedPart := strings.Split(part, " ")
		val, err := strconv.Atoi(splittedPart[0])
		if err != nil {
			return 0, 0, 0, err
		}
		switch splittedPart[1] {
		case "red":
			red = val
		case "blue":
			blue = val
		case "green":
			green = val
		}
	}
	return red, blue, green, nil
}

func convertLineToGame(line string) (Game, error) {
	colon := strings.Index(line, ":")
	firstSpace := strings.Index(line, " ")
	id := line[firstSpace+1 : colon]

	sets := make([]Set, 0)
	for _, rawSets := range strings.Split(line[colon+1:], ";") {
		colorGroups := strings.Split(strings.TrimSpace(rawSets), ", ")
		red, blue, green, err := determineColorVal(colorGroups)
		if err != nil {
			return Game{}, err
		}
		sets = append(sets, Set{
			red:   red,
			blue:  blue,
			green: green,
		})
	}

	intId, _ := strconv.Atoi(id)

	return Game{
		id:   intId,
		sets: sets,
	}, nil
}

func partOne() {
	filename := "input.txt"
	lines, err := readFile(filename)

	if err != nil {
		log.Print("Couldn't read file")
		os.Exit(1)
	}

	games, err := createGames(lines)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	var sumId = 0
	for _, game := range games {
		validSets := 0
		for _, set := range game.sets {
			if set.red <= 12 && set.blue <= 14 && set.green <= 13 {
				validSets++
				fmt.Println(set)
			}
		}
		if validSets == len(game.sets) {
			sumId += game.id
		}
	}

	fmt.Println(sumId)
}

func main() {
	partOne()
}
