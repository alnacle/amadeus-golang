package shopping

import (
	"github.com/alnacle/amadeus-golang/pkg/amadeus/client"
)

// FlightDestination represents the Flight Inspiration Search
type FlightOffers struct {
	Client *client.Client
}

func NewFlightOffers(client *client.Client) *FlightOffers {
	return &FlightOffers{client}
}

// Get JSON response
func (f *FlightOffers) Get(params ...string) ([]byte, error) {
	return f.Client.Get("/v1/shopping/flight-offers", params...)
}
