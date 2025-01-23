package users

import "airline/pkg/airline"

type Staff struct {
	airline.BaseUser
}

func NewStaff(name string) *Staff {
	return &Staff{
		BaseUser: airline.BaseUser{
			Name: name,
			Role: airline.STAFF,
		},
	}
}

func (p *Staff) GetRole() string {
	return string(p.Role)
}
