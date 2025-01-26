package airline

import "math/rand"

type flightTicket struct {
	BookingId int
	Flight    *Flight
	User      User
	SeatNo    int
}

func newBookFlight(flight *Flight, user User, seatNo int) *flightTicket {
	return &flightTicket{
		BookingId: rand.Int(),
		Flight:    flight,
		User:      user,
		SeatNo:    seatNo,
	}
}

func (b *flightTicket) GetFlightDetails() *Flight {
	return b.Flight
}

func (b *flightTicket) GetSeatNumber() int {
	return b.SeatNo
}

func (b *flightTicket) GetUser() User {
	return b.User
}
