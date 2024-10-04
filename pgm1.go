/*
	Author: Colby McClure
	Course: CS-424-01
	Assignment: Programming Assignment 1
	System Info: go version go1.23.1 windows/amd64

	Description: 	This program reads an input file containing baseball player statistics and calculates the batting average, 
					slugging percentage, and on base percentage for each player in the input file. The program then 
					prints a report of the player statistics and any errors encountered while reading and parsing the file.
*/

package main

// Import statements needed 
import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"sort"
)

// An Error struct to store error messages in the form of strings
type Error struct {
	message string
}

// A Player struct used to store player statistics, such as what should be read from the input file
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

// The BattingAverage method calculates the batting average of a player
func (p Player) BattingAverage() float64 {
	return float64(p.singles + p.doubles + p.triples + p.homeRuns) / float64(p.atBats)
} 

// The SluggingPercentage method calculates the slugging percentage of a player
func (p Player) SluggingPercentage() float64 {
	return float64(p.singles + (2 * p.doubles) + (3 * p.triples) + (4 * p.homeRuns)) / float64(p.atBats)
}

// The OnBasePercentage method calculates the on base percentage of a player
func (p Player) OnBasePercentage() float64 {
	return float64(p.singles + p.doubles + p.triples + p.homeRuns + p.walks + p.hitByPitch) / float64(p.plateAppearances)
}

// The ReadFile function reads the input and returns a slice of Player and Error structs
func ReadFile(fileName string) ([]Player, []Error) {

	// Initialize the player and error slices
	var players []Player
	var errors []Error
	
	// Open the file
	file, err := os.Open(fileName)

	// If there is an error opening the file, print the error and return nil for both slices
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil, nil 
	}

	// Create a new scanner to read the file
	scanner := bufio.NewScanner(file)

	// Initialize the line number to 0
	// This gets used to keep track of the line number in the file in case of an error
	lineNumber := 0

	// For each line in the file
	for scanner.Scan() {

		// Increment the line number
		lineNumber++

		// Read the line and split it into fields
		line := scanner.Text()
		fields := strings.Fields(line)

		// Set the first field to the first name, the second field to the last name
		firstName := fields[0]
		lastName := fields[1]
		
		// If the line does not contain sufficient data, add the error to the errors slice and process the next line
		if len(fields) != 10 {
			err := fmt.Sprintf("line %d: %s; Error: Line does not contain enough data", lineNumber, lastName)
			errors = append(errors, Error{err})
			continue
		}

		// Set the third field to the plate appearances
		plateAppearances, err := strconv.Atoi(fields[2])

		// If there is an error converting the plate appearances to an integer, print an error message and process the next line
		if err != nil {
			err := fmt.Sprintf("line %d: %s; Error: Invalid plate appearances", lineNumber, lastName)
			errors = append(errors, Error{err})
			continue
		}
		
		// Set the fourth field to the at bats
		atBats, err := strconv.Atoi(fields[3])

		// If there is an error converting the at bats to an integer, print an error message and process the next line
		if err != nil {
			err := fmt.Sprintf("line %d: %s; Error: Invalid at bats", lineNumber, lastName)
			errors = append(errors, Error{err})
			continue
		}
		
		// Set the fifth field to the singles
		singles, err := strconv.Atoi(fields[4])

		// If there is an error converting the singles to an integer, print an error message and process the next line
		if err != nil {
			err := fmt.Sprintf("line %d: %s; Error: Invalid singles", lineNumber, lastName)
			errors = append(errors, Error{err})
			continue
		}
		
		// Set the sixth field to the doubles
		doubles, err := strconv.Atoi(fields[5])
		
		// If there is an error converting the doubles to an integer, print an error message and process the next line
		if err != nil {
			err := fmt.Sprintf("line %d: %s; Error: Invalid doubles", lineNumber, lastName)
			errors = append(errors, Error{err})
			continue
		}

		// Set the seventh field to the triples
		triples, err := strconv.Atoi(fields[6])

		// If there is an error converting the triples to an integer, print an error message and process the next line
		if err != nil {
			err := fmt.Sprintf("line %d: %s; Error: Invalid triples", lineNumber, lastName)
			errors = append(errors, Error{err})
			continue
		}

		// Set the eighth field to the home runs
		homeRuns, err := strconv.Atoi(fields[7])

		// If there is an error converting the home runs to an integer, print an error message and process the next line
		if err != nil {
			err := fmt.Sprintf("line %d: %s; Error: Invalid home runs", lineNumber, lastName)
			errors = append(errors, Error{err})
			continue
		}

		// Set the ninth field to the walks
		walks, err := strconv.Atoi(fields[8])

		// If there is an error converting the walks to an integer, print an error message and process the next line
		if err != nil {
			err := fmt.Sprintf("line %d: %s; Error: Invalid walks", lineNumber, lastName)
			errors = append(errors, Error{err})
			continue
		}

		// Set the tenth field to the hit by pitch
		hitByPitch, err := strconv.Atoi(fields[9])

		// If there is an error converting the hit by pitch to an integer, print an error message and process the next line
		if err != nil {
			err := fmt.Sprintf("line %d: %s; Error: Invalid hit by pitch", lineNumber, lastName)
			errors = append(errors, Error{err})
			continue
		}

		// Append the Player struct to the players slice
		players = append(players, Player{firstName, lastName, plateAppearances, atBats, singles, doubles, triples, homeRuns, walks, hitByPitch})
	}

	// If there is an error reading the file, print the error and return nil for both slices
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
		return nil, nil 
	} 

	// File should be closed after reading
	defer file.Close()

	// When finished reading the file, return both the players and errors slices
	return players, errors
}

// A BySlugging type to sort the players by slugging percentage
type BySlugging []Player

// Implement the Len, Swap, and Less methods for the BySlugging type which are required for sorting
func (a BySlugging) Len() int { 
	return len(a) 
}

func (a BySlugging) Swap(i, j int) {
	a[i], a[j] = a[j], a[i] 
}

func (a BySlugging) Less(i, j int) bool {
	return a[i].SluggingPercentage() > a[j].SluggingPercentage() 
}

func main() {

	// Initialize the fileName variable
	var fileName string

	// Print the welcome message and prompt the user to enter the input file
	fmt.Println("Welcome to the player statistics program!\nPlease enter the input file that contains the player's statistics.")
	fmt.Println("Enter the name of the input file: ")

	// Read the input file name
	fmt.Scanln(&fileName) 

	// Call the ReadFile function and store the players and errors slices
	players, errors := ReadFile(fileName)

	// If there were any issues parsing the file, print an error message and return
	if players == nil {
		fmt.Println("Critical error, please try again")
		return
	}

	// Initialize the number of players used for the report
	numberPlayers := len(players)

	// Print the beginning of the player statistics report
	fmt.Printf("\n\nBaseball player stats report --- %d players processed\n", numberPlayers)
	fmt.Println("Player Name     Average    Slugging      On Base%")
	fmt.Println("-------------------------------------------------")

	// Sort the players by slugging percentage
	sort.Sort(BySlugging(players))

	// For each player, print the player's last name, first name, batting average, slugging percentage, and on base percentage
	for _, player := range players {
		fmt.Printf("%s, %-5s %10.3f %10.3f %10.3f\n", player.lastName, player.firstName, player.BattingAverage(), player.SluggingPercentage(), player.OnBasePercentage())
	}

	// Print the errors section
	fmt.Println("\n\nErrors encountered:")

	// For each error, print the error message
	for _, err := range errors {
		fmt.Println(err.message)
	}
}