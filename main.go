package main

import (
	"airline/pkg/airline"
	"airline/pkg/airline/users"
	"fmt"
	"time"
)

func main() {
	airIndia := airline.NewAircraft("Boeing", "747", "BO747", "AirIndia", 110)
	Indigo := airline.NewAircraft("Air Bus", "A320", "A320", "AirIndia", 110)
	flightAirIndia := airline.NewFlight(*airIndia, time.Date(2023, time.October, 5, 14, 30, 0, 0, time.UTC), "Delhi", "Mumbai", 5000.00)
	_ = airline.NewFlight(*Indigo, time.Date(2023, time.October, 5, 12, 10, 3, 0, time.UTC), "Delhi", "Mumbai", 4800.00)

	fmt.Println("==================================================================")
	user := users.NewPassenger("Abhishek")
	fmt.Println("Name of the user", user.GetName())

	airlineManagemet := airline.NewAirlineManagement(flightAirIndia, airIndia)

	fmt.Println("==================================================================")
	ticket := airlineManagemet.BookFlight(flightAirIndia, user, 12)
	fmt.Println("Booking ID: ", ticket.BookingId)

	fmt.Println("==================================================================")
	baggage, err := airlineManagemet.AddBagage(ticket, 12)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("baggage weitght: ", baggage.GetWeight())

	fmt.Println("==================================================================")

	err = airline.Search("Delhi", "Mumbai", "2023-10-05")
	if err != nil {
		fmt.Println(err.Error())
	}

}
