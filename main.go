package main

import (
	"fmt"
	"strconv"
)
  
var stationName string = "Oreva metro station"
const trainTickets int = 50
var remainingTickets int = 50
var bookings = make([]map[string]string, 0)

type userData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main() {

	greetUser()

	for { 

		//calls getUserInput function
		firstName, lastName, email, userTickets := getUserInput()

		isValidEmail, isValidNames, isValidTickets := validateUserInput(firstName, lastName, email, userTickets, uint(remainingTickets))

		//checks to make sure the user doesn't input a number > than tickets remaining
		if isValidTickets && isValidEmail && isValidNames {

			bookingLogic(userTickets, firstName, lastName, email)

			//call first name function
			firstNames := userFirstName()
			fmt.Printf("The first names of the bookings are:%v\n", firstNames)

			//checks if tickects are still available
			if remainingTickets == 0 {
				//end program
				fmt.Println("Train tickets are booked out.")
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
	//loops through the slice of map and extract the first names
	for _, val := range bookings {
		firstNames = append(firstNames, val["firstName"])
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

	//create a map a user map 
	 var userData =  make(map[string]string)
	 userData["firstName"] = firstName
	 userData["lastName"] = lastName
	 userData["email"] = email
	 userData["numberOfTickets"] = strconv.Itoa(int(userTickets))
	

	bookings = append(bookings, userData)

	fmt.Printf("list of bookings is %v:\n", bookings)
	fmt.Printf("Thank you %v %v for purchasing %v tickets, you will recieve a confirmation email at %v shortly \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are left\n", remainingTickets)

}
