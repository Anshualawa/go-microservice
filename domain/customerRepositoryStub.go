package domain

// TODO: cretae structure CustomerRepositoryStub has variable customers
type CustomerRepositoryStub struct {
	customers []Customer
}

// TODO: Implement FindAll() of receiver type CustomerRepositoryStub
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// TODO: Hendler function to return CustomerRepositoryStub object. - dummy customer record.
func NewCutomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Pappu Alawa", "2002-02-02", "Madhya Pradesh", "45454545", "1"},
		{"1001", "Pappu Alawa", "2002-02-02", "Madhya Pradesh", "45454545", "1"},
		{"1001", "Pappu Alawa", "2002-02-02", "Madhya Pradesh", "45454545", "1"},
	}
	return CustomerRepositoryStub{customers}
}
