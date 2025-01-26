package airline

import (
	"errors"
	"fmt"
)

var AllFlights []*Flight

func Search(source string, destinatin string, date string) error {
	if source == "" || destinatin == "" || date == "" {
		return errors.New("invalid search parameters")
	}

	for _, flight := range AllFlights {
		if flight.GetSource() == source && flight.GetDestination() == destinatin && flight.GetDate() == date {

			fmt.Println("Airline:", flight.GetAirline(), "\nFrom: ", flight.GetDestination(), "\nTo: ", flight.GetDestination(), "\nDate: ", flight.GetDate(), "\nPrice: ", flight.GetPrice())
			fmt.Println("***********")

		}
	}

	return nil
}
