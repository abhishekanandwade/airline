package airline

type Aircraft struct {
	Make        string
	Model       string
	ModelNumber string
	Airline     string
	TotalSeats  int
}

func NewAircraft(make string, model string, modelNumber string, airline string, seats int) *Aircraft {
	return &Aircraft{
		Make:        make,
		Model:       model,
		ModelNumber: modelNumber,
		Airline:     airline,
		TotalSeats:  seats,
	}
}

func (a *Aircraft) GetTotalSeats() int {
	return a.TotalSeats
}

func (a *Aircraft) GetAirline() string {
	return a.Airline
}
