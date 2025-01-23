package airline

type User interface {
	GetRole() string
}

type UserType string

const (
	PASSENGER UserType = "passenger"
	CREW      UserType = "crew"
	STAFF     UserType = "staff"
	ADMIN     UserType = "admin"
)

type BaseUser struct {
	Name string
	Role UserType
}
