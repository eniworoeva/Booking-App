package main

import "fmt"

func main()  {
	var stationName string = "Oreva metro station"
	const trainTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("Welcome to %v.\n", stationName)
	fmt.Printf("we have a total of %v tickets but %v left.\n",trainTickets, remainingTickets )
	fmt.Println("Please book your tickets")

	var userName string
	var userTickets uint

	fmt.Println("Please enter your name:")
	fmt.Scan(&userName) 
	fmt.Println("how many tickets would you like to purchase?")
	fmt.Scan(&userTickets)

	fmt.Printf("%v has booked %v tickets.\n",userName, userTickets )

	 





}
