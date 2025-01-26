package airline

import (
	"errors"
	"sync"
	"time"
)

type Flight struct {
	Aircraft       Aircraft
	DepartureTime  time.Time
	Source         string
	Destination    string
	AvailableSeats int
	mu             sync.RWMutex
}

func NewFlight(aircraft Aircraft, departureTime time.Time, source string, destination string) *Flight {
	flight := &Flight{
		Aircraft:      aircraft,
		DepartureTime: departureTime,
		Source:        source,
		Destination:   destination,
	}

	AllFlights = append(AllFlights, flight)
	return flight
}

func (f *Flight) BookSeat(count int) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.AvailableSeats >= count {
		f.AvailableSeats -= count
		return nil
	}

	return errors.New("Seats Unavailable")
}

func (f *Flight) GetAvailableSeats() int {
	f.mu.RLock()
	defer f.mu.Unlock()
	return f.AvailableSeats
}

func (f *Flight) GetSource() string {
	return f.Source
}

func (f *Flight) GetDestination() string {
	return f.Destination
}

func (f *Flight) GetDate() string {
	return f.DepartureTime.Format(time.DateOnly)
}
