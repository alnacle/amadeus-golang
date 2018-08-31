package amadeus

import (
	"fmt"
	"os"

	"github.com/alnacle/amadeus-golang/pkg/amadeus/client"
)

// Location defines a handy list of location types, to be used in the locations API
type Location string

const (
	// Airport type
	Airport Location = "AIRPORT"
	// City type
	City Location = "CITY"
	// Any location type
	Any Location = "AIRPORT,CITY"
)

// Direction types, to be used in the Flight Busiest Period API:
type Direction string

const (
	// Arriving type
	Arriving Direction = "ARRIVING"
	// Departing type
	Departing Direction = "DEPARTING"
)

// Amadeus client library for accessing the travel APIs
type Amadeus struct {
	client *client.Client

	// Shopping namespace
	Shopping *ShoppingStruct

	// Version of the SDK
	Version string
}

// Create creates a new ClIent instance
func Create(a ...string) *Amadeus {

	clientID := os.Getenv("AMADEUS_CLIENT_ID")
	secretID := os.Getenv("AMADEUS_CLIENT_SECRET")

	c := client.NewClient(clientID, secretID, "https://test.api.amadeus.com")
	s := NewShopping(c)

	return &Amadeus{
		client:   c,
		Shopping: s,
		Version:  Version,
	}
}

func init() {
	fmt.Println("amadeus package initialized")
}
