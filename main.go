package main

import (
	"fmt"
	"strings"
)

// package-level vars must be declared with "var" and "=" sign
var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

// var bookings [50]string
var bookings []string

func main() {
	//all vars (or package) must be used in GO
	fmt.Printf("Type : %T\n", conferenceName)

	greetUsers()
	//fmt.Printf("Welcome to %v booking APP!\n", conferenceName)

	for {
		firstName, lastName, email, userTickets := getUserInput()
		// in GO you can return multiple values from func(), in that case = 3
		isValidTicketsNumber, isValidName, isValidEmail := validateUserInput(userTickets, firstName, lastName, email)

		// you can use ":=" instead of "=" if you don't want to write "var"
		if isValidTicketsNumber && isValidName && isValidEmail {

			bookTicket(userTickets, firstName, lastName, email)
			// function to print first names from "booking" slice
			//PrintFirstNames(bookings)
			firstNames := getFirstNames(bookings)

			fmt.Printf("The first names of  bookings are: %v\n", firstNames)

			var noTicketsRemainning bool = remainingTickets == 0
			if noTicketsRemainning {
				// end program
				fmt.Println("Our conference is booked out. Come back next year!!")
				break
			}
		} else {
			if !isValidEmail {
				fmt.Println("you entered invalid email, please check the @ sign in your address")
			}
			if !isValidName {
				fmt.Println("first or last name is too short, try write it again")
			}
			if !isValidTicketsNumber {
				fmt.Println("you entered invalid number of tickets, pay more attention.. ")
			}
			//fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets)
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking app", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// bookings []string = internal parameters but the second "[]string" is output parameters
func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
		// fmt.Println(index)   "_" means that we ignore "index"
	}

	//fmt.Printf("The first names of  bookings are: %v\n", firstNames)
	return firstNames
}

func validateUserInput(userTickets uint, firstName string, lastName string, email string) (bool, bool, bool) {
	var isValidTicketsNumber = userTickets <= remainingTickets && remainingTickets > 0
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")

	return isValidTicketsNumber, isValidName, isValidEmail

}

func getUserInput() (string, string, string, uint) {
	var firstName string

	var lastName string

	var email string
	var userTickets uint
	// ask user for their name
	//	fmt.Scan(&userName)
	/*fmt.Println(remainingTickets)
	fmt.Println(&remainingTickets)
	fmt.Printf("%T\n", remainingTickets)
	fmt.Printf("%T\n", &remainingTickets)
	*/
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email name: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets pls: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	/*fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The 1st element of array: %v\n", bookings[0])
	fmt.Printf("The type array: %T\n", bookings)
	fmt.Printf("The whole array: %v\n", len(bookings))
	*/
	fmt.Printf("Thank you %v %v for booking %v tickets. You'll receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
