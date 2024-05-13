package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50
	bookings := []string{}

	//fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceTickets)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

for{
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	
	fmt.Println("Enter your  email address:")
	fmt.Scan(&email)
	
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)
	
	remainingTickets = remainingTickets - int(userTickets)
	bookings = append(bookings, firstName + " " + lastName)

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The whole value: %v\n", bookings[0])
	// fmt.Printf("Slice type: %T\n", bookings)
	// fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n",remainingTickets, conferenceName)

	firstName := []string{}
	for index, booking := range bookings{
		strings.Fields(booking)
	}
	fmt.Printf("These are all our bookings: %v\n", bookings)
  }
}