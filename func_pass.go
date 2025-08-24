package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Go conference"
	var remainingTickets uint
	const conferenceTickets = 50
	var userName string
	var userTickets uint
	var bookings []string
	remainingTickets = 50
	for {
		if !checkBooked(userTickets, remainingTickets, bookings) {
			break
		}
		intro(conferenceName, remainingTickets, conferenceTickets)
		userName, userTickets = get()
		if check(userTickets, remainingTickets, bookings) {
			remainingTickets, bookings = book(userName, userTickets, remainingTickets, bookings)
		}
		continue
	}
}

func intro(conferenceName string, remainingTickets uint, conferenceTickets uint) {
	fmt.Printf("Welcome to %v !\n", conferenceName)
	fmt.Printf("%v of %v tickets remaining", remainingTickets, conferenceTickets)
}

func get() (string, uint) {
	var userName string
	var userTickets uint
	fmt.Println("\nEnter your name")
	fmt.Scan(&userName)
	fmt.Println("Enter no. of tickets")
	fmt.Scan(&userTickets)
	return userName, userTickets
}

func check(userTickets uint, remainingTickets uint, bookings []string) bool {
	if userTickets > remainingTickets {
		fmt.Printf("Sorry we only have %v tickets remaining\n", remainingTickets)
		return false
	}
	return true
}

func checkBooked(userTickets uint, remainingTickets uint, bookings []string) bool {
	if remainingTickets == 0 {
		fmt.Printf("\nBooked Out, bookings list: %v", bookings)
		return false
	}
	return true
}

func book(userName string, userTickets uint, remainingTickets uint, bookings []string) (uint, []string) {
	bookings = append(bookings, userName)
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Congratulations %v, you've booked %v tickets\n", userName, userTickets)
	return remainingTickets, bookings
}
