package users

import "airline/pkg/airline"

type Admin struct {
	airline.BaseUser
}

func NewAdmin(name string) *Admin {
	return &Admin{
		BaseUser: airline.BaseUser{
			Name: name,
			Role: airline.ADMIN,
		},
	}
}

func (c *Admin) GetRole() string {
	return string(c.Role)
}
