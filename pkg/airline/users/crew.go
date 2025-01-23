package users

import "airline/pkg/airline"

type Crew struct {
	airline.BaseUser
}

func NewCrew(name string) *Crew {
	return &Crew{
		BaseUser: airline.BaseUser{
			Name: name,
			Role: airline.CREW,
		},
	}
}

func (c *Crew) GetRole() string {
	return string(c.Role)
}
