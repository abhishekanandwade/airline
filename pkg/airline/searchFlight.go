package airline

import "errors"

var AllFlights []*Flight

func Search(source string, destinatin string, date string) ([]*Flight, error) {
	if source == "" || destinatin == "" || date == "" {
		return nil, errors.New("invalid search parameters")
	}

	var resultFlights []*Flight

	for _, flight := range AllFlights {
		if flight.GetSource() == source && flight.GetDestination() == destinatin && flight.GetDate() == date {
			resultFlights = append(resultFlights, flight)
		}
	}

	return resultFlights, nil
}
