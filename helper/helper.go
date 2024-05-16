package helper

import "strings"

var MyVar = "somevalue"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= uint(remainingTickets)
	return isValidName, isValidEmail, isValidTicketNumber
}