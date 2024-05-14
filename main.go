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

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

for  {
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

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")

	if userTickets < uint(remainingTickets) {
		remainingTickets = remainingTickets - int(userTickets)
		bookings = append(bookings, firstName + " " + lastName)
	
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n",remainingTickets, conferenceName)
	
		firstNames := []string{}
		for  _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
	
		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
	}
} else if userTickets == uint(remainingTickets) {

} else {
	fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets", remainingTickets, userTickets)
   }
  }
}