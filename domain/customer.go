package domain

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerReposity interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, error)
}
