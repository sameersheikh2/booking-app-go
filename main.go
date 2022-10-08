package main

import (
	"fmt"
	"strings"
)

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
		if userTickets > ticketsAvailable {
			fmt.Printf("We only have %v tickets remaining , so you can't book %v tickets.\nTry again\n", ticketsAvailable, userTickets)
			continue
		}

		bookings = append(bookings, firstName+" "+lastName)
		ticketsAvailable = ticketsAvailable - userTickets

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
		fmt.Println(bookings)

		/*for initialization; condition; update {
					statement(s)
				  }.
		The initialization initializes and/or declares variables and is executed only once.
		Then, the condition is evaluated. If the condition is true, the body of the for loop is executed.
		The update updates the value of initialization.
		The condition is evaluated again. The process continues until the condition is false.
		If the condition is false, the for loop ends*/

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)       //string.Fields split strings eg. sameer sheikh = "sameer" "sheikh"
			firstNames = append(firstNames, names[0]) //name holding value eg. "sameer" "sheikh" so name[0] at index 0 we have value sameer
		}
		// above we defined an slice of strings(firstNames) and looped through BOOKINGS for both its index and value using the for..range keyword

		if ticketsAvailable == 0 {
			fmt.Println("The tickets are all sold out. Come back next year.")
			break
		}

		fmt.Printf("%v tickets remaining for Go Conference.\n", ticketsAvailable)
		fmt.Printf("The first name of bookings are %v\n", firstNames)
	}

}
