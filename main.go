package main

import (
	"fmt"
	"strings"
)

func main() {
	//all vars (or package) must be used in GO
	conferenceName := "Go Conference"
	fmt.Printf("Type : %T\n", conferenceName)
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	//var bookings [50]string
	var bookings []string
	fmt.Printf("Welcome to %v booking APP!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
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

		var isValidTicketsNumber = userTickets <= remainingTickets && remainingTickets > 0
		var isValidName = len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")

		// you can use ":=" instead of "=" if you don't want to write "var"
		if isValidTicketsNumber && isValidName && isValidEmail {
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

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
				// fmt.Println(index)   "_" means that we ignore "index"
			}

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
