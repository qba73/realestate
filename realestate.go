package realestate

import "time"

// Agent represents a property agent working in the real estate company.
type Agent struct {
	Id          int
	Name        string
	Photo       string
	Description string
	Email       string
	Phone       string
	IsMVP       bool
	HireDate    time.Time
}

// Listing represents a property listed for sale by the agent.
type Listing struct {
	ID           int
	Agent        int
	Title        string
	Address      string
	City         string
	County       string
	PostCode     string
	Description  string
	Price        int
	Bedrooms     int
	Bathrooms    int
	Garage       int
	Area         int
	BackyardSize int
	IsPublished  bool
	ListDate     time.Time
	PhotoMain    string
	Photo1       string
	Photo2       string
	Photo3       string
	Photo4       string
	Photo5       string
	Photo6       string
}
