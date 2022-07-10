package main

import (
	"fmt"
	"strings"
	"booking-app/helper"
)
  
var stationName string = "Oreva metro station"

const trainTickets int = 50

var remainingTickets int = 50
var bookings []string

func main() {

	greetUser()

	for { 

		//calls getUserInput function
		firstName, lastName, email, userTickets := getUserInput()

		isValidEmail, isValidNames, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets)

		//checks to make sure the user doesn't input a number > than tickets remaining
		if isValidTickets && isValidEmail && isValidNames {

			bookingLogic(userTickets, firstName, lastName, email)

			//call first name function
			firstNames := userFirstName()
			fmt.Printf("The first names of the bookings are:%v\n", firstNames)

			//checks if tickects are still available
			if remainingTickets == 0 {
				//end program
				fmt.Println("Train tickets are booked out.  ")
				break
			}
		} else {
			if !isValidEmail {
				fmt.Println("enter a valid email address")
			}
			if !isValidNames {
				fmt.Println("your name is to short")
			}
			if !isValidTickets {
				fmt.Printf("your input data is invalid")
			}

		}

	}

}

func greetUser() {
	fmt.Printf("welcome to %v ", stationName)
	fmt.Printf("we have a total of %v tickets but %v left.\n", trainTickets, remainingTickets)
	fmt.Println("Please book your tickets")

}
func userFirstName() []string {
	firstNames := []string{}
	//loop through the values and ignore the indices
	for _, val := range bookings {
		//seperate the strings using whitespaces
		names := strings.Fields(val)
		//place the first seperated string into the storage container
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email:")
	fmt.Scan(&email)
	fmt.Println("how many tickets would you like to purchase?")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
func bookingLogic(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - int(userTickets)
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for purchasing %v tickets, you will recieve a confirmation email at %v shortly \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are left\n", remainingTickets)

}
