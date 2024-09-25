package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type Error struct {
	message string
}

type Player struct {
	firstName string
	lastName  string
	plateAppearances int
	atBats int
	singles int
	doubles int
	triples int
	homeRuns int
	walks int
	hitByPitch int
}

func (p Player) BattingAverage() float64 {
	return float64(p.singles + p.doubles + p.triples + p.homeRuns) / float64(p.atBats)
} 

func (p Player) SluggingPercentage() float64 {
	return float64(p.singles + (2 * p.doubles) + (3 * p.triples) + (4 * p.homeRuns)) / float64(p.atBats)
}

func ReadFile(fileName string) ([]Player, []Error) {

	var players []Player
	var errors []Error
	
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil, nil 
	}

	scanner := bufio.NewScanner(file)

	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		fields := strings.Fields(line)


		firstName := fields[0]
		lastName := fields[1]
		plateAppearances, err := strconv.Atoi(fields[2])

		if err != nil {
			fmt.Println("Invalid plate appearances: ", fields[2])
			continue
		}

		if len(fields) != 10 {
			err := fmt.Sprintf("line %d: %s; Error: Line does not contain enough data", lineNumber, lastName)
			errors = append(errors, Error{err})
			continue
		}
		
		atBats, err := strconv.Atoi(fields[3])

		if err != nil {
			err := fmt.Sprintf("line %d: %s; Error: Invalid at bats", lineNumber, lastName)
			errors = append(errors, Error{err})
			continue
		}

		singles, err := strconv.Atoi(fields[4])

		if err != nil {
			err := fmt.Sprintf("line %d: %s; Error: Invalid singles", lineNumber, lastName)
			errors = append(errors, Error{err})
			continue
		}

		doubles, err := strconv.Atoi(fields[5])

		if err != nil {
			err := fmt.Sprintf("line %d: %s; Error: Invalid doubles", lineNumber, lastName)
			errors = append(errors, Error{err})
			continue
		}

		triples, err := strconv.Atoi(fields[6])

		if err != nil {
			err := fmt.Sprintf("line %d: %s; Error: Invalid triples", lineNumber, lastName)
			errors = append(errors, Error{err})
			continue
		}

		homeRuns, err := strconv.Atoi(fields[7])

		if err != nil {
			err := fmt.Sprintf("line %d: %s; Error: Invalid home runs", lineNumber, lastName)
			errors = append(errors, Error{err})
			continue
		}

		walks, err := strconv.Atoi(fields[8])

		if err != nil {
			err := fmt.Sprintf("line %d: %s; Error: Invalid walks", lineNumber, lastName)
			errors = append(errors, Error{err})
			continue
		}

		hitByPitch, err := strconv.Atoi(fields[9])

		if err != nil {
			err := fmt.Sprintf("line %d: %s; Error: Invalid hit by pitch", lineNumber, lastName)
			errors = append(errors, Error{err})
			continue
		}

		players = append(players, Player{firstName, lastName, plateAppearances, atBats, singles, doubles, triples, homeRuns, walks, hitByPitch})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
		return nil, nil 
	} 

	defer file.Close()

	return players, errors
}

func main() {

	var fileName string

	fmt.Println("Welcome to the player statistics program!\nPlease enter the input file that contains the player's statistics.")
	fmt.Println("Enter the name of the input file: ")
	fmt.Scanln(&fileName) 

	players, errors := ReadFile(fileName)

	for _, player := range players {
		fmt.Printf("%s %s %.3f %.3f\n", player.firstName, player.lastName, player.BattingAverage(), player.SluggingPercentage())
	}

	for _, err := range errors {
		fmt.Println(err.message)
	}



}