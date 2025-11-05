package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcom to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "and", remainingTickets, "are still available")
	fmt.Println("You can buy tickets from here.")

	var bookings = []string{}
	for {

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

		if userTickets <= remainingTickets {
			fmt.Printf("We only have %v tickets, Please try again\n", remainingTickets)
			continue
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The whole slice %v\n", bookings)
		fmt.Printf("The first value %v\n", bookings[0])
		fmt.Printf("The length %v\n", len(bookings))
		fmt.Printf("The Type %T\n", bookings)

		fmt.Printf("Thank you %v %v for your reservation for %v tickets. You will get confirmation to %v soon\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		fmt.Printf("There are all our bookings in the application %v\n", bookings)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
		if remainingTickets == 0 {
			//end the program
			fmt.Println("Conference is fully booked. Please try next year.")
			break
		}

	}

}
