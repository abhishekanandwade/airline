package airline

type AirlineMangement struct {
	AddFlight *Flight
	Aircraft  *Aircraft
}

func NewAirlineManagement(flight *Flight, aircraft *Aircraft) *AirlineMangement {
	return &AirlineMangement{
		AddFlight: flight,
		Aircraft:  aircraft,
	}
}

func (a *AirlineMangement) BookFlight(flight *Flight, user *User, seatNo int) *flightTicket {
	return newBookFlight(flight, user, seatNo)
}

func (a *AirlineMangement) AddBagage(ticket *flightTicket, weight int) *baggage {
	return newBaggage(ticket, weight)
}
