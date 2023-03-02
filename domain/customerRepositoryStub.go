package domain

type CustomerReposityStub struct {
	customers []Customer
}

func (s CustomerReposityStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerReposityStub {
	customers := []Customer{
		{"1001", "Virat", "Banglore", "000000", "1993-05-8", "1"},
		{"1002", "Rohit", "Mumbai", "111111", "1994-04-2", "1"},
	}
	return CustomerReposityStub{customers}
}
