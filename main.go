package main

import "fmt"

func main() {

	//assign variable wihtout datatype because go will automatically understand data type by looking at value assign
	var totalTickets = 50
	var ticketsAvailable uint = 50 //uint var type allow only positive number
	// var bookings [50]string        //this is array.first it need array size and second is value type that comes in array
	// var bookings []string //slice in go, same as array but dynamic and no need to define size as we do in array
	bookings := []string{}
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Welcome to Go Conference booking application")

	fmt.Printf("We have total of %v tickets and %v are still available\n", totalTickets, ticketsAvailable) //%v stand for value. formatting using printf

	fmt.Println("Get your tickets here to attend")

	for {
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email name:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		bookings = append(bookings, firstName+""+lastName)
		ticketsAvailable = ticketsAvailable - userTickets

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)

		fmt.Printf("%v tickets remaining for Go Conference.\n", ticketsAvailable)
		fmt.Printf("%v bookings\n", bookings)
	}

}
