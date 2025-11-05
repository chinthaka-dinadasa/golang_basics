package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := captureUserInput()

		remainingTickets = remainingTickets - userTickets

		isValidName, isValidEmail, isValidUserTicketCount := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTicketCount {

			processBooking(firstName, lastName, email, userTickets)

			firstNames := printFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				//end the program
				fmt.Println("Conference is fully booked. Please try next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("Firstname and lastname you entered is too short\n")
			}
			if !isValidEmail {
				fmt.Printf("Email entered is not correct\n")
			}
			if !isValidUserTicketCount {
				fmt.Printf("You have entered wrong number of tickets\n")
			}

		}

	}

}

func greetUsers() {
	fmt.Printf("Welcom to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "and", remainingTickets, "are still available")
	fmt.Println("You can buy tickets from here.")
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func captureUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email")
	fmt.Scan(&email)
	fmt.Println("Number of tickets ? ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func processBooking(firstName string, lastName string, email string, userTickets uint) {
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %v %v for your reservation for %v tickets. You will get confirmation to %v soon\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("The whole slice %v\n", bookings)
	fmt.Printf("The first value %v\n", bookings[0])
	fmt.Printf("The length %v\n", len(bookings))
	fmt.Printf("The Type %T\n", bookings)
	fmt.Printf("There are all our bookings in the application %v\n", bookings)
}
