package main

import "fmt"

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

		remainingTickets = remainingTickets - int(userTickets)
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for purchasing %v tickets, you will recieve a confirmation email at %v shortly \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are left\n", remainingTickets)

	fmt.Printf("All bookings:%v\n", bookings)
	}

	

}
