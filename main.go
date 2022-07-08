package main

import (
	"fmt"
	"strings"
)

func main() {
	var stationName string = "Oreva metro station"
	const trainTickets int = 50
	var remainingTickets int = 50
	var bookings []string

	fmt.Printf("Welcome to %v.\n", stationName)
	fmt.Printf("we have a total of %v tickets but %v left.\n", trainTickets, remainingTickets)
	fmt.Println("Please book your tickets")
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	for {

		fmt.Println("Please enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Please enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Please enter your email:")
		fmt.Scan(&email)
		fmt.Println("how many tickets would you like to purchase?")
		fmt.Scan(&userTickets)

		//checks to make sure the user doesn't input a number > than tickets remaining
		if userTickets <= uint(remainingTickets) {

			remainingTickets = remainingTickets - int(userTickets)
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for purchasing %v tickets, you will recieve a confirmation email at %v shortly \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets are left\n", remainingTickets)

			//create a storage container
			firstNames := []string{}
			//loop through the values and ignore the indices
			for _, val := range bookings {
				//seperate the strings using whitespaces
				names := strings.Fields(val)
				//place the first seperated string into the storage container
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of the bookings are:%v\n", firstNames)

			//checks if tickects are still available
			if remainingTickets == 0 {
				fmt.Println("Train tickets are booked out.  ")
				break
			}
		} else {
			fmt.Printf("we only have %v tickets left, you can not book %v tickets.\n", remainingTickets, userTickets)
			
		}

	}

}
