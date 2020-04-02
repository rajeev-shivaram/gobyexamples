package magazine

//Subscriber blah
type Subscriber struct {
	Name    string
	Rate    float64
	Active  bool
	Address // anonymous struct field
}

// Address struct
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

// Employee exported from here
type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address
}
