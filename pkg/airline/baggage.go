package airline

type baggage struct {
	BookedFlight *flightTicket
	Weight       int
}

func newBaggage(bookedFlight *flightTicket, weight int) *baggage {
	return &baggage{
		BookedFlight: bookedFlight,
		Weight:       weight,
	}
}

func (b *baggage) GetFlight() *flightTicket {
	return b.BookedFlight
}

func (b *baggage) GetWeight() int {
	return b.Weight
}
