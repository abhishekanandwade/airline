package users

import "airline/pkg/airline"

type Passenger struct {
	airline.BaseUser
}

func NewPassenger(name string) *Passenger {
	return &Passenger{
		BaseUser: airline.BaseUser{
			Name: name,
			Role: airline.PASSENGER,
		},
	}
}

func (p *Passenger) GetRole() string {
	return string(p.Role)
}

func (p *Passenger) GetName() string {
	return p.Name
}
