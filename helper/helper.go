package helper

import "strings"

func ValidateUserInput(userTickets uint, firstName string, lastName string, email string, remainingTickets uint) (bool, bool, bool) {
	var isValidTicketsNumber = userTickets <= remainingTickets && remainingTickets > 0
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")

	return isValidTicketsNumber, isValidName, isValidEmail

}
