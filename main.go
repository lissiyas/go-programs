package main

import (
	"fmt"
)


func main()  {
	var conferenceName = "welcome"
	const conferenceTicket = 50

	var  remainingTickets uint = 50
	var bookings   [50]string

	fmt.Printf("conference tickets is %T , remainingtickets is %T, conferenceName is %T\n",conferenceTicket,remainingTickets,conferenceName )
	fmt.Printf("welcome to the %v booking application\n", conferenceName)
	fmt.Printf("we have total off %v tickets and %v are still avsailable\n", conferenceTicket, remainingTickets)
	fmt.Println("get your tickets here to attend")
	


	
	
	var firstaName string 
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("enter your first name:")
	fmt.Scan(&firstaName)
	fmt.Println("enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("enter your email:")
	fmt.Scan(&email)
	fmt.Println("enter your tickets:")
	fmt.Scan(&userTickets)

	remainingTickets= remainingTickets - userTickets

	bookings[0] = firstaName + " " + lastName

	fmt.Printf("%v %v  with email id  %v booked %v number of tickets\n", firstaName,lastName,email,userTickets)
	fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName)
 	
}