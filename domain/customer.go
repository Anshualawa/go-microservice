package domain

// struct
type Customer struct {
	Id          string
	Name        string
	DateOfBirth string
	City        string
	Zipcode     string
	Status      string
}

// Interface

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
