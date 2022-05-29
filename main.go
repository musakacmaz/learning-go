package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName = "Golang Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTickets := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTickets {
			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry, we are sold out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Invalid name")
			}
			if !isValidEmail {
				fmt.Println("Invalid email")
			}
			if !isValidTickets {
				fmt.Println("Invalid number of tickets")
			}
		}

	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTickets
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for first name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	// ask user for last name
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	// ask user for email
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	// ask user for number of tickets
	fmt.Println("How many tickets do you want to buy: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("We have %v tickets left for %v \n", remainingTickets, conferenceName)
}
