package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Golang Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings []string

	for {
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

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("We have %v tickets left for %v \n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry, we are sold out")
				break
			}
		} else {
			fmt.Printf("Sorry, we only have %v tickets left, so you cannot book %v tickets.\n", remainingTickets, userTickets)
		}

	}
}
